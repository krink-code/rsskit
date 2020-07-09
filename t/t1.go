
package main

import "fmt"
import "encoding/xml"

type Root struct {
	Source Source `xml:"encoding"`
}

type Source struct {
	Data string `xml:",chardata"`
}

func main() {
	x := `<root><source:encoding><![CDATA[Hello]]></source:encoding></root>`
	r := Root {}
	xml.Unmarshal([]byte(x), &r)
	fmt.Println(r)
}

