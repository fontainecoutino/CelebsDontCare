package retrieve

import (
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
	getTweetsfromUser(userID, Path+"raw.json")
}

/**
 *  The format of the file after the function is as follows. The tweets are
 *  stored to keep only the text and date of creation. Each tweet in the file is stored as a map
 * 	and a comma is added after wards. This is true for every tweet; which means that the function
 *  will produce a file that is not in json format. This should be fixed by the caller.
 *  all current tweets from TwitterUserID and store in tweets.json
 */
func getTweetsfromUser(userID string, destination string) {
	nextToken := ""
	for {
		// executes command to get data and stores it in database/retrieve/raw.json
		prg := Path + "retrieve"
		_, err := exec.Command("bash", prg, userID, nextToken).Output()
		if err != nil {
			fmt.Println("> Error executing bash command: " + err.Error())
			return
		}

		// gets raw data from file
		rawDataFile, _ := os.Open(destination)
		rawData, err := ioutil.ReadAll(rawDataFile)
		if err != nil {
			fmt.Println("> Error tranformig " + destination + " to []byte: " + err.Error())
		}
		rawDataFile.Close()

		var data map[string]interface{}
		json.Unmarshal(rawData, &data)

		// gets all tweets, modifies them to keep only tweets and writes them to tweets.json
		currentTweets := data["data"].([]interface{})
		for _, tweet := range currentTweets {
			tweet := tweet.(map[string]interface{})
			delete(tweet, "id")
			writeTweetToFile(tweet, Path+"tweets.json", ",")
		}

		// gets metadata to update retrieve
		metaData := data["meta"].(map[string]interface{})
		if token, ok := metaData["next_token"].(string); ok {
			nextToken = token
		} else {
			break
		}
	}
}

func writeTweetToFile(tweet map[string]interface{}, file string, comma string) {
	bytes, _ := json.Marshal(tweet)

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("> Error storing tweet in tweets.json: " + err.Error())
		return
	}
	defer f.Close()

	if _, err = f.WriteString(string(bytes) + comma); err != nil {
		fmt.Println("> Error storing tweet in tweets.json: " + err.Error())
	}
}

/**
 * read tweets.json to get trip information and store in trips.json
 */
