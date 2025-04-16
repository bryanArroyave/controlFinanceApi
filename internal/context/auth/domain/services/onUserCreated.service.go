package services

import (
	"fmt"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
)

func UserCreated(userID int, repo ports.IUserEventsRepository) {
	err := repo.Publish("user.registered", userID)
	if err != nil {
		fmt.Println("Error publishing event:", err)
	}
}
