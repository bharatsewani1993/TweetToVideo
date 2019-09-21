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
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"github.com/ChimeraCoder/anaconda"
	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
    "golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
//	"reflect"
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


type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

//Write text on image
func AddLabel(img *image.RGBA, x, y int, label string, fontface *basicfont.Face, col color.RGBA) {   
    point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

    d := &font.Drawer{
        Dst:  img,
        Src:  image.NewUniform(col),
        Face: fontface,
        Dot:  point,
    }
    d.DrawString(label)
}

//CreateImages create images for video
func CreateImages() {

	//Resize user profile pic to 50 X 50
	src, err := imaging.Open("1.jpg")
	if err != nil{
	  log.Fatal("Error %v", err)
	}

	src = imaging.Fill(src,50,50, imaging.Center, imaging.Lanczos)
	err = imaging.Save(src,"profilepic.jpg")
	if err != nil{
	  log.Fatal("Error %v",err)
	}
 
	//Convert profile pic to circle
	//open profile pic
	profilepic, err := os.Open("profilepic.jpg")
	if err != nil {
		fmt.Println(err)
	}
	profilepicd, err := jpeg.Decode(profilepic)

	if err != nil {
		fmt.Println(err)
	}

	//create canvas for profile pic
	profilepiccanvas := image.NewRGBA(image.Rect(0,0,50,50))
	p := image.Pt(25,25)
	//draw mask on profile pic
	draw.DrawMask(profilepiccanvas, profilepiccanvas.Bounds(), profilepicd, image.ZP, &circle{p, 25}, image.ZP, draw.Over)

	//save as circle cropped 
	pout, err := os.Create("./roundpic.png")
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(pout, profilepiccanvas) 


	//Merge profile pic on white background
	backGround, err := os.Open("background.png")
	roundpic, err := os.Open("roundpic.png")
	backGroundd, err := png.Decode(backGround)
	roundpicd, err := png.Decode(roundpic)

	if err != nil {
		fmt.Println(err)
	}

	//Draw profilepic on white backgroud
	offset := image.Pt(300, 200)
    b := backGroundd.Bounds()
    canvas := image.NewRGBA(b)
    draw.Draw(canvas, b, backGroundd, image.ZP, draw.Src)
    draw.Draw(canvas, roundpicd.Bounds().Add(offset), roundpicd, image.ZP, draw.Over)

	//Write text on Image	

	s := make([]string, 12)
    s[0] = "a"
    s[1] = "b"
	s[2] = "c"
	s[3] = "d"
    s[4] = "e"
	s[5] = "f"
	s[6] = "g"
    s[7] = "h"
	s[8] = "i"
	s[9] = "j"
    s[10] = "k"
	s[11] = "l"
	
		//set font color
	/*	col := color.RGBA{0, 0, 0, 255}
		AddLabel(canvas, 360, 215, "Bharat Sewani",inconsolata.Bold8x16,col)
		col = color.RGBA{0, 0, 0, 150}
		AddLabel(canvas, 360, 230, "@bharatsewani199",inconsolata.Regular8x16,col)
		col = color.RGBA{0, 0, 0, 255}
		AddLabel(canvas, 310, 270, "No matter how far you have gone on the wrong road,",inconsolata.Bold8x16,col)
		AddLabel(canvas, 310, 290, "You can still turn around.",inconsolata.Bold8x16,col)
		col = color.RGBA{27,149,224,255}
		AddLabel(canvas, 310, 310, "#TuesdayMotivation",inconsolata.Bold8x16,col)
		col = color.RGBA{0, 0, 0, 150}
		AddLabel(canvas, 310, 340, "9:00 AM . Jul 31, 2018",inconsolata.Regular8x16,col) */

	//	fmt.Println(reflect.TypeOf(inconsolata.Regular8x16))
		
	for i := 0; i < 10; i++ {
		col := color.RGBA{0, 0, 0, 255}
		AddLabel(canvas, 360 + i*10, 215, s[i],inconsolata.Bold8x16,col)
		//save frame
		s := strconv.Itoa(i)
		out, err := os.Create("./frame"+s+".jpg")
		if err != nil {
			fmt.Println(err)
		}
		var opt jpeg.Options
		opt.Quality = 80
		jpeg.Encode(out, canvas, &opt) 
	}

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
