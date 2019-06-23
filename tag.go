package wallhaven

import "strconv"

type TagID int64

func (t TagID) String() string {
	return strconv.Itoa(int(t))
}

func GetTagInfo(t TagID) {
	getWithBase("/tag/"+string(t))
}