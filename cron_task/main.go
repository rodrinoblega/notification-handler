package main

import (
	"github.com/rodrinoblega/notification_handler/config"
	"github.com/rodrinoblega/notification_handler/src/adapters/notification_providers"
	"github.com/rodrinoblega/notification_handler/src/adapters/repositories"
	database "github.com/rodrinoblega/notification_handler/src/frameworks/databases"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"log"
	"os"
	"time"
)

// Simulación del cron job
func runCronTask(useCase *usecases.ReprocessFailureNotificationUseCase) {
	log.Println("Starting cron task execution...")

	useCase.Reprocess()

	log.Println("Finishing cron task execution...")
}

func main() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	envConf := config.Load(os.Getenv("ENV"))
	db := database.NewPostgres(envConf)
	repo := repositories.NewNotificationRepository(db)
	userActionRepository := repositories.NewUserActionRepository(db)
	notificationProviderA := notification_providers.NewNotificationProviderA()
	notificationProviderB := notification_providers.NewNotificationProviderB()
	reprocessFailureNotificationUseCase := usecases.NewReprocessFailureNotificationUseCase(repo, userActionRepository, notificationProviderA, notificationProviderB)

	for {
		select {
		case <-ticker.C:
			runCronTask(reprocessFailureNotificationUseCase)
		}
	}
}
