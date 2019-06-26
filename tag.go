package wallhaven

//GetTagInfo returns metadata regarding a single tag when given its int ID
func GetTagInfo(t TagID) (*Tag, error) {
	resp, err := get("/tag/" + string(t))
	if err != nil {
		return nil, err
	}
	res := &TagResult{}
	processResponse(resp, res)
	return &res.Data, nil
}
