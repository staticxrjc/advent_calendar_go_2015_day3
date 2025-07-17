package main

import "testing"

func Test(t *testing.T) {
	testRoute := []struct {
		name     string
		route    string
		expected int
	}{
		{
			name:     "> delvers to 2 houses",
			route:    ">",
			expected: 2,
		},
	}
	for _, test := range testRoute {
		t.Run(test.name, func(t *testing.T) {
			location := Location{
				X: 0,
				Y: 0,
			}
			santa := Santa{
				location: &location,
			}
			santa.Route(test.route)

			got := santa.uniqueDeliveries
			want := test.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
