package retrieve

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/fontainecoutino/CelebsDontCare/database"
	"github.com/fontainecoutino/CelebsDontCare/trip"
	_ "github.com/fontainecoutino/CelebsDontCare/trip"
)

const URL = "http://localhost:5000/api/"

/**
 *  Given the source, it retrieves data from trips and stors them into the database
 */
func getData(source string) (int, error) {
	var newTrips int
	var err error
	if source == "twitter" {
		// get all possible sources
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		results, err := database.DB.QueryContext(ctx, `SELECT * FROM twitter_sources `)
		if err != nil {
			log.Println(err.Error())
			return 0, err
		}
		defer results.Close()

		sources := make([]Twitter_Source, 0)
		for results.Next() {
			var source Twitter_Source
			results.Scan(
				&source.Twitter_id,
				&source.Username,
				&source.Last_tweet_id)

			sources = append(sources, source)
		}

		// for all users write to data base
		var nt int
		for _, source := range sources {
			nt, err = writeTweetsToDatabase(source)
			newTrips += nt
			if err != nil {
				return newTrips, err
			}
		}
	}
	return newTrips, err
}

/**
 * Based on the twitter source, get all th valid tweets that are trips an stores
 * them into data base. It then updates the twitter source to store their latest
 * tweet id. Lastly, it writes a log if at least one trip was written to database.
 */
func writeTweetsToDatabase(source Twitter_Source) (int, error) {
	allTweets, oldestID := getTweetsfromUser(source.Twitter_id, source.Last_tweet_id)

	// write to trips to database
	var num int
	for _, tweet := range allTweets {
		num += insertTripIntoDatabase(tweet)
	}

	// update to twitter_sources
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`UPDATE twitter_sources SET last_tweet_id = $1 WHERE twitter_id = $2`,
		oldestID, source.Twitter_id)

	if err != nil {
		log.Println(err.Error())
		return num, err
	}

	// update trip log
	if allTweets != nil {
		tweets := &Tweets{}
		tweets.Tweets = allTweets
		j, err := json.Marshal(tweets)
		ctx1, cancel1 := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel1()
		_, err = database.DB.ExecContext(ctx1, `INSERT INTO new_trip_log (log)  VALUES ($1)`, j)

		if err != nil {
			log.Println(err.Error())
			return num, err
		}
	}

	return num, nil
}

/**
 * Based on the text it strips the info to store into the trip. It then
 * stores into database. Returns 1 if successful.
 */
func insertTripIntoDatabase(tweet map[string]interface{}) int {
	tripText := strings.Replace(tweet["text"].(string), "\n", "", -1)
	fmt.Println(tripText)
	split := strings.Split(tripText, "~")

	for i, info := range split {
		split[i] = strings.TrimSpace(info)
	}

	// name
	name := strings.Split(split[0], "'")[0]

	// destination / mileage
	var fromTo string
	var distance int
	if strings.Contains(split[0], "mile") {
		fromTo = strings.TrimSpace(strings.Split(split[0], "from")[1])
		m := strings.Split(strings.TrimSpace(strings.Split(split[0], "mile")[0]), " ")
		distance, _ = strconv.Atoi(strings.Replace(m[len(m)-1], ",", "", -1))
	}

	// gallons
	gallons, _ := strconv.Atoi(strings.Replace(strings.Split(split[1], " ")[0], ",", "", -1))

	// cost
	cost, _ := strconv.Atoi(strings.Replace(strings.Split(split[3], " ")[0][1:], ",", "", -1))

	newTrip := trip.Trip{
		TimeStamp:   tweet["created_at"].(string),
		Name:        name,
		Distance:    distance,
		GallonsUsed: gallons,
		CostOfFuel:  cost,
		Flight:      fromTo,
	}

	jsonStr, _ := json.Marshal(newTrip)
	req, _ := http.NewRequest("POST", URL+"trips", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	return 1
}

/**
 *  Based on the userid, it gets al the tweets that are trips and have not been
 *  evaluated since the lastTweetId. Temporarily stores into a raw file
 *  but then deletes.
 */
func getTweetsfromUser(userID string, lastTweetID string) ([]map[string]interface{}, string) {
	var tripLogTweets []map[string]interface{}
	oldestID := lastTweetID
	var nextToken string
	for {
		// gets all tweets and appends them to temp file
		var data map[string]interface{}
		json.Unmarshal(getRawData(userID, nextToken), &data)

		currentTweets := data["data"].([]interface{})
		metaData := data["meta"].(map[string]interface{})

		// append tweets only if a trip log
		lastTweetReached := false
		for _, tweet := range currentTweets {
			tweet := tweet.(map[string]interface{})
			if tweet["id"].(string) == lastTweetID {
				lastTweetReached = true
				break
			}
			if strings.Contains(tweet["text"].(string), "'s") &&
				strings.Contains(tweet["text"].(string), "gallons") &&
				strings.Contains(tweet["text"].(string), "CO2 emissions") {

				oid, _ := strconv.Atoi(oldestID)
				tid, _ := strconv.Atoi(tweet["id"].(string))
				if tid > oid {
					oldestID = tweet["id"].(string)
				}

				tripLogTweets = append(tripLogTweets, tweet)
			}
		}

		if lastTweetReached {
			break
		}

		token, tokenExists := metaData["next_token"].(string)
		if !tokenExists {
			break
		}
		nextToken = token
	}

	// gets rid of file since there is no use for it anymore
	_, err := exec.Command("rm", Path+"raw.json").Output()
	check(err, "> Error deleting "+Path+"raw.json"+": ")
	return tripLogTweets, oldestID
}

/**
 * Executes bash command to get 100 tweets from userID and store in raw.json.
 */
func getRawData(userID string, nextToken string) []byte {
	// executes command to get data and stores it in database/retrieve/raw.json
	prg := Path + "retrieve"
	_, err := exec.Command("bash", prg, userID, nextToken).Output()
	check(err, "> Error executing bash command: ")

	// gets raw data from file
	rawDataFile, _ := os.Open(Path + "raw.json")
	rawData, err := ioutil.ReadAll(rawDataFile)
	check(err, "> Error tranformig "+Path+"raw.json"+" to []byte: ")
	rawDataFile.Close()

	return rawData
}

/**
 * Checks for an error. Displays message and panics if needed.
 */
func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg + err.Error())
	}
}
