package http_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/linklux/luxaur/data"
)

const BASE_API_URL = "https://aur.archlinux.org/rpc/"

type aurFindResponse struct {
	ResultCount int          `json:"resultcount"`
	Package     data.Package `json:"results"`
}

type aurSearchResponse struct {
	ResultCount int            `json:"resultcount"`
	Packages    []data.Package `json:"results"`
}

type AurClient struct {
	endpoint string
	options  []string
}

func (a AurClient) Search(query string) []data.Package {
	response := request("?type=search&arg=" + query)

	res := aurSearchResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		panic(err)
	}

	return res.Packages
}

func (a AurClient) Find(query string) data.Package {
	response := request("?type=info&arg=" + query)

	res := aurFindResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		panic(err)
	}

	return res.Package
}

func request(endpoint string) []byte {
	client := http.Client{Timeout: time.Second * 5}
	url := BASE_API_URL + endpoint

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		panic(reqErr)
	}

	aurResponse, getErr := client.Do(req)
	if getErr != nil {
		panic(getErr)
	}

	if aurResponse.Body != nil {
		defer aurResponse.Body.Close()
	}

	body, readErr := ioutil.ReadAll(aurResponse.Body)
	if readErr != nil {
		panic(readErr)
	}

	return body
}
