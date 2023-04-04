package wallhaven

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetUserSettings returns user preferences for the current logged in user. This method requires an API key
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

// GetUserCollections returns the user's collections. This method requires an API key
// If NO username is provided (i.e: username == "") then this method will returns the authenticated user collections (public and private).
// If one username is provided (i.e: username != "") then this method will only returns user PUBLIC collections
func GetUserCollections(apiKey string, username string) (*[]Collection, error) {
	var u *url.URL
	if username == "" {
		u, _ = url.Parse(getWithBase("/collections"))
	} else {
		u, _ = url.Parse(getWithBase("/collections/" + username))
	}
	q := u.Query()
	q.Add("apiKey", apiKey)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	collections := &CollectionsResult{}
	processResponse(resp, collections)
	return &collections.Data, nil
}

// GetUserCollectionWallpapers returns the Wallpapers as array from the user's collections. This method requires an API key.
// Username is required.
func GetUserCollectionWallpapers(apiKey string, username string, id CollectionID) (*[]Wallpaper, error) {
	u, _ := url.Parse(getWithBase(fmt.Sprintf("%s/%s/%s", "/collections", username, id)))
	q := u.Query()
	q.Add("apiKey", apiKey)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	wallpapers := &CollectionWallpapersResult{}
	processResponse(resp, wallpapers)
	return &wallpapers.Data, nil

}
