package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrinoblega/notification_handler/config"
	"github.com/rodrinoblega/notification_handler/src/adapters/controllers"
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

	kafkaProducer, err := messaging.NewKafkaProducer(envConf)
	if err != nil {
		log.Fatal(err)
	}

	userActionRepository := repositories.NewUserActionRepository(db)
	publishEventUseCase := usecases.NewPublishEventUseCase(kafkaProducer, userActionRepository)
	publishEventController := controllers.NewPublishEventController(publishEventUseCase)

	showNotificationsUseCase := usecases.NewShowNotificationsUseCase(repo)
	notificationController := controllers.NewNotificationController(showNotificationsUseCase)

	templateRepository := repositories.NewPostgresTemplateRepository(db)
	templateController := controllers.NewTemplateController(templateRepository)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/event/publish", publishEventController.SendNotification)
	api.GET("/notifications", notificationController.GetNotificationsHandler)

	api.POST("/templates", templateController.CreateTemplate)
	api.GET("/templates", templateController.GetTemplates)
	api.PUT("/templates/:type", templateController.UpdateTemplate)
	api.DELETE("/templates/:type", templateController.DeleteTemplate)

	router.Run(":8080")
}
