package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	auth "./packages/auth"
	req "./packages/req"
	"github.com/unrolled/secure"
)

func Preview_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {

		var user string
		var key string
		var result string
		r.ParseForm()
		user = r.Form["user"][0]
		key = r.Form["key"][0]

		resp, flg := req.PreviewProcess(user)

		if key != "0" {

			for _, temp := range resp {
				result = result + "<div class=\"press\" id=\"" + temp.ID.Hex() + "\"><div class=\"pressbar\" id=\"pressbar-" + temp.ID.Hex() + "\"><form class=\"updown\" id=\"updown-" + temp.ID.Hex() + "\"><div class=\"value-button\"  id=\"decrease\" onclick=\"decreaseValue('like-" + temp.ID.Hex() + "')\" value=\"Decrease Value\"></div><p class=\"like\" id=\"like-" + temp.ID.Hex() + "\">+" + strconv.Itoa(len(temp.Likes)) + "</p><div class=\"value-button\" id=\"increase\" onclick=\"increaseValue('like-" + temp.ID.Hex() + "')\" value=\"Increase Value\"></div></form><div class=\"value-button\" id=\"share\" onclick=\"getlink()\" value=\"id-" + temp.ID.Hex() + "\"></div><div class=\"a1\" style=\"color:rgb(251, 255, 0);\">" + strconv.Itoa(len(temp.Comments)) + " نظر </div><div class=\"a1\" style=\"color:orangered;\">" + temp.Date + "  " + temp.Time + " </div> <div class=\"editbtn\" ></div>	</div><div class=\"content\" id=\"content-" + temp.ID.Hex() + "\"><div class=\"matn\" id=\"matn-" + temp.ID.Hex() + "\"><p1>:" + temp.Author + "</p1><p1>" + temp.Text + "</div><div class=\"ax\" id=\"ax-" + temp.ID.Hex() + "\" style=\"background-image:url(" + temp.File + ")\" ></div></div></div>"
			}

		} else {

			for _, temp := range resp {
				result = result + "<div class=\"press\" id=\"" + temp.ID.Hex() + "\"><div class=\"pressbar\" id=\"pressbar-" + temp.ID.Hex() + "\"><form class=\"updown\" id=\"updown-" + temp.ID.Hex() + "\"><div class=\"value-button\"  id=\"decrease\" onclick=\"decreaseValue('like-" + temp.ID.Hex() + "')\" value=\"Decrease Value\"></div><p class=\"like\" id=\"like-" + temp.ID.Hex() + "\">+" + strconv.Itoa(len(temp.Likes)) + "</p><div class=\"value-button\" id=\"increase\" onclick=\"increaseValue('like-" + temp.ID.Hex() + "')\" value=\"Increase Value\"></div></form><div class=\"value-button\" id=\"share\" onclick=\"getlink()\" value=\"id-" + temp.ID.Hex() + "\"></div><div class=\"a1\" style=\"color:rgb(251, 255, 0);\">" + strconv.Itoa(len(temp.Comments)) + " نظر </div><div class=\"a1\" style=\"color:orangered;\">" + temp.Date + "  " + temp.Time + " </div>	</div><div class=\"content\" id=\"content-" + temp.ID.Hex() + "\"><div class=\"matn\" id=\"matn-" + temp.ID.Hex() + "\"><p1>:" + temp.Author + "</p1><p1>" + temp.Text + "</div><div class=\"ax\" id=\"ax-" + temp.ID.Hex() + "\" style=\"background-image:url(" + temp.File + ")\" ></div></div></div>"
			}

		}

		if flg {
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

		r.ParseForm()
		user = r.Form["user"][0]
		pass = r.Form["pass"][0]

		code, flg := auth.LoginProcess(user, pass)

		if flg {
			fmt.Fprintf(w, "1&"+code.Username)
			return
		} else {
			fmt.Fprintf(w, "0&"+code.Username)
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

	app1 := getNews_rout.Handler(http.HandlerFunc(Preview_ctrl))
	app2 := login_rout.Handler(http.HandlerFunc(login_ctrl))
	app3 := submit_rout.Handler(http.HandlerFunc(submit_ctrl))
	app4 := like_rout.Handler(http.HandlerFunc(like_ctrl))
	app5 := dislike_rout.Handler(http.HandlerFunc(dislike_ctrl))

	http.Handle("/getnews", app1)
	http.Handle("/login", app2)
	http.Handle("/submit", app3)
	http.Handle("/like", app4)
	http.Handle("/dislike", app5)

	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))

}
