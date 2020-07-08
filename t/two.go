
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "bytes"
    "os"
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

type Rss struct {
   XMLName     xml.Name `xml:"rss"`
   Version     string   `xml:"version,attr"`
   Description string   `xml:"description"`
   Title       string   `xml:"title"`
   Link        string   `xml:"link"`
}

func main() {

    url := os.Args[1]

    //feed, err := getContent("https://feeds.twit.tv/sn.xml")
    feed, err := getContent(url)
    if err != nil {
        fmt.Printf("Failed to get XML: %v \n", err)
        return
    } /* else {
        fmt.Println("Received XML")
    } */

    //fmt.Println("Received XML: %v", data)

    fmt.Println(string(feed))

    //f := strings(feed)
    f := bytes.NewReader(feed)
    byteValue, _ := ioutil.ReadAll(f)
    rss := Rss{}
    xml.Unmarshal(byteValue, &rss)
    fmt.Println("Rss Version: " + rss.Version)
    fmt.Println("Rss Title: " + rss.Title)
    fmt.Println("Rss Link: " + rss.Link)
    fmt.Println("Rss Description: " + rss.Description)

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


