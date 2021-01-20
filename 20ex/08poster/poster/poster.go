package poster

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	URL    = "http://www.omdbapi.com/"
	Apikey = "8918582d"
)

func SearchMoive(title string) (*MovieInfo, error) {
	formatUrl := URL + "?apikey=" + Apikey + "&t=" + url.QueryEscape(title)
	resp, err := http.Get(formatUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search movie failed: %d", resp.StatusCode)
	}

	var result MovieInfo
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func DownloadImage(u, dir, name string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(dir + "/" + name)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}
