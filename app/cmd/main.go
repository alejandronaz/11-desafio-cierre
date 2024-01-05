package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/app/internal/tickets/repository"
)

func main() {

	// Create the list of tickets
	ticketRepo := repository.NewTicketRepo()

	// Load the data from the csv file
	err := ticketRepo.LoadData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("----------------------------------------------------")

	// Print the list of tickets
	ticketRepo.PrintTickets()

	fmt.Println("----------------------------------------------------")

	// Get the total tickets to Argentina
	argentinaTotal, err := ticketRepo.GetTotalTickets("Argentina")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total tickets to Argentina: ", argentinaTotal)

	fmt.Println("----------------------------------------------------")

	// Get the total tickets to Argentina by period
	earlyMornings, err := ticketRepo.GetCountByPeriod(repository.EarlyMorning)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total early morning tickets to Argentina: ", earlyMornings)

	mornings, err := ticketRepo.GetCountByPeriod(repository.Morning)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total morning tickets to Argentina: ", mornings)

	afternoon, err := ticketRepo.GetCountByPeriod(repository.Afternoon)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total afternoon tickets to Argentina: ", afternoon)

	evening, err := ticketRepo.GetCountByPeriod(repository.Evening)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total evening tickets to Argentina: ", evening)

	fmt.Println("----------------------------------------------------")

	// Get the average tickets to Argentina
	avgArg, err := ticketRepo.AverageDestination("Argentina")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Average tickets to Argentina: %.2f %%\n", avgArg)

	fmt.Println("----------------------------------------------------")

}
