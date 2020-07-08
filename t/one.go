
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
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


func main() {

if data, err := getContent("https://feeds.twit.tv/sn.xml"); err != nil {
    fmt.Printf("Failed to get XML: %v", err)
} else {
    fmt.Println("Received XML:")
    fmt.Println(string(data))
}

	fmt.Printf("done\n")

}



