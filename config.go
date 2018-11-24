package main

import (
	"flag"

	"github.com/ChimeraCoder/anaconda"
)

// set up keys to use  twitter api
const (
	consumerKey       = "" //set your consumer key
	consumerSecret    = "" //set your consumer secret
	accessToken       = "" //set your token
	accessTokenSecret = "" //set your token secret
)

// returns pointer to Twiiter api
func getApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return api
}

// input from terminal
var (
	word   = flag.String("word", "", "put a key word thay you want to search")
	length = flag.Int("length", 10, "how many tweets do you need ?")
)
