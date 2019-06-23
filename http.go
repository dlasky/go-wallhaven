package wallhaven

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
