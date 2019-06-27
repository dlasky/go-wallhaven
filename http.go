package wallhaven

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func processResponse(resp *http.Response, out interface{}) error {
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", byt)
	resp.Body.Close()
	return json.Unmarshal(byt, out)

}

const baseURL = "https://wallhaven.cc/api/v1"

func getWithBase(p string) string {
	return baseURL + p
}

var client = &http.Client{}

func get(p string) (*http.Response, error) {
	return getWithValues(p, url.Values{})
}

func getWithValues(p string, v url.Values) (*http.Response, error) {
	u, err := url.Parse(getWithBase(p))
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	return getAuthedResponse(u.String())
}

func getAuthedResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Key", os.Getenv("WH_API_KEY"))
	return client.Do(req)
}

func download(filepath string, resp *http.Response) error {

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
