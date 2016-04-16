package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//func fooHandler(w http.ResponseWriter, r *http.Request) {
//  fmt.Fprintf(w, "Hello foo, %q", html.EscapeString(r.URL.Path))
//}

const (
	PushoverURL   = "https://api.pushover.net/1/messages.json"
	PushoverToken = "aEKdQe2dTNYx28zgjuVgBNoAMAceBJ"
)

type sms struct {
	id        string
	timeStamp string
	to        string
	from      string
	message   string
}

func smsToPushover(s sms) {
	log.Printf("%v\n", s)

	var message string
	message = s.from + " " + s.message

	resp, err := http.PostForm(PushoverURL,
		url.Values{
			"token":   {PushoverToken},
			"user":    {"u9uFaQmgrGrh6ggG3PPJcURtUHAiUC"},
			"message": {message}})
	if err != nil {
		// handle error
	}
	//defer resp.Body.Close()
	//ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
}

func main() {
	//fmt.Println("hello world")
	//http.Handle("/foo", fooHandler)

	http.HandleFunc("/sms-notify", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Method %v %v\n", r.Method, r.RequestURI)
		if r.Method == "GET" {
			var s sms
			s.id = r.FormValue("ID")
			s.timeStamp = r.FormValue("TIMESTAM")
			s.to = r.FormValue("TO")
			s.from = r.FormValue("FROM")
			s.message = r.FormValue("MESSAGE")
			go smsToPushover(s)

			defer r.Body.Close()
		} else {
			log.Println("Only supports GET")
		}

		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
