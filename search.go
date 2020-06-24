package wallhaven

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Search Types

//Category is an enum used to represent wallpaper categories
type Category int

//Enum for Category Types
const (
	General Category = 0x100
	Anime   Category = 0x010
	People  Category = 0x001
)

func (c Category) String() string {
	return strconv.FormatInt(int64(c), 2)
}

//Purity is an enum used to represent
type Purity int

//Enum for purity types
const (
	SFW     Purity = 0x100
	Sketchy Purity = 0x010
	NSFW    Purity = 0x001
)

func (p Purity) String() string {
	return strconv.FormatInt(int64(p), 2)
}

//Sort enum specifies the various sort types accepted by WH api
type Sort int

//Sort Enum Values
const (
	DateAdded Sort = iota + 1
	Relevance
	Random
	Views
	Favorites
	Toplist
)

func (s Sort) String() string {
	str := [...]string{"", "date_added", "relevance", "random", "views", "favorites", "topList"}
	return str[s]
}

//Order enum specifies the sort orders accepted by WH api
type Order int

//Sort Enum Values
const (
	Desc Order = iota + 1
	Asc
)

func (o Order) String() string {
	str := [...]string{"", "desc", "asc"}
	return str[o]
}

//TopRange is used to specify the time window for 'top' result when topList is chosen as sort param
type TopRange int

//Enum for TopRange values
const (
	Day TopRange = iota + 1
	ThreeDay
	Week
	Month
	SixMonth
	Year
)

func (t TopRange) String() string {
	str := [...]string{"", "1d", "3d", "1w", "1m", "6m", "1y"}
	return str[t]
}

//Resolution specifies the image resolution to find
type Resolution struct {
	Width  int64
	Height int64
}

func (r Resolution) String() string {
	return fmt.Sprintf("%vx%v", r.Width, r.Height)
}

func (r Resolution) isValid() bool {
	return r.Width > 0 && r.Height > 0
}

//Ratio may be used to specify the aspect ratio of the search
type Ratio struct {
	Vertical   int64
	Horizontal int64
}

func (r Ratio) String() string {
	return fmt.Sprintf("%vx%v", r.Horizontal, r.Vertical)
}

func (r Ratio) isValid() bool {
	return r.Vertical > 0 && r.Horizontal > 0
}

//Q is used to hold the Q params for various fulltext options that the WH Search supports
type Q struct {
	Tags       []string
	ExcudeTags []string
	UserName   string
	TagID      int
	Type       string //Type is one of png/jpg
	Like       WallpaperID
}

func (q Q) toQuery() url.Values {
	var sb strings.Builder
	for _, tag := range q.Tags {
		sb.WriteString("+")
		sb.WriteString(tag)
	}
	for _, etag := range q.ExcudeTags {
		sb.WriteString("-")
		sb.WriteString(etag)
	}
	if len(q.UserName) > 0 {
		sb.WriteString("@")
		sb.WriteString(q.UserName)
	}
	if len(q.Type) > 0 {
		sb.WriteString("type:")
		sb.WriteString(q.Type)
	}
	out := url.Values{}
	val := sb.String()
	if len(val) > 0 {
		out.Set("q", val)
	}
	return out
}

//Search provides various parameters to search for on wallhaven
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
	Colors      []string //Colors is an array of hex colors represented as strings in #RRGGBB format
	Page        int64
}

func (s Search) toQuery() url.Values {
	v := s.Query.toQuery()
	if s.Categories > 0 {
		v.Add("categories", s.Categories.String())
	}
	if s.Purities > 0 {
		v.Add("purity", s.Purities.String())
	}
	if s.Sorting > 0 {
		v.Add("sorting", s.Sorting.String())
	}
	if s.Order > 0 {
		v.Add("order", s.Order.String())
	}
	if s.TopRange > 0 && s.Sorting == Toplist {
		v.Add("topRange", s.TopRange.String())
	}
	if s.AtLeast.isValid() {
		v.Add("atleast", s.AtLeast.String())
	}
	if len(s.Resolutions) > 0 {
		outRes := []string{}
		for _, res := range s.Resolutions {
			if res.isValid() {
				outRes = append(outRes, res.String())
			}
		}
		if len(outRes) > 0 {
			v.Add("resolutions", strings.Join(outRes, ","))
		}
	}
	if len(s.Ratios) > 0 {
		outRat := []string{}
		for _, rat := range s.Ratios {
			if rat.isValid() {
				outRat = append(outRat, rat.String())
			}
		}
		if len(outRat) > 0 {
			v.Add("ratios", strings.Join(outRat, ","))
		}
	}
	if len(s.Colors) > 0 {
		v.Add("colors", strings.Join([]string(s.Colors), ","))
	}
	return v
}

//SearchWallpapers performs a search on WH given a set of criteria.
//Note that this API behaves slightly differently than the various
//single item apis as it also includes the metadata for paging purposes
func SearchWallpapers(search *Search) (*SearchResults, error) {

	resp, err := getWithValues("/search/", search.toQuery())
	out := &SearchResults{}
	err = processResponse(resp, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
