
package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

type PostData struct {
	Channel struct {
		Posts []struct {
			Title       string   `xml:"title"`
			Link        string   `xml:"link"`
			Category    []string `xml:"category"`
			Creator     string   `xml:"creator"`
			PubDate     string   `xml:"pubDate"`
			Updated     string   `xml:"updated"`
			License     string   `xml:"license"`
			Encoded     string   `xml:"encoded"`
			Description string   `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func main() {

    resp, err := http.Get(os.Args[1])

    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    data := &PostData{}
    err = xml.Unmarshal(body, data)
    for _, item := range data.Channel.Posts {
        fmt.Println(item.Title) 
    }

}


