package learngo

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	gacil := Passenger{"Angel", 1, false}
	fmt.Println(gacil)

	var (
		ricad = Passenger{Name: "Ricad", TicketNumber: 2}
		eric  = Passenger{Name: "Eric", TicketNumber: 2}
	)
	fmt.Println(ricad, eric)

	var joni Passenger
	joni.Name = "Rezon"
	joni.TicketNumber = 3
	fmt.Println(joni)

	gacil.Boarded = true
	ricad.Boarded = true

	if ricad.Boarded {
		fmt.Println("Ricad has boarded the bus")
	}

	if gacil.Boarded {
		fmt.Println(gacil.Name, " has boarded the bus")
	}

	joni.Boarded = true
	bus := Bus{joni}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, " is in the front seat!")
}
