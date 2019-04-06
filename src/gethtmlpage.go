package largeoutput

import (
	"bytes"
	"fmt"
	"github.com/iwdgo/testingfiles"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
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

/* for benchmark purposes, the test is reduced to call the comparison */
func GetHTMLPageString() error {
	testingfiles.OutputDir("output")
	pfileName := "pagegot.html"
	testingfiles.StringToFile(pfileName, getHTMLPage())
	i, _, _, _ := runtime.Caller(0)
	if funcname := strings.SplitAfter(filepath.Base(runtime.FuncForPC(i).Name()), "."); len(funcname) == 1 {
		return fmt.Errorf("Func name not found")
	} else {
		return testingfiles.FileCompare(pfileName, "pagewant.html") // second element is the func name
	}
}

/* Buffer to file, iso String. Then comparing files. No real gain */
func GetHTMLPageBuffer() error {
	testingfiles.OutputDir("output")
	pfileName := "pagegot.html"
	testingfiles.BufferToFile(pfileName, bytes.NewBuffer(getHTMLPage()))
	i, _, _, _ := runtime.Caller(0)
	if funcname := strings.SplitAfter(filepath.Base(runtime.FuncForPC(i).Name()), "."); len(funcname) == 1 {
		return fmt.Errorf("Func name not found")
	} else {
		return testingfiles.FileCompare(pfileName, "pagewant.html") // second element is the func name
	}
}

/* No got file. Comparing buffer to want file. Got file created only if different */
func GetHTMLPageBufferNoGotFile() error {
	testingfiles.OutputDir("output") // for want file
	i, _, _, _ := runtime.Caller(0)
	if funcname := strings.SplitAfter(filepath.Base(runtime.FuncForPC(i).Name()), "."); len(funcname) == 1 {
		return fmt.Errorf("Func name not found")
	} else {
		return testingfiles.BufferCompare(bytes.NewBuffer(getHTMLPage()), "pagewant.html")
	}
}
