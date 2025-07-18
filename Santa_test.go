package main

import "testing"

func TestRouteSanta(t *testing.T) {
	testRoute := []struct {
		name       string
		route      string
		deliveries int
	}{
		{
			name:       "> delivers to 2 houses",
			route:      ">",
			deliveries: 2,
		},
		{
			name:       "^>v< delivers presents to 4 houses",
			route:      "^>v<",
			deliveries: 4,
		},
		{
			name:       "^v^v^v^v^v delivers presents to 2 houses",
			route:      "^v^v^v^v^v",
			deliveries: 2,
		},
	}
	for _, test := range testRoute {
		t.Run(test.name, func(t *testing.T) {
			santa := NewSanta()
			santa.Route(test.route)

			got := santa.uniqueDeliveries
			want := test.deliveries

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
	testLocation := []struct {
		name   string
		route  string
		expect Location
	}{
		{
			name:   "Santa starts at 0,0",
			route:  "",
			expect: Location{0, 0, House{1}},
		},
		{
			name:   "> leaves santa at x = 1, y = 0",
			route:  ">",
			expect: Location{1, 0, House{1}},
		},
		{
			name:   "^>v< leaves santa at x = 0, y = 0",
			route:  "^>v<",
			expect: Location{0, 0, House{2}},
		},
		{
			name:   "^v^v^v^v^v leaves santa at x = 0, y = 0",
			route:  "^v^v^v^v^v",
			expect: Location{0, 0, House{6}},
		},
	}
	for _, test := range testLocation {
		t.Run(test.name, func(t *testing.T) {
			santa := NewSanta()
			santa.Route(test.route)

			got := santa.location
			want := test.expect

			if got.X != want.X || got.Y != want.Y {
				t.Errorf("got X:%d,Y:%d want X:%d,Y:%d", got.X, got.Y, want.X, want.Y)
			}

		})
	}
}

func TestRouteSantaRobot(t *testing.T) {
	testRoute := []struct {
		name       string
		route      string
		deliveries int
	}{
		{
			name:       "^v delivers to 3 houses",
			route:      "^v",
			deliveries: 3,
		},
		{
			name:       "^>v< delivers presents to 3 houses",
			route:      "^>v<",
			deliveries: 3,
		},
		{
			name:       "^v^v^v^v^v delivers presents to 11 houses",
			route:      "^v^v^v^v^v",
			deliveries: 11,
		},
	}
	for _, test := range testRoute {
		t.Run(test.name, func(t *testing.T) {
			santa := NewSantaWithRobot()
			santa.Route(test.route)

			got := santa.uniqueDeliveries
			want := test.deliveries

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
	testLocation := []struct {
		name        string
		route       string
		expectSanta Location
		expectRobot Location
	}{
		{
			name:        "Santa and robot starts at 0,0",
			route:       "",
			expectSanta: Location{0, 0, House{1}},
			expectRobot: Location{0, 0, House{1}},
		},
		{
			name:        "> leaves santa at x = 1, y = 0 and robot x = 0, y = 0",
			route:       ">",
			expectSanta: Location{1, 0, House{1}},
			expectRobot: Location{0, 0, House{1}},
		},
		{
			name:        "^>v< leaves santa and robot at x = 0, y = 0",
			route:       "^>v<",
			expectSanta: Location{0, 0, House{1}},
			expectRobot: Location{0, 0, House{1}},
		},
		{
			name:        "^v^v^v^v^v leaves santa at x = 0, y = 5 and robot x = 0, y = -5",
			route:       "^v^v^v^v^v",
			expectSanta: Location{0, 5, House{1}},
			expectRobot: Location{0, -5, House{1}},
		},
	}
	for _, test := range testLocation {
		t.Run(test.name, func(t *testing.T) {
			santa := NewSantaWithRobot()
			santa.Route(test.route)

			got1 := santa.location
			got2 := santa.locationRobot
			want1 := test.expectSanta
			want2 := test.expectRobot

			if got1.X != want1.X || got1.Y != want1.Y {
				t.Errorf("got X:%d,Y:%d want X:%d,Y:%d", got1.X, got1.Y, want1.X, want1.Y)
				t.Errorf("got X:%d,Y:%d want X:%d,Y:%d", got2.X, got2.Y, want2.X, want2.Y)
			}

		})
	}
}
