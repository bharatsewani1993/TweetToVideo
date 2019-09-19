package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
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

//CreateImages create images for video
func CreateImages() {
	backGround, err := os.Open("background.png")
	profilepic, err := os.Open("profilepic.jpg")
	if err != nil {
		fmt.Println(err)
	}

	img1, err := png.Decode(backGround)
	img2, err := jpeg.Decode(profilepic)


	if err != nil {
		fmt.Println(err)
	}

	offset := image.Pt(300, 200)
    b := img1.Bounds()
    canvas := image.NewRGBA(b)
    draw.Draw(canvas, b, img1, image.ZP, draw.Src)
    draw.Draw(canvas, img2.Bounds().Add(offset), img2, image.ZP, draw.Over)

	out, err := os.Create("./reset.jpg")
	if err != nil {
		fmt.Println(err)
	}

	var opt jpeg.Options
	opt.Quality = 80

	jpeg.Encode(out, canvas, &opt) 
}

func main() {
	//	api := TwitterLogin()
	//	tweeturl := GetTweetURL()
	//	tweet := GetTweet(tweeturl, api)
	//fmt.Println(tweet.Text)

	text := "My Name is Bharat and I am a Programmer."
	fmt.Println(text)

	CreateImages()
	fmt.Println("image created")

}
