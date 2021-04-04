package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
)

type History struct {
	Favicon    string `json:"favicon_url"`
	Transition string `json:"page_transition"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	ClientId   string `json:"client_id"`
	Time       string `json:"time_usec"`
}

type BrowserHistory struct {
	Histories []History `json:"Browser History"`
}

func main() {
	fileFlag := flag.String("f", "", "Target file path")
	flag.Parse()

	if err := parseJson2Text(*fileFlag); err != nil {
		os.Exit(0)
	}

}

func parseJson2Text(filePath string) error {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var bhs BrowserHistory

	json.Unmarshal(raw, &bhs)

	var basename = path.Base(filePath)
	file, err := os.Create(basename + ".txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, history := range bhs.Histories {
		_, err := file.WriteString(history.Url + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
