package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	word              = flag.String("word", "", "put a key word that you want to search")
	number            = flag.Int("number", 10, "how many tweets do you need?")
	consumerKey       = flag.String("consumer-key", "", "set your consumer key")
	consumerSecret    = flag.String("consumer-secret", "", "set your consumer secret")
	accessToken       = flag.String("access-token", "", "set your accessToken")
	accessTokenSecret = flag.String("accessTokenSecret", "", "set your accessTokenSecret")
)

// Client for twitter API
type Client struct {
	API *anaconda.TwitterApi
}

// New Creates twitter-api client
func New(key, secret, token, tokenSecret string) *Client {
	if key == "" {
		key = os.Getenv("TWITTER_CONSUMER_KEY")
	}
	if secret == "" {
		secret = os.Getenv("TWITTER_CONSUMER_SECRET")
	}
	if token == "" {
		token = os.Getenv("TWITTER_ACCESS_TOKEN")
	}
	if tokenSecret == "" {
		tokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	}
	anaconda.SetConsumerKey(key)
	anaconda.SetConsumerSecret(secret)
	api := anaconda.NewTwitterApi(token, tokenSecret)

	return &Client{
		API: api,
	}
}

func main() {
	flag.Parse()

	cli := New(*consumerKey, *consumerSecret, *accessToken, *accessTokenSecret)
	tweets := cli.GetTweets(*word, *number)
	if len(tweets) == 0 {
		log.Printf("buzz-tweet not found with keyword %s", *word)
		os.Exit(0)
	}

	for _, tweet := range tweets {
		fmt.Fprintln(os.Stdout, tweet.Text)
	}
}

// GetBuzzTweet get buzz tweet
func (cli *Client) GetTweets(word string, num int) []anaconda.Tweet {
	// get a stream filtered by the word from flag
	stream := cli.API.PublicStreamFilter(url.Values{
		"track": []string{"#" + word},
	})
	defer stream.Stop()

	var tweets []anaconda.Tweet
	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			continue
		}

		// buzz-tweet recognize tweet with 10000 favoriteCount as "buzz-tweet"
		if t.FavoriteCount > 10000 {
			tweets = append(tweets, t)
		}
		if len(tweets) == num {
			break
		}
	}
	return tweets
}