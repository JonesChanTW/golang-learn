package webserver

import (
	"fmt"
	"net/http"
)

type handl int

func (h handl) ServerDefault(w http.ResponseWriter, req *http.Request) {

}

// OriginHTTPRun http server 使用
func OriginHTTPRun() {
	// var m handl
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data := ""
		switch req.Method {
		case "GET":
			data = fmt.Sprintf("Visit main page")
		case "POST":
			data = fmt.Sprintf("Main page can't add anything\n")
			data += fmt.Sprintf("post data \n")
			if err := req.ParseForm(); err != nil {
				break
			}
			for key, val := range req.Form {
				data += fmt.Sprintf("key %v = %v \n", key, val)
			}
		case "DELETE":
			data = fmt.Sprintf("Main page can't Del anything")
		}
		fmt.Fprintln(w, data)

	}))
	http.ListenAndServe(":80", nil)
}
