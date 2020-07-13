
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "os"
    "encoding/json"
    "bytes"
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
   PubDate     string `xml:"pubDate"`
   Guid        string `xml:"guid"`
}

type Feed struct {
   XMLName     xml.Name `xml:"feed"`
   Title       string   `xml:"title"`
   Link        Link     `xml:"link"`
   Entries     []Entry  `xml:"entry"`
}

type Entry struct {
   Name        string `xml:"name,attr"`
   Title       string `xml:"title"`
   Link        Link   `xml:"link"`
   Summary     string `xml:"summary"`
   Updated     string `xml:"updated"`
   Id          string `xml:"id"`
}

type Link struct {
   Href        string `xml:"href,attr"`
}

type Json struct {
   Title       string `json:"title,omitempty"`
   Link        string `json:"link,omitempty"`
   Description string `json:"description,omitempty"`
   PubDate     string `json:"pubdate,omitempty"`
   Summary     string `json:"summary,omitempty"`
   Updated     string `json:"updated,omitempty"`
}


func main() {

    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s url\n", os.Args[0])
        os.Exit(1)
    }

    url := os.Args[1]

    data, err := getContent(url)
    if err != nil {
        fmt.Printf("Failed to get XML: %v \n", err)
        os.Exit(1)
    }

    f := bytes.NewReader(data)
    byteValue, _ := ioutil.ReadAll(f)

    rss := Rss{}
    xml.Unmarshal(byteValue, &rss)

    fmt.Println("{")

    if rss.Channel.Items != nil {
        fmt.Println(`"title": "` + rss.Channel.Title + `",`)
        fmt.Println(`"link": "` + rss.Channel.Link + `",`)
        fmt.Println(`"posts": [`)

        c := len(rss.Channel.Items)
        for _, item := range rss.Channel.Items {
            c--
            jdata := &Json{Title: item.Title, Link: item.Link,
                       Description: item.Description, PubDate: item.PubDate}
            j, _ := json.MarshalIndent(jdata, " ", " ")
            if c == 0 {
                fmt.Println(string(j))
            } else {
                fmt.Println(string(j) + `,`)
            }
        }
        fmt.Println(`]`)
    }

    feed := Feed{}
    xml.Unmarshal(byteValue, &feed)

    if feed.Entries != nil {
        fmt.Println(`"title": "` + feed.Title + `",`)
        fmt.Println(`"link": "` + feed.Link.Href + `",`)
        fmt.Println(`"posts": [`)

        c := len(feed.Entries)
        for _, entry := range feed.Entries {
            jdata := &Json{Title: entry.Title, Link: entry.Link.Href,
                       Summary: entry.Summary, Updated: entry.Updated}
            j, _ := json.MarshalIndent(jdata, " ", " ")
            if c == 0 {
                fmt.Println(string(j))
            } else {
                fmt.Println(string(j) + `,`)
            }
        }
        fmt.Println(`]`)
    }
    fmt.Println("}")

}

