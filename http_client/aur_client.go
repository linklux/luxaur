package http_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/linklux/luxaur/data"
)

const RPC_API_URL = "https://aur.archlinux.org/rpc/?v=5"
const CGIT_API_URL = "https://aur.archlinux.org"

type aurInfoResponse struct {
	ResultCount int                   `json:"resultcount"`
	Packages    []data.AurPackageInfo `json:"results"`
}

type aurSearchResponse struct {
	ResultCount int                     `json:"resultcount"`
	Packages    []data.AurPackageSearch `json:"results"`
}

type AurClient struct {
	endpoint string
	options  []string
}

func (a AurClient) Search(query string) (int, []data.AurPackageSearch) {
	response := request("&type=search&arg=" + query)

	res := aurSearchResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		return 0, []data.AurPackageSearch{}
	}

	return res.ResultCount, res.Packages
}

func (a AurClient) Find(query []string) (int, []data.AurPackageInfo) {
	args := ""
	for _, arg := range query {
		args += "&arg[]=" + arg
	}

	response := request("&type=info" + args)

	res := aurInfoResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		return 0, []data.AurPackageInfo{}
	}

	return res.ResultCount, res.Packages
}

func request(endpoint string) []byte {
	client := http.Client{Timeout: time.Second * 5}
	url := RPC_API_URL + endpoint

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
