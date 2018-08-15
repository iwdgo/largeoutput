package largeoutput

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/*
Write a program to download the contents of http://www.{tech}.com/belgique/index.html (the {tech} homepage for Belgium),
and then save the contents of the page to a new local file, with all occurrences of "{tech}" replaced by "MyTech".
*/

const (
	techName  = "SAP"
	myTech    = "MyTech"
	myCountry = "belgique"
)

func getTechHomePage() []byte {
	resp, err := http.Get("http://www." + techName + ".com/" + myCountry + "/index.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
}

func getHTMLPage() []byte {
	//It is assume that replacement is case sensitive
	return bytes.Replace(getTechHomePage(), []byte(techName), []byte(myTech), -1)
}
