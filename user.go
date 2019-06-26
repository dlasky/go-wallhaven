package wallhaven

import (
	"net/http"
	"net/url"
)

//GetUserSettings returns user preferences for the current logged in user. This method requires an API key
func GetUserSettings(apiKey string) (*User, error) {
	u, _ := url.Parse(getWithBase("/settings"))
	q := u.Query()
	q.Add("apiKey", apiKey)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	info := &UserResult{}
	processResponse(resp, info)
	return &info.Data, nil
}
