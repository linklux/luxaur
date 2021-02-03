package aur_util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/linklux/luxaur/model"
)

// TODO Store in config
const RPC_API_URL = "https://aur.archlinux.org/rpc/?v=5"

type aurInfoResponse struct {
	ResultCount int                    `json:"resultcount"`
	Packages    []model.AurPackageInfo `json:"results"`
}

type aurSearchResponse struct {
	ResultCount int                      `json:"resultcount"`
	Packages    []model.AurPackageSearch `json:"results"`
}

func Search(query string) (int, []model.AurPackageSearch) {
	response := request("&type=search&arg=" + query)

	res := aurSearchResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		return 0, []model.AurPackageSearch{}
	}

	return res.ResultCount, res.Packages
}

func Find(query []string) (int, []model.AurPackageInfo) {
	args := ""
	for _, arg := range query {
		args += "&arg[]=" + arg
	}

	response := request("&type=info" + args)

	res := aurInfoResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		return 0, []model.AurPackageInfo{}
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
