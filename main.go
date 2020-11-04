package main

import (
        "net/http"
        "fmt"
        "os"
        "time"
        "encoding/json"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Json_struct struct {
		Json_string string `json:"timezone"`
}
func main() {
	metrics := os.Getenv("METRICS")
	if metrics == "TRUE" {
		//default prometheus handler
		http.Handle("/metrics", promhttp.Handler())
        }
        http.HandleFunc("/timezone", PostHandler)
        http.ListenAndServe(":80", nil)
}
func PostHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		var t Json_struct
        err := decoder.Decode(&t)
        if err != nil {
			panic(err)
        }
        // get time func with loc parameter
        loc, err := time.LoadLocation(t.Json_string)
        fmt.Fprintln(rw, time.Now().In(loc))
     } else {
          fmt.Fprintln(rw, "not a POST request")
       }
}
