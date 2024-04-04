package main
import (
	"fmt"
	"os"
	"github.com/harkaitz/go-sharewith"
)

var HELP string =
`Usage: sharewith all|PLATFORM img|URL [TEXT] [HASHTAG]

Print a share with PLATFORM URL or platform's logo as "data:" to be
embedded in HTML files.

Supported platforms are: Facebook, Twitter, WhatsApp, Telegram,
Instagram.

Copyright (c) 2024 Harkaitz Agirre, harkaitz.aguirre@gmail.com`


func main() {
	var platform, url, text, hashtag string
	
	if len(os.Args) < 3 {
		fmt.Println(HELP)
		os.Exit(1)
	}
	platform = os.Args[1]
	url = os.Args[2]
	if len(os.Args) > 3 { text = os.Args[3] }
	if len(os.Args) > 4 { hashtag = os.Args[4] }
	
	if url == "img" {
		switch platform {
			case "Facebook":  fmt.Println(sharewith.FacebookImgURL())
			case "Twitter":   fmt.Println(sharewith.TwitterImgURL())
			case "WhatsApp":  fmt.Println(sharewith.WhatsAppImgURL())
			case "Telegram":  fmt.Println(sharewith.TelegramImgURL())
			case "Instagram": fmt.Println(sharewith.InstagramImgURL())
			default: fmt.Println("Unknown platform:", os.Args[1])
		}
		os.Exit(0)
	}
	
	if platform == "all" {
		fmt.Println(string(sharewith.A(url, text, hashtag)))
		os.Exit(0)
	}
	
	switch platform {
	case "Facebook":  fmt.Println(sharewith.Facebook(url, text, hashtag))
	case "Twitter":   fmt.Println(sharewith.Twitter(url, text, hashtag))
	case "WhatsApp":  fmt.Println(sharewith.WhatsApp(url, text, hashtag))
	case "Telegram":  fmt.Println(sharewith.Telegram(url, text, hashtag))
	case "Instagram": fmt.Println(sharewith.Instagram(url, text, hashtag))
	default: fmt.Println("Unknown platform:", os.Args[1])
	}
	
	os.Exit(0)
}
