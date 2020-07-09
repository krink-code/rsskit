package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Item struct {
	XMLName     xml.Name `xml:"item"`

	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`

	Items   []Item   `xml:"channel>item"`
}

func main() {
	file, err := ioutil.ReadFile("file.rss")

 	if err != nil {
 		log.Println(err)
 		return
 	}

	v := Rss{}
	err = xml.Unmarshal(file, &v)
	if err != nil {
        log.Println(err.Error())
        os.Exit(1)
    }

	for _, item := range v.Items {
		log.Println("Title: ", item.Title)
		log.Println("Description: ", item.Description)
		log.Println("Link: ", item.Link)
		log.Println("PubDate: ", item.PubDate)
	}
}


