package wallhaven

import "net/http"

func GetWallpaperInfo(id WallpaperID) (*WallpaperInfo, error) {

	resp, err := http.Get(getWithBase("/w/" + string(id)))
	if err != nil {
		return nil, err
	}
	info := &WallpaperInfo{}
	err = processResponse(resp, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
