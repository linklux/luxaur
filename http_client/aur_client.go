package http_client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/linklux/luxaur/model"
)

// TODO Store in config
const RPC_API_URL = "https://aur.archlinux.org/rpc/?v=5"
const CGIT_API_URL = "https://aur.archlinux.org"

type aurInfoResponse struct {
	ResultCount int                    `json:"resultcount"`
	Packages    []model.AurPackageInfo `json:"results"`
}

type aurSearchResponse struct {
	ResultCount int                      `json:"resultcount"`
	Packages    []model.AurPackageSearch `json:"results"`
}

type AurClient struct {
	endpoint string
	options  []string
}

func (a AurClient) Search(query string) (int, []model.AurPackageSearch) {
	response := request("&type=search&arg=" + query)

	res := aurSearchResponse{}
	if err := json.Unmarshal(response, &res); err != nil {
		return 0, []model.AurPackageSearch{}
	}

	return res.ResultCount, res.Packages
}

func (a AurClient) Find(query []string) (int, []model.AurPackageInfo) {
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

// TODO Move this somewhere centralized to make it reusable
func Download(pkg *model.AurPackageInfo) (string, error) {
	base, _ := os.UserHomeDir()
	dir := fmt.Sprintf("%s/.luxaur/data/%s-%s", base, pkg.Name, pkg.Version)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Sprintf("Failed while creating path: %s", err.Error())
	}

	res, err := http.Get(CGIT_API_URL + pkg.Url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	dest := dir + "/" + path.Base(res.Request.URL.String())

	out, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return dest, err
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
