package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	tickets "github.com/bootcamp-go/desafio-go-bases/app/internal/tickets/model"
	"github.com/bootcamp-go/desafio-go-bases/app/internal/utils"
)

// implements the TicketRepoInterface
type TicketRepo struct {
	List []tickets.Ticket
}

const ticketsFile = "./app/data/tickets.csv"

// NewTicketRepo creates a new ticket list
func NewTicketRepo() TicketRepo {
	return TicketRepo{
		List: []tickets.Ticket{},
	}
}

// AddTicket adds a ticket to the list
func (t *TicketRepo) AddTicket(ticket tickets.Ticket) {

	// check if ticket already exists
	// for _, t := range t.List {
	// 	if ticket.Id == t.Id {
	// 		return errors.New("ticket already exists")
	// 	}
	// }

	t.List = append(t.List, ticket)
	// return nil
}

// LoadData loads the data from the csv file
func (t *TicketRepo) LoadData() (err error) {

	// open the csv
	file, err := os.Open(ticketsFile)
	if err != nil {
		return ErrLoadData
	}
	defer file.Close()

	// create a reader
	reader := csv.NewReader(file)

	for {
		// get a row
		record, err := reader.Read()
		if err != nil { // EOF
			break
		}

		// create the ticket
		ticket := tickets.NewTicket(record[0], record[1], record[2], record[3], record[4], 0)
		// parse the price
		ticket.Price, err = strconv.ParseFloat(record[5], 64)
		if err != nil {
			return ErrLoadData
		}

		// add the ticket to the list
		t.List = append(t.List, ticket)

	}

	return nil
}

// PrintTickets prints the tickets in the list
func (t *TicketRepo) PrintTickets() {
	for _, ticket := range t.List {
		fmt.Println(ticket)
	}
}

// GetTotalTickets returns the total number of tickets for a destination
func (t *TicketRepo) GetTotalTickets(destination string) (total int, err error) {

	if len(t.List) == 0 {
		return 0, ErrEmptyList
	}

	for _, t := range t.List {
		if t.Destination == destination {
			total++
		}
	}

	return total, nil

}

// GetCountByPeriod returns the total number of tickets for a time
type TicketTime int

const (
	EarlyMorning TicketTime = iota
	Morning
	Afternoon
	Evening
)

func (t *TicketRepo) GetCountByPeriod(time TicketTime) (total int, err error) {

	for _, t := range t.List {

		var isBtw bool

		switch time {
		case EarlyMorning:
			isBtw, err = isBetween(t, 0, 7)
		case Morning:
			isBtw, err = isBetween(t, 7, 13)
		case Afternoon:
			isBtw, err = isBetween(t, 13, 20)
		case Evening:
			isBtw, err = isBetween(t, 20, 24)
		}

		if err != nil {
			return 0, err
		}
		if isBtw {
			total++
		}
	}

	return total, nil

}

// function to calculate if a ticket is between two hours
// floor is not included, ceil is included
func isBetween(ticket tickets.Ticket, floor, ceil int) (bool, error) {

	// convert time to hour and min
	hour, min, err := utils.ConvertTime(ticket.FlightTime)
	if err != nil {
		return false, err
	}

	if min == 0 && hour >= floor && hour < ceil {
		return true, nil
	}

	if min != 0 && hour >= floor && hour < ceil {
		return true, nil
	}

	return false, nil
}

// AverageDestination returns the average tickets for a destination
func (t *TicketRepo) AverageDestination(destination string) (avg float64, err error) {

	if len(t.List) == 0 {
		return 0, ErrEmptyList
	}

	totalTickets := len(t.List)
	totalDestino, err := t.GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}

	avg = 100 * float64(totalDestino) / float64(totalTickets)

	return avg, nil

}
