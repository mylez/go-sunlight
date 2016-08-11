package internal

import (
	"encoding/json"
	"github.com/mylez/go-sunlight/sunlight"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

import "log"

func GenerateURL(root string, resource ...string) string {
	ret := root + "/" + strings.Join(resource, "/")
	return ret
}

func QueryURL(root string, params map[string]string, resource ...string) string {
	apikey := sunlight.GetAPIKey()
	uri := GenerateURL(root, resource...) + "?apikey=" + apikey
	for k, v := range params {
		uri = uri + "&" + url.QueryEscape(k) + "=" + url.QueryEscape(v)
	}
	return uri
}


func GetURL(target interface{}, root string, params map[string]string, resource ...string) (err error) {
	url := QueryURL(root, params, resource...)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, target)
	return
}
