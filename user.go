package wallhaven

import (
	"net/http"
	"net/url"
)

type UserInfo struct {
	Data struct {
		ThumbSize     string   `json:"thumb_size"`
		PerPage       string   `json:"per_page"`
		Purity        []string `json:"purity"`
		Categories    []string `json:"categories"`
		Resolutions   []string `json:"resolutions"`
		AspectRatios  []string `json:"aspect_ratios"`
		ToplistRange  string   `json:"toplist_range"`
		TagBlacklist  []string `json:"tag_blacklist"`
		UserBlacklist []string `json:"user_blacklist"`
	} `json:"data"`
}

func GetUserSettings(apiKey string) (*UserInfo, error) {
	u, _ := url.Parse(getWithBase("/settings"))
	q := u.Query()
	q.Add("apiKey", apiKey)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	info := &UserInfo{}
	processResponse(resp, info)
	return info, nil
}
