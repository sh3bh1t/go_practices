package main

import (
	"fmt"
	"net/url"
)

func main() {

	myurl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"
	parsedurl, err := url.Parse(myurl) // to parse the provided url from string to an url object
	if err != nil {
		fmt.Println("the error while parsing is : ", err)
	}

	// printing certain aspects of the provided url
	fmt.Println("scheme of url : ", parsedurl.Scheme)
	fmt.Println("host of url : ", parsedurl.Host)
	fmt.Println("path of url : ", parsedurl.Path)
	fmt.Println("raw query of url : ", parsedurl.RawQuery)

	//modifying the given url or some of its aspects

	parsedurl.Path = "/checkout"
	parsedurl.RawQuery = "username=kattua"
	newURL := parsedurl.String()
	fmt.Println("the modified url is : ", newURL)

}
