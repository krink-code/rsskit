
package main

import (
    "encoding/xml"
    "fmt"
    "net/http"
    "os"
)

type Enclosure struct {
    Url    string `xml:"url,attr"`
    Length int64  `xml:"length,attr"`
    Type   string `xml:"type,attr"`
}

type Item struct {
    Title     string    `xml:"title"`
    Link      string    `xml:"link"`
    Desc      string    `xml:"description"`
    Guid      string    `xml:"guid"`
    Enclosure Enclosure `xml:"enclosure"`
    PubDate   string    `xml:"pubDate"`
}

type Channel struct {
    Title string `xml:"title"`
    Link  string `xml:"link"`
    Desc  string `xml:"description"`
    Items []Item `xml:"item"`
}

type Rss struct {
    Channel Channel `xml:"channel"`
}

func main() {

    url := os.Args[1]

    //resp, err := http.Get("https://feeds.twit.tv/sn.xml")
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error GET: %v\n", err)
        return
    }
    defer resp.Body.Close()

    rss := Rss{}

    decoder := xml.NewDecoder(resp.Body)
    err = decoder.Decode(&rss)
    if err != nil {
        fmt.Printf("Error Decode: %v\n", err)
        return
    }

    fmt.Printf("Channel title: %v\n", rss.Channel.Title)
    fmt.Printf("Channel link: %v\n", rss.Channel.Link)

    for i, item := range rss.Channel.Items {
        fmt.Printf("%v. item title: %v\n", i, item.Title)
    }
}

