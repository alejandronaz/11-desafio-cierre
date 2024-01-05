package main

import (
	"fmt"

	tickets "github.com/bootcamp-go/desafio-go-bases/app/internal/tickets/repostitory"
)

func main() {

	// Create the list of tickets
	ticketList := tickets.NewTicketRepo()

	// Load the data from the csv file
	err := ticketList.LoadData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("----------------------------------------------------")

	// Print the list of tickets
	ticketList.PrintTickets()

	fmt.Println("----------------------------------------------------")

	// Get the total tickets to Argentina
	argentinaTotal, err := ticketList.GetTotalTickets("Argentina")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total tickets to Argentina: ", argentinaTotal)

	fmt.Println("----------------------------------------------------")

	// Get the total tickets to Argentina by period
	earlyMornings, err := ticketList.GetCountByPeriod(tickets.EarlyMorning)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total early morning tickets to Argentina: ", earlyMornings)

	mornings, err := ticketList.GetCountByPeriod(tickets.Morning)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total morning tickets to Argentina: ", mornings)

	afternoon, err := ticketList.GetCountByPeriod(tickets.Afternoon)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total afternoon tickets to Argentina: ", afternoon)

	evening, err := ticketList.GetCountByPeriod(tickets.Evening)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total evening tickets to Argentina: ", evening)

	fmt.Println("----------------------------------------------------")

	// Get the average tickets to Argentina
	avgArg, err := ticketList.AverageDestination("Argentina")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Average tickets to Argentina: %.2f %%\n", avgArg)

	fmt.Println("----------------------------------------------------")

}
