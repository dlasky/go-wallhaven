package wallhaven

import "testing"

func TestQueryPagination(t *testing.T) {
	s := Search{
		Page: 2,
	}

	query := s.toQuery()
	queryEncoded := query.Encode()

	if queryEncoded != "page=2" {
		t.Fail()
	}
}