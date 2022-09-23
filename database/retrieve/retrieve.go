package retrieve

const Path = "./database/retrieve/"

type UserTweets struct {
	User_id         string                   `json:"user_id"`
	Oldest_tweet_id string                   `json:"oldest_tweet_id"`
	Tweets          []map[string]interface{} `json:"tweets"`
}
