package retrieve

const Path = "./database/retrieve/"

type Twitter_Source struct {
	Twitter_id    string
	Username      string
	Last_tweet_id string
}

type Tweets struct {
	Tweets []map[string]interface{} `json:tweets`
}
