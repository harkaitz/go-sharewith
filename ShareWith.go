package sharewith

import (
	"net/url"
	"html/template"
	_ "embed"
)

func Facebook(u string, e ...string) template.URL {
	return template.URL("https://www.facebook.com/sharer/sharer.php?u=" + url.QueryEscape(u))
}

func Twitter(u string, e ...string) template.URL {
	r := "https://twitter.com/intent/tweet?" + "url=" + url.QueryEscape(u)
	if len(e) >= 1 {
		r += "&text=" + url.QueryEscape(e[0])
	}
	if len(e) >= 2 {
		r += "&hashtags=" + url.QueryEscape(e[1])
	}
	return template.URL(r)
}

func WhatsApp(u string, e ...string) template.URL {
	return template.URL("https://api.whatsapp.com/send?text=" + url.QueryEscape(u))
}

func Telegram(u string, e ...string) template.URL {
	r := "https://telegram.me/share/url?url=" + url.QueryEscape(u)
	if len(e) >= 1 {
		r += "&text=" + url.QueryEscape(e[0])
	}
	return template.URL(r)
}

func Instagram(u string, e ...string) template.URL {
	return template.URL("https://www.instagram.com/?url=" + url.QueryEscape(u))
}

func FacebookImgURL()  template.URL { return template.URL(imgpre + FacebookB64PNG)  }
func TwitterImgURL()   template.URL { return template.URL(imgpre + TwitterB64PNG)   }
func WhatsAppImgURL()  template.URL { return template.URL(imgpre + WhatsappB64PNG)  }
func TelegramImgURL()  template.URL { return template.URL(imgpre + TelegramB64PNG)  }
func InstagramImgURL() template.URL { return template.URL(imgpre + InstagramB64PNG) }

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

var imgpre string = "data:image/png;charset=utf-8;base64,"

//go:embed gen/icon-orig-facebook.b64
var FacebookB64PNG string

//go:embed gen/icon-orig-telegram.b64
var TelegramB64PNG string

//go:embed gen/icon-orig-whatsapp.b64
var WhatsappB64PNG string

//go:embed gen/icon-orig-instagram.b64
var InstagramB64PNG string

//go:embed gen/icon-orig-twitter.b64
var TwitterB64PNG string
