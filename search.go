package wallhaven

import (
	"image"
	"net/http"
	"strconv"
)

type Category int

const (
	General = 0x100
	Anime   = 0x010
	People  = 0x001
)

func (c Category) String() string {
	return strconv.FormatInt(int64(c), 2)
}

type Purity int

const (
	SFW     = 0x100
	Sketchy = 0x010
	NSFW    = 0x001
)

func (p Purity) String() string {
	return strconv.FormatInt(int64(p), 2)
}

type Q struct {
	Term       string
	Tags       []string
	ExcudeTags []string
	UserName   string
	TagID      int
	Type       string //Type is one of png/jpg
	Like       WallpaperID
}

type Sort string

type Order string

type TopRange string

type Resolution struct {
	Width  int64
	Height int64
}

func (r Resolution) String() string {
	return fmt.Sprintf("%vx%v", r.Vertical, r.Horizontal)
}

type Ratio struct {
	Vertical   int64
	Horizontal int64
}

func (r Ratio) String() string {
	return fmt.Sprintf("%vx%v", r.Vertical, r.Horizontal)
}

type Search struct {
	Query       Q
	Categories  Category
	Purities    Purity
	Sorting     Sort
	Order       Order
	TopRange    TopRange
	AtLeast     Resolution
	Resolutions []Resolution
	Ratios      []Ratio
	Colors      []image.RGBA
	Page        int64
}

func SearchWallpapers(search *Search) (*SearchResults, error) {
	resp, err := http.Get(getWithBase("/search/"))
	if err != nil {
		return nil, err
	}
	out := &SearchResults{}
	err = processResponse(resp, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
