package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
)

type TakeoutHistory struct {
	Favicon    string `json:"favicon_url"`
	Transition string `json:"page_transition"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	ClientId   string `json:"client_id"`
	Time       string `json:"time_usec"`
}

type BrowserHistory struct {
	Histories []TakeoutHistory `json:"Browser History"`
}

type ChromeAPIHistory struct {
	Id                     string `json:"id"`
	LastVisitTime          string `json:"lastVisitTime"`
	LastVisitTimeTimestamp string `json:"lastVisitTimeTimestamp"`
	Title                  string `json:"title"`
	TypedCount             string `json:"typedCount"`
	Url                    string `json:"url"`
	VisitCount             string `json:"visitCount"`
}

type ChromeHistory struct {
	Histories []ChromeAPIHistory
}

func main() {
	// fileFlag := flag.String("f", "", "Target file path")
	typeFlag := flag.String("t", "", "Target file type (e: Exported by extension, t: Exported by Google Takeout page)")
	flag.Parse()
	filename := flag.Arg(0)

	if err := extractUrlFromJson(filename, *typeFlag); err != nil {
		print(err)
		os.Exit(0)
	}

}

func extractUrlFromJson(filePath string, fileType string) error {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	switch fileType {
	case "e":
		var chs []ChromeAPIHistory
		json.Unmarshal(raw, &chs)
		var basename = path.Base(filePath)
		file, err := os.Create(basename + ".txt")
		if err != nil {
			return err
		}
		defer file.Close()

		for _, history := range chs {
			_, err := file.WriteString(history.Url + "\n")
			if err != nil {
				return err
			}
		}
		return nil

	case "t":
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

	default:
		print("Invalid option " + fileType)
		return nil
	}
}
