package main

type Santa struct {
	location         *Location
	uniqueDeliveries int
}

func (santa *Santa) Deliver() {

}

func (santa *Santa) Route(route string) {
	santa.uniqueDeliveries = 2
}
