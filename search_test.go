package wallhaven

import "testing"

func TestQueryDesignatedAspectRatio(t *testing.T) {
	s := Search{
		Ratios: []Ratio{
			{
				// Using designated struct members, so that their order aren't relevant.
				Horizontal: 16,
				Vertical:   9,
			},
		},
	}

	query := s.toQuery()
	queryEncoded := query.Encode()

	if queryEncoded != "ratios=16x9" {
		t.Fail()
	}
}

func TestQueryAspectRatio(t *testing.T) {
	s := Search{
		Ratios: []Ratio{
			{16, 9}, // No struct designators, we ensure that Horizontal comes first, then it's vertical.
		},
	}

	query := s.toQuery()
	queryEncoded := query.Encode()

	if queryEncoded != "ratios=16x9" {
		t.Fail()
	}
}
