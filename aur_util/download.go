package aur_util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const CGIT_API_URL = "https://aur.archlinux.org"

func Download(url string, name string, version string) (string, error) {
	base, _ := os.UserHomeDir()
	dir := fmt.Sprintf("%s/.luxaur/data/%s-%s", base, name, version)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Sprintf("Failed while creating path: %s", err.Error())
	}

	res, err := http.Get(CGIT_API_URL + url)
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
