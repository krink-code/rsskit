
package main

import (
    "fmt"
    //"flag"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "bytes"
    "os"
    "log"
    "io"
    //"strings"
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

func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
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

func main() {

    db := os.Args[1]
    url := os.Args[2]

    fmt.Println(db)

    writer := WriteToFile("result.txt", "Hello Readers of golangcode.com\n")
    if writer != nil {
        log.Fatal(writer)
    }

    //feed, err := getContent("https://feeds.twit.tv/sn.xml")
    data, err := getContent(url)
    if err != nil {
        fmt.Printf("Failed to get XML: %v \n", err)
        os.Exit(1)
    } /* else {
        fmt.Println("Received XML")
    } */

    //fmt.Println("Received XML: %v", data)

    fmt.Println(string(data))

    //f := strings(feed)
    f := bytes.NewReader(data)
    byteValue, _ := ioutil.ReadAll(f)
    rss := Rss{}
    xml.Unmarshal(byteValue, &rss)
    feed := Feed{}
    xml.Unmarshal(byteValue, &feed)

    fmt.Println("Rss Version: " + rss.Version)
    fmt.Println("Rss Title: " + rss.Title)
    fmt.Println("Rss Link: " + rss.Link)
    fmt.Println("Rss Description: " + rss.Description)

    fmt.Println("Rss.Channel Title: " + rss.Channel.Title)

    fmt.Println("Feed.Title: " + feed.Title)
    fmt.Println("Feed.Link: " + feed.Link)

    for i, item := range rss.Channel.Items {
      fmt.Println(i)
      fmt.Println(item.Title)
      writer := WriteToFile("result.txt", item.Title + "\n")
      if writer != nil {
        log.Fatal(writer)
      }
    }

    for i, entry := range feed.Entries {
      fmt.Println(i)
      fmt.Println(entry.Title)
    }

    /*
    if data, err := getContent("https://feeds.twit.tv/sn.xml"); err != nil {
        fmt.Printf("Failed to get XML: %v", err)
    } else {
        fmt.Println("Received XML:")
        fmt.Println(string(data))
    }
    */

    //byteValue, _ := ioutil.ReadAll(string(data))
    /*
    rss := Rss{}
    xml.Unmarshal(byteValue, &rss)
    fmt.Println("Rss version: " + rss.Version)
    fmt.Println("Rss Title: " + rss.Title)
    fmt.Println("Rss Link: " + rss.Link)
    fmt.Println("Rss Description: " + rss.Description)
    */
    //for i, item := range rss.Channel.Items {
    //    fmt.Println(item.Description)
    //}

	fmt.Printf("done\n")

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
