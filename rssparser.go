
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "bytes"
    "os"
    "strings"
    "time"
    "encoding/json"
)

func getContent(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Read body: %v", err)
    }

    return data, nil
}

type Rss struct {
   XMLName     xml.Name `xml:"rss"`
   Version     string   `xml:"version,attr"`
   Channel     Channel  `xml:"channel"`
   Description string   `xml:"description"`
   Title       string   `xml:"title"`
   Link        string   `xml:"link"`
}

type Channel struct {
   XMLName     xml.Name `xml:"channel"`
   Title       string   `xml:"title"`
   Link        string   `xml:"link"`
   Description string   `xml:"description"`
   Items       []Item   `xml:"item"`
}

type Item struct {
   Name        string `xml:"name,attr"`
   Title       string `xml:"title"`
   Link        string `xml:"link"`
   Description string `xml:"description"`
   PubDate     string `xml:"pubdate"`
   Guid        string `xml:"guid"`
}

type Feed struct {
   XMLName     xml.Name `xml:"feed"`
   Title       string   `xml:"title"`
   Link        string   `xml:"link"`
   Entries     []Entry  `xml:"entry"`
}

type Entry struct {
   Name        string `xml:"name,attr"`
   Title       string `xml:"title"`
   Link        string `xml:"link"`
   Summary     string `xml:"summary"`
   PubDate     string `xml:"pubdate"`
   Guid        string `xml:"guid"`
}

type Json struct {
   Title      string `json:"title"`
   Link       string `json:"link"`
}

func main() {

    if len(os.Args) != 3 {
        fmt.Printf("Usage: %s db url\n", os.Args[0])
        os.Exit(1)
    }

    db := os.Args[1]
    url := os.Args[2]

    file, err := os.OpenFile(db,
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    data, err := getContent(url)
    if err != nil {
        fmt.Printf("Failed to get XML: %v \n", err)
        os.Exit(1)
    }

    f := bytes.NewReader(data)
    byteValue, _ := ioutil.ReadAll(f)
    rss := Rss{}
    xml.Unmarshal(byteValue, &rss)
    feed := Feed{}
    xml.Unmarshal(byteValue, &feed)

    /*
    if rss.Channel.Items != nil {
        fmt.Println("Rss Version: " + rss.Version)
    }
    */

    b, err := ioutil.ReadFile(db)
    if err != nil {
        panic(err)
    }
    s := string(b)

    t := time.Now()

    for _, item := range rss.Channel.Items {
      if ! strings.Contains(s, item.Title) {
          jdata := &Json{Title: item.Title, Link: item.Link }
          j, _ := json.MarshalIndent(jdata, " ", " ")
          fmt.Println(string(j))
          file.WriteString(item.Title + "|" + t.Format(time.RFC3339) + "\n")
      }
    }

    for _, entry := range feed.Entries {
      if ! strings.Contains(s, entry.Title) {
          jdata := &Json{Title: entry.Title, Link: entry.Link }
          j, _ := json.MarshalIndent(jdata, " ", " ")
          fmt.Println(string(j))
          file.WriteString(entry.Title + "|" + t.Format(time.RFC3339) + "\n")
      }
    }

    file.Close()
	//fmt.Printf("done\n")

}

/*
 <?xml version="1.0" encoding="UTF-8"?>
  <feed xmlns="http://www.w3.org/2005/Atom"><title>
*/

/*
<?xml version="1.0" encoding="utf-8" ?>
<rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:creativeCommons="http://backend.userland.com/creativeCommonsRssModule" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:rawvoice="http://www.rawvoice.com/rawvoiceRssModule/" version="2.0">
	<channel>
		<title>
*/

