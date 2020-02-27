package main

import (
	"github.com/nik/image-fetcher-service/internal/downloader"
	"github.com/nik/image-fetcher-service/internal/utility"
	"os"
	"path/filepath"
)

type AutoGenerated struct {
	Query struct {
		Apikey string `json:"apikey"`
		Q      string `json:"q"`
		Tbm    string `json:"tbm"`
		Device string `json:"device"`
		URL    string `json:"url"`
	} `json:"query"`
	RelatedSearches []interface{} `json:"related_searches"`
	ImageResults    []struct {
		Position  int    `json:"position"`
		Thumbnail string `json:"thumbnail"`
		SourceURL string `json:"sourceUrl"`
		Title     string `json:"title"`
		Link      string `json:"link"`
		Source    string `json:"source"`
	} `json:"image_results"`
}

var app AutoGenerated

func main() {

	mydir, _ := os.Getwd()
	println(mydir)

	configPath := filepath.FromSlash(mydir + "/config/config.json")

	config, error := utility.LoadConfiguration(configPath)
	if error != nil {
		panic(error)
	}

	downloader := downloader.NewDownloader(config.Url, config.ApiKey, config.SearchImageQuery)

	links, _ := downloader.GetLinks()
	println(links)

	/*queryResponse,_:= downloader.GetSearchResponse()
	for _, element:= range queryResponse.ImageResults {
		println(element.Link)
	}*/

	images, err := downloader.GetImages(links)
	if err != nil {
		println("Images are downloaded and total size is %d", len(images))
	}
}
