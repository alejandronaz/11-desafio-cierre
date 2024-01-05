package repository_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/app/internal/tickets/model"
	"github.com/bootcamp-go/desafio-go-bases/app/internal/tickets/repository"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {

	t.Run("should return the total of tickets for a given destination", func(t *testing.T) {
		// Arrange
		repo := repository.NewTicketRepo()
		t1 := model.NewTicket("1", "Juan Doe", "juan@example.com", "Argentina", "15:34", 456.0)
		t2 := model.NewTicket("2", "Lionel Messi", "liomessi@example.com", "Argentina", "18:00", 200.0)
		t3 := model.NewTicket("3", "Facundo Campazzo", "facu@example.com", "Spain", "20:00", 100.0)

		repo.AddTicket(t1)
		repo.AddTicket(t2)
		repo.AddTicket(t3)

		// Act
		expected := 2
		total, err := repo.GetTotalTickets("Argentina")

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, total)
	})

	t.Run("should return an error for empty repository", func(t *testing.T) {
		// Arrange
		repo := repository.NewTicketRepo()

		// Act
		_, err := repo.GetTotalTickets("Argentina")

		// Assert
		require.Error(t, err)
	})

}

func TestGetCountByPeriod(t *testing.T) {

	t.Run("should return the total of tickets for a given period", func(t *testing.T) {
		// Arrange
		repo := repository.NewTicketRepo()
		t1 := model.NewTicket("1", "Juan Doe", "juan@example.com", "Argentina", "13:00", 456.0)
		t2 := model.NewTicket("2", "Lionel Messi", "liomessi@example.com", "Argentina", "19:59", 200.0)
		t3 := model.NewTicket("3", "Facundo Campazzo", "facu@example.com", "Spain", "20:00", 100.0)

		repo.AddTicket(t1)
		repo.AddTicket(t2)
		repo.AddTicket(t3)

		// Act
		expected := 2
		total, err := repo.GetCountByPeriod(repository.Afternoon)

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, total)
	})

	t.Run("should return an error due to incorrect time", func(t *testing.T) {
		// Arrange
		repo := repository.NewTicketRepo()
		t1 := model.NewTicket("1", "Juan Doe", "juan@example.com", "Argentina", "13:00", 456.0)
		t2 := model.NewTicket("2", "Lionel Messi", "liomessi@example.com", "Argentina", "19:59", 200.0)
		t3 := model.NewTicket("3", "Facundo Campazzo", "facu@example.com", "Spain", "27:00", 100.0)

		repo.AddTicket(t1)
		repo.AddTicket(t2)
		repo.AddTicket(t3)

		// Act
		_, err := repo.GetCountByPeriod(repository.Afternoon)

		// Assert
		require.Error(t, err)
	})

}

func TestAverageDestination(t *testing.T) {

	t.Run("should return the average tickets for a given destination", func(t *testing.T) {
		// Arrange
		repo := repository.NewTicketRepo()
		t1 := model.NewTicket("1", "Juan Doe", "juan@example.com", "Argentina", "13:00", 456.0)
		t2 := model.NewTicket("2", "Lionel Messi", "liomessi@example.com", "Argentina", "19:59", 200.0)
		t3 := model.NewTicket("3", "Facundo Campazzo", "facu@example.com", "Spain", "20:00", 100.0)

		repo.AddTicket(t1)
		repo.AddTicket(t2)
		repo.AddTicket(t3)

		// Act
		expected := 100 * 2.0 / 3.0
		total, err := repo.AverageDestination("Argentina")

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, total)
	})

	t.Run("should return an error for empty repository", func(t *testing.T) {
		// Arrange
		repo := repository.NewTicketRepo()

		// Act
		_, err := repo.AverageDestination("Argentina")

		// Assert
		require.Error(t, err)
	})

}
