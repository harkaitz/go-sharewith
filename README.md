GO SHAREWITH
============

## Go programs

    Usage: sharewith all|PLATFORM img|URL [TEXT] [HASHTAG]
    
    Print a share with PLATFORM URL or platform's logo as "data:" to be
    embedded in HTML files.
    
    Supported platforms are: Facebook, Twitter, WhatsApp, Telegram,
    Instagram.
    
    Copyright (c) 2024 Harkaitz Agirre, harkaitz.aguirre@gmail.com

## Go documentation

    package sharewith // import "github.com/harkaitz/go-sharewith"
    
    func A(u, txt, hashtag string) template.HTML
    func Facebook(uri string, e ...string) template.URL
    func FacebookImgURL() template.URL
    func Instagram(uri string, e ...string) template.URL
    func InstagramImgURL() template.URL
    func Telegram(uri string, e ...string) template.URL
    func TelegramImgURL() template.URL
    func Twitter(uri string, e ...string) template.URL
    func TwitterImgURL() template.URL
    func WhatsApp(uri string, e ...string) template.URL
    func WhatsAppImgURL() template.URL

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
