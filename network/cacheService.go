package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"

  "github.com/patrickmn/go-cache"
)

var pageCache = cache.New(5*time.Minute, 10*time.Minute)

func cacheService(w http.ResponseWriter, req *http.Request) {
  var siteUrl string
  
  if(req.Host == "localhost:8000" ) {
    fmt.Println("Calling http")
	siteUrl = "http://godoc.org"
  } else {  
    fmt.Println("Calling https")
	siteUrl = "https://godoc.org" 
  }
  
  url := siteUrl + req.URL.Path
  cachedResponse, found := pageCache.Get(req.URL.Path)

  if found {
	fmt.Println("Cached-site")
    fmt.Fprintf(w, cachedResponse.(string))
  } else {
	fmt.Println("Direct-site")
    siteResponse, err := http.Get(url)

    if err != nil {
      panic(err)
    }

    defer siteResponse.Body.Close()

    siteResponseBody, err := ioutil.ReadAll(siteResponse.Body)
    if err != nil {
      log.Fatal(err)
    }
    responseBody := string(siteResponseBody)

    pageCache.Set(req.URL.Path, responseBody, cache.DefaultExpiration)

    fmt.Fprintf(w, responseBody)
  }

}

func main() {
	http.HandleFunc("/github.com/stretchr/testify/assert", cacheService)
	go http.ListenAndServe(":8000", nil)
	err1 := http.ListenAndServeTLS(":8090", "host.cert", "host.key", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}