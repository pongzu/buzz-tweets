package main

import (
	"flag"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	flag.Parse()
	createDir()

}

// create directory to save conetents of tweets
func createDir() {
	var name = "buzz_tweets"
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.MkdirAll(name, 0755); err != nil {
			log.Fatal(err)
		}
	}
}

func GetResult(api *anaconda.TwitterApi, word string, length int) {
	var counter int
	// get a stream filterd by the word from flag
	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#" + word},
	})
	defer stream.Stop()

	var subCounter int
	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Println("got unexepected value")
			continue
		}
		if !judgeBuzzed(t.FavoriteCount) {
			subCounter++
			if subCounter > 1000 {
				log.Println("can not findmore tweets ")
				break
			}
			continue
		}
		// save the tweet that is liked by 10000 people
		save(t.Text, counter)
		counter++
	}
}

func judgeBuzzed(count int) bool {
	return count > 10000
}

func save(text string, counter int) {
	fileName := "buzz_tweets/text_" + strconv.Itoa(counter) + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(file, text)
	if err != nil {
		log.Fatal(err)
	}
}
