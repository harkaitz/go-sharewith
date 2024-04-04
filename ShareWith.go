package sharewith

import (
	"net/url"
	"html/template"
	_ "embed" /**/
)

/* ---------------------------------------------------------------- */
var imgpre string = "data:image/png;charset=utf-8;base64,"

// Facebook returns a share with facebook URL.
func Facebook(uri string, e ...string) template.URL {
	return template.URL("https://www.facebook.com/sharer/sharer.php?u=" + url.QueryEscape(uri))
}

// Twitter returns a share in twitter URL. You can specify optionally
// a "text" (second arg) or a hashtag (third argument).
func Twitter(uri string, e ...string) template.URL {
	r := "https://twitter.com/intent/tweet?" + "url=" + url.QueryEscape(uri)
	if len(e) >= 1 && e[0] != "" {
		r += "&text=" + url.QueryEscape(e[0])
	}
	if len(e) >= 2 && e[1] != "" {
		r += "&hashtags=" + url.QueryEscape(e[1])
	}
	return template.URL(r)
}

// WhatsApp returns a share with facebook URL.
func WhatsApp(uri string, e ...string) template.URL {
	return template.URL("https://api.whatsapp.com/send?text=" + url.QueryEscape(uri))
}

// Telegram returns a share with telegram URL, you can specify a text
// as second argument optionally.
func Telegram(uri string, e ...string) template.URL {
	r := "https://telegram.me/share/url?url=" + url.QueryEscape(uri)
	if len(e) >= 1 && e[0] != "" {
		r += "&text=" + url.QueryEscape(e[0])
	}
	return template.URL(r)
}

// Instagram returns a share with instagram URL.
func Instagram(uri string, e ...string) template.URL {
	return template.URL("https://www.instagram.com/?url=" + url.QueryEscape(uri))
}

/* ---------------------------------------------------------------- */

// FacebookImgURL returns a "data:" URL with a facebook logo.
func FacebookImgURL()  template.URL { return template.URL(imgpre + facebookB64PNG)  }

// TwitterImgURL returns a "data:" URL with a Twitter logo.
func TwitterImgURL()   template.URL { return template.URL(imgpre + twitterB64PNG)   }

// WhatsAppImgURL returns a "data:" URL with a WhatsApp logo.
func WhatsAppImgURL()  template.URL { return template.URL(imgpre + whatsappB64PNG)  }

// TelegramImgURL returns a "data:" URL with a Telegram logo.
func TelegramImgURL()  template.URL { return template.URL(imgpre + telegramB64PNG)  }

// InstagramImgURL returns a "data:" URL with a Instagram logo.
func InstagramImgURL() template.URL { return template.URL(imgpre + instagramB64PNG) }

// A returns HTML code for a list of sharewith buttons.
func A(u, txt, hashtag string) template.HTML {
	var html string
	
	for i, f := range []struct {
		name string
		uri  func (u string, e ... string) template.URL
		img  func () template.URL
	} {
		{"Facebook"  , Facebook,  FacebookImgURL  },
		{"Twitter"   , Twitter ,  TwitterImgURL   },
		{"WhatsApp"  , WhatsApp,  WhatsAppImgURL  },
		{"Telegram"  , Telegram,  TelegramImgURL  },
		{"Instagram" , Instagram, InstagramImgURL },
	} {
		if i != 0 {
			html += "&nbsp;"
		}
		html += `<a href="` + string(f.uri(u, txt, hashtag)) + `">`
		html += `<img src="` + string(f.img()) + `" alt="` + f.name +`" style="width:32px;"/>`
		html += `</a>`
	}
	
	return template.HTML(html)
}

/* ---------------------------------------------------------------- */

//go:embed gen/icon-orig-facebook.b64
var facebookB64PNG string

//go:embed gen/icon-orig-telegram.b64
var telegramB64PNG string

//go:embed gen/icon-orig-whatsapp.b64
var whatsappB64PNG string

//go:embed gen/icon-orig-instagram.b64
var instagramB64PNG string

//go:embed gen/icon-orig-twitter.b64
var twitterB64PNG string
