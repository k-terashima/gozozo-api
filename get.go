package gozozo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const domain = "http://zozo.jp"

type SnapRanking struct {
	Snap struct {
		Data []struct {
			Ranking  string `json:"ranking"`
			Count    string `json:"count"`
			Shop     string `json:"shop"`
			Shopurl  string `json:"shopurl"`
			Name     string `json:"name"`
			Usertype string `json:"usertype"`
			Main     struct {
				Img string `json:"img"`
				URL string `json:"url"`
				Sex string `json:"sex"`
				Alt string `json:"alt"`
			} `json:"main"`
			Sub []struct {
				Img          string `json:"img"`
				URL          string `json:"url"`
				Title        string `json:"title"`
				Price        string `json:"price"`
				Pricesale    string `json:"pricesale"`
				Discountrate string `json:"discountrate"`
				Pricetype    string `json:"pricetype"`
				Weboff       string `json:"weboff"`
			} `json:"sub"`
		} `json:"data"`
	} `json:"snap"`
}

func (p *SnapRanking) GetRanking() error {
	url := filepath.Join(domain, "/data/snapranking.txt")
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// delete BOM
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	var snap SnapRanking
	if err := json.Unmarshal(body, &snap); err != nil {
		return err
	}

	for i, v := range snap.Snap.Data {
		fmt.Printf("%d-%s: %s(%s)\n", i, v.Ranking, v.Name, v.Shop)
	}

	return nil
}
