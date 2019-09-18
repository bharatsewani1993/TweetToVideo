package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

//TwitterLogin is used for authentication purpose
func TwitterLogin() *anaconda.TwitterApi {
	//API Key and Access Token
	consumerkey := "rQhHlYBINp4AdteYlA2KdOVZL"
	consumersecret := "5HBHhezNwUQkfNvKIlXnhAgt6DSfqbLS4q6c9polwQEDcCd61P"
	accesstoken := "3181658732-NZn8b404vANoRJJ6gzdA3QXHLGEGfyGoqEQrszD"
	accesstokensecret := "YApNM25YmLr6QjCd7fBYY4BEvE49fankFhZS293E0cYvZ"

	//Authentication With Twitter
	anaconda.SetConsumerKey(consumerkey)
	anaconda.SetConsumerSecret(consumersecret)
	return anaconda.NewTwitterApi(accesstoken, accesstokensecret)
}

//GetTweet fetches tweet from Twitter
func GetTweet(tweeturl string, api *anaconda.TwitterApi) anaconda.Tweet {

	//get TweetId from url.
	result := strings.SplitAfter(tweeturl, "/status/")

	fmt.Println("TweetId is", result[1])

	TweetID, err := strconv.ParseInt(strings.TrimSpace(result[1]), 10, 64)
	fmt.Println("value of TweetID is", TweetID)
	if err != nil {
		fmt.Println("TweetID error!")
		fmt.Println(err)
	}

	//set options
	options := url.Values{}
	options.Set("id", result[1])

	tweet, err := api.GetTweet(TweetID, options)
	if err != nil {
		fmt.Println("Error in Fetching Tweet")
		fmt.Println(err)
	}
	return tweet
}

//GetTweetURL collect Twitter Tweet URL from user.
func GetTweetURL() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Tweet URL: ")
	url, _ := reader.ReadString('\n')

	//	fmt.Println("url is ==>", url)

	//validate url
	if IsURL(url) {
		err := errors.New("url: Tweet url is not valid")
		fmt.Println(err)
		os.Exit(1)
	}
	return url
}

//IsURL check if url is valid
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	api := TwitterLogin()
	tweeturl := GetTweetURL()
	tweet := GetTweet(tweeturl, api)
	fmt.Println(tweet.Text)
}
