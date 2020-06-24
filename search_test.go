package wallhaven

import "testing"

// Ensure the aspect ratio respects the format required by the wallhaven API.
// Horizontal comes first, then it's vertical.
func TestQueryAspectRatio(t *testing.T) {
	s := Search{
		Ratios: []Ratio{
			{
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
