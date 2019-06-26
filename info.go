package wallhaven

//GetWallpaperInfo returns a single wallpaper's data given its ID
func GetWallpaperInfo(id WallpaperID) (*Wallpaper, error) {
	resp, err := get("/w/" + string(id))
	if err != nil {
		return nil, err
	}
	info := &WallpaperResult{}
	err = processResponse(resp, info)
	if err != nil {
		return nil, err
	}
	return &info.Data, nil
}
