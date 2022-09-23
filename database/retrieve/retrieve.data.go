package retrieve

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

/**
 *  Given an userID, gets all of their tweets. The tweets are stored in the file
 *  tweets.json. The tweets are stored to keep only the text and date of creation.
 */
func getData(source string) (int, error) {
	userID := "1450174360346574850" // @CelebJets
	writeTweetsToFile(userID, Path+"tweets.json")

	return 0, nil
}

/**
 *  The format of the file after the function is as follows. The tweets are
 *  stored to keep only the text and date of creation. Each tweet in the file is stored as a map
 * 	and a comma is added after wards. This is true for every tweet; which means that the function
 *  will produce a file that is not in json format. This should be fixed by the caller.
 *  all current tweets from TwitterUserID and store in tweets.json
 */
func writeTweetsToFile(userID string, destination string) {
	allTweets, oldestID := getTweetsfromUser(userID)
	userTweets := UserTweets{
		User_id:         userID,
		Oldest_tweet_id: oldestID,
		Tweets:          allTweets,
	}

	bytes, _ := json.MarshalIndent(userTweets, "", " ")
	err := ioutil.WriteFile(destination, bytes, 0644)
	check(err, true, "> Error writing to "+destination+": ")

	fmt.Println("> Done writing tweets to " + destination)
}

/**
 *  The format of the file after the function is as follows. The tweets are
 *  stored to keep only the text and date of creation. Each tweet in the file is stored as a map
 * 	and a comma is added after wards. This is true for every tweet; which means that the function
 *  will produce a file that is not in json format. This should be fixed by the caller.
 *  all current tweets from TwitterUserID and store in tweets.json
 */
func getTweetsfromUser(userID string) ([]map[string]interface{}, string) {
	var tripLogTweets []map[string]interface{}
	var oldestID string
	var nextToken string
	for {
		// gets all tweets and appends them to temp file
		var data map[string]interface{}
		json.Unmarshal(getRawData(userID, nextToken), &data)

		currentTweets := data["data"].([]interface{})
		metaData := data["meta"].(map[string]interface{})

		// append tweets only if a trip log
		for _, tweet := range currentTweets {
			tweet := tweet.(map[string]interface{})
			if strings.Contains(tweet["text"].(string), "gallons") &&
				strings.Contains(tweet["text"].(string), "CO2 emissions") {
				oldestID = metaData["oldest_id"].(string)
				tripLogTweets = append(tripLogTweets, tweet)
			}
		}

		token, ok := metaData["next_token"].(string)
		if !ok {
			break
		}
		nextToken = token
	}

	// gets rid of file since there is no use for it anymore
	_, err := exec.Command("rm", Path+"raw.json").Output()
	check(err, false, "> Error deleting "+Path+"raw.json"+": ")
	return tripLogTweets, oldestID
}

/**
 * Executes bash command to get 100 tweets from userID and store in raw.json.
 */
func getRawData(userID string, nextToken string) []byte {
	// executes command to get data and stores it in database/retrieve/raw.json
	prg := Path + "retrieve"
	_, err := exec.Command("bash", prg, userID, nextToken).Output()
	check(err, false, "> Error executing bash command: ")

	// gets raw data from file
	rawDataFile, _ := os.Open(Path + "raw.json")
	rawData, err := ioutil.ReadAll(rawDataFile)
	check(err, false, "> Error tranformig "+Path+"raw.json"+" to []byte: ")
	rawDataFile.Close()

	return rawData
}

/**
 * Checks for an error. Displays message and panics if needed.
 */
func check(err error, panicCheck bool, msg string) {
	if err != nil {
		fmt.Println(msg + err.Error())
		if panicCheck {
			panic(err)
		}
	}
}
