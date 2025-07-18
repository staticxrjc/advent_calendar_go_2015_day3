package main

type Santa struct {
	locations        []*Location
	location         *Location
	uniqueDeliveries int
}

func (santa *Santa) Route(route string) {
	for _, c := range route {
		switch c {
		case '>':
			santa.move(1, 0)
			break
		case 'v':
			santa.move(0, -1)
			break
		case '<':
			santa.move(-1, 0)
			break
		case '^':
			santa.move(0, 1)
		}

		santa.location.House.numPresents += 1
		if santa.location.House.numPresents == 1 {
			santa.uniqueDeliveries += 1
		}
	}
}

func (santa *Santa) move(x int, y int) {
	for _, tmpLocation := range santa.locations {
		if ((santa.location.X + x) == tmpLocation.X) && (santa.location.Y+y) == tmpLocation.Y {
			santa.location = tmpLocation
			return
		}
	}
	santa.location = &Location{
		X:     santa.location.X + x,
		Y:     santa.location.Y + y,
		House: House{0},
	}
	santa.locations = append(santa.locations, santa.location)
}

func NewSanta() *Santa {
	location := Location{X: 0, Y: 0, House: House{1}}
	locations := []*Location{&location}

	return &Santa{
		location:         &location,
		locations:        locations,
		uniqueDeliveries: 1,
	}
}
