package wallhaven

//WallPaperID is a string representing a wallpaper
type WallpaperID string

type WallpaperInfo struct {
	Data Data `json:"data"`
}
type Avatar struct {
	Two00Px  string `json:"200px"`
	One28Px  string `json:"128px"`
	Three2Px string `json:"32px"`
	Two0Px   string `json:"20px"`
}
type Uploader struct {
	Username string `json:"username"`
	Group    string `json:"group"`
	Avatar   Avatar `json:"avatar"`
}

type Tags struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	CreatedAt  string `json:"created_at"`
}

type Data struct {
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
	Tags       []Tags   `json:"tags"`
}

type SearchResults struct {
	Data []Data `json:"data"`
	Meta Meta   `json:"meta"`
}
type Thumbs struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

type Meta struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
	Query       interface{} `json:"query"`
}
