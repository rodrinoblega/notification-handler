package main

import (
	"github.com/rodrinoblega/notification_handler/config"
	"github.com/rodrinoblega/notification_handler/src/adapters/notification_providers"
	"github.com/rodrinoblega/notification_handler/src/adapters/repositories"
	database "github.com/rodrinoblega/notification_handler/src/frameworks/databases"
	"github.com/rodrinoblega/notification_handler/src/frameworks/messaging"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"log"
	"os"
)

func main() {
	envConf := config.Load(os.Getenv("ENV"))
	db := database.NewPostgres(envConf)
	repo := repositories.NewNotificationRepository(db)
	notificationProviderA := notification_providers.NewNotificationProviderA()
	notificationProviderB := notification_providers.NewNotificationProviderB()

	saveNotificationUseCase := usecases.NewSaveNotificationUseCase(repo, notificationProviderA, notificationProviderB)

	kafkaConsumer, err := messaging.NewKafkaConsumer(envConf, saveNotificationUseCase)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Subscribe()
}
