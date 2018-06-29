package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	auth "./packages/auth"
	req "./packages/req"
	"github.com/unrolled/secure"
)

func Preview_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string

		var result string
		r.ParseForm()
		user = r.Form["user"][0]

		resp, flg := req.PreviewProcess(user)

		if flg {

			for _, temp := range resp {
				result = result + "<article><h1 id=\"head" + temp.ID.Hex() + "\">" + temp.Title + "</h1><img id=\"img" + temp.ID.Hex() + "\" src=\"" + temp.File + "\" alt=\"\" /><p>" + temp.Text[0:300] + "...</p><a href=\"#\" class=\"rm\" onclick=\"more('" + temp.ID.Hex() + "')\">بیشتر</a></article>"
			}
			fmt.Fprintf(w, result)
			return
		} else {
			fmt.Fprintf(w, "0")
			return
		}

	}

}

func login_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string
		var pass string
		var rate string
		var stat string

		r.ParseForm()
		user = r.Form["user"][0]
		pass = r.Form["pass"][0]

		code, flg := auth.LoginProcess(user, pass)

		if flg {

			for i := 0; i < 5; i++ {

				if i < int(code.Rate) {
					rate = rate + "<i class=\"fas fa-star\"></i>"
				} else {
					rate = rate + "<i class=\"far fa-star\"></i>"
				}
			}

			stat = "<p1 onclick=\"showFollowers()\"> دوستان : " + strconv.Itoa(len(code.Followers)) + "</p1><br>" + "<p1 onclick=\"showFollowings()\">دنبال شده : " + strconv.Itoa(len(code.Followings)) + "</p1><br>" + "<p1>" + code.Usertype + "</p1>"

			fmt.Fprintf(w, "1&"+code.Username+"&"+rate+"&"+stat)
			return
		} else {
			fmt.Fprintf(w, "0&"+code.Username+"&"+rate+"&"+stat)
			return
		}

	}

}

func submit_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string
		var pass string
		r.ParseForm()
		user = r.Form["user"][0]
		pass = r.Form["pass"][0]

		code, flg := auth.SubmitProcess(user, pass)

		if !flg {
			fmt.Fprintf(w, "0&"+code.Username+"&"+code.Password)
			return
		} else {
			fmt.Fprintf(w, "1&"+code.Username+"&"+code.Password)
			return
		}

	}

}

func like_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string
		var newsID string
		r.ParseForm()
		user = r.Form["username"][0]
		newsID = r.Form["newsID"][0]

		flg := req.Like(user, newsID)

		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, "1")
			return
		}

	}

}

func dislike_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string
		var newsID string
		r.ParseForm()
		user = r.Form["username"][0]
		newsID = r.Form["newsID"][0]

		flg := req.DisLike(user, newsID)

		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, "1")
			return
		}

	}

}

func more_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var newsID string
		var user string
		r.ParseForm()
		newsID = r.Form["newsID"][0]
		user = r.Form["pressUser"][0]
		log.Print(user)
		result, flg := req.More(newsID)

		if !flg {

			fmt.Fprintf(w, "0")
			return
		} else {

			heart := "<i class=\"far fa-heart\"></i>"
			toggle := "like"
			for i, _ := range result.Likes {
				if result.Likes[i] == user {
					heart = "<i class=\"fas fa-heart\"></i>"
					toggle = "dislike"
				}
			}
			resp := "<button id=\"clsCm\" onclick=\"clsCm()\">X</button><img src=\"" + result.File + "\"><br><div id=\"newsStat\"><p id=\"likes\" onclick=\"" + toggle + "('" + result.ID.Hex() + "')\">" + heart + "</p><p id=\"count\">" + strconv.Itoa(len(result.Likes)) + "</p><p id=\"edit\" onclick=\"edit('" + result.ID.Hex() + "')\"><i class=\"far fa-edit\"></i>   ویرایش  </p><p id=\"date\">" + result.Date + "</p></div><input id=\"inputCm\" placeholder=\"  نظر بدهید...\" ><button id=\"sendCm\" onclick=\"sendCm('" + result.ID.Hex() + "')\">ارسال</button><div id=\"cms\">"
			for _, temp := range result.Comments {

				resp = resp + "<div class=\"cm\"><p>" + strings.Split(temp, "@")[0] + ":<p>" + "<p>" + strings.Split(temp, "@")[1] + "</p>" + "</div>"
			}

			resp = resp + "</div><div class=\"textCm\"><p2>@" + result.Author + ":</p2><br>" + result.Text + "</div>"

			fmt.Fprintf(w, resp)
			return
		}

	}

}

func cm_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string
		var newsID string
		var cm string
		r.ParseForm()
		user = r.Form["username"][0]
		cm = r.Form["Cm"][0]
		newsID = r.Form["newsID"][0]
		flg := req.Cm(newsID, user, cm)

		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			res := "<p>" + user + ":<p>" + "<p>" + cm + "</p>"
			fmt.Fprintf(w, res)
			return
		}

	}

}

func upload_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		// Post
		r.ParseForm()
		file, handler, err := r.FormFile("image")
		filename := r.Form["name"][0]
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./files/"+filename+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		return

	}
}

func sendNews_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		r.ParseForm()
		user := r.Form["username"][0]
		title := r.Form["title"][0]
		keywords := r.Form["keywords"][0]
		time := r.Form["time"][0]
		text := r.Form["text"][0]

		flg, id := req.SendNews(user, title, keywords, time, text)

		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, id)
			return
		}

	}

}

func searchNews_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var result string

	if r.Method == "POST" {

		r.ParseForm()
		category := r.Form["cat"][0]
		query := r.Form["query"][0]

		resp, flg := req.SearchNews(category, query)

		if flg {

			for _, temp := range resp {
				result = result + "<article><h1>" + temp.Title + "</h1><img src=\"" + temp.File + "\" alt=\"\" /><p>" + temp.Text[0:300] + "...</p><a href=\"#\" class=\"rm\" onclick=\"more('" + temp.ID.Hex() + "')\">بیشتر</a></article>"
			}
			fmt.Fprintf(w, result)
			return
		} else {
			fmt.Fprintf(w, "0")
			return
		}

	}

}

func editNews_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		r.ParseForm()
		user := r.Form["username"][0]
		title := r.Form["title"][0]
		keywords := r.Form["keywords"][0]
		time := r.Form["time"][0]
		text := r.Form["text"][0]
		newsID := r.Form["id"][0]

		flg := req.ModifyNews(user, title, keywords, time, text, newsID)

		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, newsID)
			return
		}

	}

}

func getFollowers_ctrl(w http.ResponseWriter, r *http.Request) {

	var resp string
	var userrate string
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		r.ParseForm()
		user := r.Form["username"][0]
		//pass := r.Form["password"][0]

		list, flg := req.Followers(user)

		for _, tmp := range list {
			for i := 0; i < 5; i++ {
				if i < int(tmp.Rate) {
					userrate = userrate + "<i class=\"fas fa-star\"></i>"
				} else {
					userrate = userrate + "<i class=\"far fa-star\"></i>"
				}
			}
			resp = resp + "<div class=\"contact\"><p3>" + tmp.Username + "</p3><div id=\"frate\">" + userrate + "</div><p3 onclick=\"latest('" + tmp.Username + "')\">مطالب اخیر</p3><p3 onclick=\"unfollow('" + tmp.Username + "')\">دنبال نکردن</p3></div><br>"
			userrate = ""
		}
		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, resp)
			return
		}

	}

}

func getFollowings_ctrl(w http.ResponseWriter, r *http.Request) {

	var resp string
	var userrate string
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		r.ParseForm()
		user := r.Form["username"][0]
		//pass := r.Form["password"][0]

		list, flg := req.Followings(user)

		for _, tmp := range list {
			for i := 0; i < 5; i++ {
				if i < int(tmp.Rate) {
					userrate = userrate + "<i class=\"fas fa-star\"></i>"
				} else {
					userrate = userrate + "<i class=\"far fa-star\"></i>"
				}
			}
			resp = resp + "<div class=\"contact\"><p3>" + tmp.Username + "</p3><div id=\"frate\">" + userrate + "</div><p3 onclick=\"latest('" + tmp.Username + "')\">مطالب اخیر</p3><p3 onclick=\"unfollow('" + tmp.Username + "')\">دنبال نکردن</p3></div><br>"
			userrate = ""
		}
		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, resp)
			return
		}

	}

}

func unfollow_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		r.ParseForm()
		user := r.Form["username"][0]
		//pass := r.Form["password"][0]
		unflw := r.Form["unflw"][0]

		flg := req.Unfollow(user, unflw)

		if !flg {
			fmt.Fprintf(w, "0")
			return
		} else {
			fmt.Fprintf(w, "1")
			return
		}

	}

}

func main() {

	getNews_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	login_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	submit_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	like_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	dislike_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	more_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	cm_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	upload_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	sendNews_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	searchNews_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	editNews_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	getFollowers_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	getFollowings_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	unfollow_rout := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	app1 := getNews_rout.Handler(http.HandlerFunc(Preview_ctrl))
	app2 := login_rout.Handler(http.HandlerFunc(login_ctrl))
	app3 := submit_rout.Handler(http.HandlerFunc(submit_ctrl))
	app4 := like_rout.Handler(http.HandlerFunc(like_ctrl))
	app5 := dislike_rout.Handler(http.HandlerFunc(dislike_ctrl))
	app6 := more_rout.Handler(http.HandlerFunc(more_ctrl))
	app7 := cm_rout.Handler(http.HandlerFunc(cm_ctrl))
	app8 := upload_rout.Handler(http.HandlerFunc(upload_ctrl))
	app9 := sendNews_rout.Handler(http.HandlerFunc(sendNews_ctrl))
	app10 := searchNews_rout.Handler(http.HandlerFunc(searchNews_ctrl))
	app11 := editNews_rout.Handler(http.HandlerFunc(editNews_ctrl))
	app12 := getFollowers_rout.Handler(http.HandlerFunc(getFollowers_ctrl))
	app13 := getFollowings_rout.Handler(http.HandlerFunc(getFollowings_ctrl))
	app14 := unfollow_rout.Handler(http.HandlerFunc(unfollow_ctrl))

	http.Handle("/getnews", app1)
	http.Handle("/login", app2)
	http.Handle("/submit", app3)
	http.Handle("/like", app4)
	http.Handle("/dislike", app5)
	http.Handle("/more", app6)
	http.Handle("/Cm", app7)
	http.Handle("/upload", app8)
	http.Handle("/sendNews", app9)
	http.Handle("/searchNews", app10)
	http.Handle("/editNews", app11)
	http.Handle("/getFollowers", app12)
	http.Handle("/getFollowings", app13)
	http.Handle("/unfollow", app14)

	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))

}
