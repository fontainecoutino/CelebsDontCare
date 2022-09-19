package retrieve

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const Path = "./database/retrieve/"

/**
 *  Given an userID, gets all of their tweets. The tweets are stored in the file
 *  tweets.json. The tweets are stored to keep only the text and date of creation.
 */
func GetData() {
	userID := "1450174360346574850" // @CelebJets
	writeTweetsToFile(userID, Path+"tweets.json")
}

/**
 *  The format of the file after the function is as follows. The tweets are
 *  stored to keep only the text and date of creation. Each tweet in the file is stored as a map
 * 	and a comma is added after wards. This is true for every tweet; which means that the function
 *  will produce a file that is not in json format. This should be fixed by the caller.
 *  all current tweets from TwitterUserID and store in tweets.json
 */
func writeTweetsToFile(userID string, destination string) {
	tempFile := Path + "tempTweets.json"

	oldestID := getTweetsfromUser(userID, tempFile)

	outfile, err := os.Create(destination)
	if err != nil {
		fmt.Println("> Error creating " + destination + ": " + err.Error())
		return
	}
	defer outfile.Close()

	// add beggining of json object
	usrSting := "\"user_id\": \"" + userID + "\",\n"
	oldestSting := "\"oldest_tweet_id\": \"" + oldestID + "\",\n"
	tweetsString := "\"tweets\":[\n"
	_, err = outfile.WriteString("{ " + usrSting + oldestSting + tweetsString)
	if err != nil {
		fmt.Println("> Error writing to " + destination + ": " + err.Error())
		return
	}

	// write contents from tmp to desination file
	f, err := os.Open(tempFile)
	if err != nil {
		fmt.Println("> Error opening " + tempFile + ": " + err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, err = outfile.WriteString(scanner.Text() + "\n")
		if err != nil {
			fmt.Println("> Error writing to " + destination + ": " + err.Error())
			return
		}
	}

	// write end of file to finish object
	_, err = outfile.WriteString("]}")
	if err != nil {
		fmt.Println("> Error writing to " + destination + ": " + err.Error())
		return
	}

	// ensure all lines are written
	outfile.Sync()
	outfile.Close()
	f.Close()

	// gets rid of file since there is no use for it anymore
	_, err = exec.Command("rm", tempFile).Output()
	if err != nil {
		fmt.Println("> Error deleting " + tempFile + ": " + err.Error())
		return
	}

	fmt.Println("> Done writing tweets to " + destination)
}

/**
 *  The format of the file after the function is as follows. The tweets are
 *  stored to keep only the text and date of creation. Each tweet in the file is stored as a map
 * 	and a comma is added after wards. This is true for every tweet; which means that the function
 *  will produce a file that is not in json format. This should be fixed by the caller.
 *  all current tweets from TwitterUserID and store in tweets.json
 */
func getTweetsfromUser(userID string, tempFile string) string {
	nextToken := ""
	oldestID := ""
	for {
		// executes command to get data and stores it in database/retrieve/raw.json
		prg := Path + "retrieve"
		_, err := exec.Command("bash", prg, userID, nextToken).Output()
		if err != nil {
			fmt.Println("> Error executing bash command: " + err.Error())
			return oldestID
		}

		// gets raw data from file
		rawDataFile, _ := os.Open(Path + "raw.json")
		rawData, err := ioutil.ReadAll(rawDataFile)
		if err != nil {
			fmt.Println("> Error tranformig " + Path + "raw.json" + " to []byte: " + err.Error())
		}
		rawDataFile.Close()

		// gets all tweets and appends them to temp file
		var data map[string]interface{}
		json.Unmarshal(rawData, &data)

		currentTweets := data["data"].([]interface{})
		metaData := data["meta"].(map[string]interface{})
		breakLoopFlag := false
		for index, tweet := range currentTweets {
			tweet := tweet.(map[string]interface{})
			oldestID = metaData["oldest_id"].(string)
			delete(tweet, "id")

			if token, ok := metaData["next_token"].(string); ok {
				nextToken = token
				appendTweetToFile(tweet, tempFile, ",")
			} else if index != len(currentTweets)-1 {
				appendTweetToFile(tweet, tempFile, ",")
			} else {
				appendTweetToFile(tweet, tempFile, "")
				breakLoopFlag = true
			}
		}
		if breakLoopFlag {
			break
		}
	}

	// gets rid of file since there is no use for it anymore
	_, err := exec.Command("rm", Path+"raw.json").Output()
	if err != nil {
		fmt.Println("> Error deleting " + Path + "raw.json" + ": " + err.Error())
		return oldestID
	}
	return oldestID
}

/**
 * Append a tweet to the given file
 */
func appendTweetToFile(tweet map[string]interface{}, file string, comma string) {
	bytes, _ := json.Marshal(tweet)

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("> Error storing tweet in " + file + " " + err.Error())
		return
	}
	defer f.Close()

	if _, err = f.WriteString(string(bytes) + comma + "\n"); err != nil {
		fmt.Println("> Error storing tweet in " + file + " " + err.Error())
	}
}

/**
 * read tweets.json to get trip information and store in trips.json
 */
