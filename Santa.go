package main

type Santa struct {
	locations        []*Location
	location         *Location
	locationRobot    *Location
	uniqueDeliveries int
	hasRobot         bool
}

func (santa *Santa) Route(route string) {
	for i, c := range route {
		switch c {
		case '>':
			if santa.hasRobot && (i%2 != 0) {
				santa.moveRobot(1, 0)
			} else {
				santa.move(1, 0)
			}
			break
		case 'v':
			if santa.hasRobot && (i%2 != 0) {
				santa.moveRobot(0, -1)
			} else {
				santa.move(0, -1)
			}
			break
		case '<':
			if santa.hasRobot && (i%2 != 0) {
				santa.moveRobot(-1, 0)
			} else {
				santa.move(-1, 0)
			}
			break
		case '^':
			if santa.hasRobot && (i%2 != 0) {
				santa.moveRobot(0, 1)
			} else {
				santa.move(0, 1)
			}
		}

		santa.location.House.numPresents += 1
		if santa.location.House.numPresents == 1 {
			santa.uniqueDeliveries += 1
		}
		if santa.hasRobot {
			santa.locationRobot.House.numPresents += 1
			if santa.locationRobot.House.numPresents == 1 {
				santa.uniqueDeliveries += 1
			}
		}
	}
}

func (santa *Santa) moveRobot(x int, y int) {
	for _, tmpLocation := range santa.locations {
		if ((santa.locationRobot.X + x) == tmpLocation.X) && (santa.locationRobot.Y+y) == tmpLocation.Y {
			santa.locationRobot = tmpLocation
			return
		}
	}
	santa.locationRobot = &Location{
		X:     santa.locationRobot.X + x,
		Y:     santa.locationRobot.Y + y,
		House: House{0},
	}
	santa.locations = append(santa.locations, santa.locationRobot)
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
		hasRobot:         false,
	}
}

func NewSantaWithRobot() *Santa {
	santa := NewSanta()
	santa.hasRobot = true
	santa.locationRobot = santa.location

	return santa
}
