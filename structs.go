package wallhaven

import "strconv"

//WallpaperID is a string representing a wallpaper
type WallpaperID string

//TagID is an int representing a tag on wh
type TagID int64

func (t TagID) String() string {
	return strconv.Itoa(int(t))
}

//Avatar user's avatar
type Avatar struct {
	Two00Px  string `json:"200px"`
	One28Px  string `json:"128px"`
	Three2Px string `json:"32px"`
	Two0Px   string `json:"20px"`
}

//Uploader information on who uploaded a given wallpaper
type Uploader struct {
	Username string `json:"username"`
	Group    string `json:"group"`
	Avatar   Avatar `json:"avatar"`
}

//Tag full data on a given wallpaper tag
type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	CreatedAt  string `json:"created_at"`
}

//Wallpaper information about a given wallpaper
type Wallpaper struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	ShortURL   string   `json:"short_url"`
	Uploader   Uploader `json:"uploader"`
	Views      int      `json:"views"`
	Favorites  int      `json:"favorites"`
	Source     string   `json:"source"`
	Purity     string   `json:"purity"`
	Category   string   `json:"category"`
	DimensionX int      `json:"dimension_x"`
	DimensionY int      `json:"dimension_y"`
	Resolution string   `json:"resolution"`
	Ratio      string   `json:"ratio"`
	FileSize   int      `json:"file_size"`
	FileType   string   `json:"file_type"`
	CreatedAt  string   `json:"created_at"`
	Colors     []string `json:"colors"`
	Path       string   `json:"path"`
	Thumbs     Thumbs   `json:"thumbs"`
	Tags       []Tag    `json:"tags"`
}

//Thumbs paths for thumbnail images of wallpaper
type Thumbs struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

//Meta paging and stats metadata for search results
type Meta struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
	Query       interface{} `json:"query"`
}

//User User preference Data (set on wallhaven's user GUI)
type User struct {
	ThumbSize     string   `json:"thumb_size"`
	PerPage       string   `json:"per_page"`
	Purity        []string `json:"purity"`
	Categories    []string `json:"categories"`
	Resolutions   []string `json:"resolutions"`
	AspectRatios  []string `json:"aspect_ratios"`
	ToplistRange  string   `json:"toplist_range"`
	TagBlacklist  []string `json:"tag_blacklist"`
	UserBlacklist []string `json:"user_blacklist"`
}

//Result Structs -- server responses

//SearchResults a wrapper containing search results from wh
type SearchResults struct {
	Data []Wallpaper `json:"data"`
	Meta Meta        `json:"meta"`
}

//TagResult a wrapper containing a single tag result when TagInfo is requested
type TagResult struct {
	Data Tag `json:"data"`
}

//UserResult a wrapper containing user preference data
type UserResult struct {
	Data User `json:"data"`
}

//WallpaperResult a wrapper containing a single wallpaper result when WallpaperInfo is requested
type WallpaperResult struct {
	Data Wallpaper `json:"data"`
}
