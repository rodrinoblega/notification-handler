package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrinoblega/notification_handler/config"
	"github.com/rodrinoblega/notification_handler/src/adapters/controllers"
	"github.com/rodrinoblega/notification_handler/src/frameworks/messaging"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"log"
	"os"
)

func main() {
	envConf := config.Load(os.Getenv("ENV"))

	kafkaProducer, err := messaging.NewKafkaProducer(envConf)
	if err != nil {
		log.Fatal(err)
	}

	useCase := usecases.NewPublishEventUseCase(kafkaProducer)
	controller := controllers.NewPublishEventController(useCase)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/event/publish", controller.SendNotification)

	router.Run(":8080")
}
