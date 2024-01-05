package tickets

import "fmt"

type Ticket struct {
	Id          string
	Name        string
	Email       string
	Destination string
	FlightTime  string
	Price       float64
}

func NewTicket(id, name, email, destination, flightTime string, price float64) Ticket {
	return Ticket{
		Id:          id,
		Name:        name,
		Email:       email,
		Destination: destination,
		FlightTime:  flightTime,
		Price:       price,
	}
}

func (t Ticket) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Email: %s, Destination: %s, FlightTime: %s, Price: %.2f\n", t.Id, t.Name, t.Email, t.Destination, t.FlightTime, t.Price)
}
