package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"net/http"
)

type PublishEventController struct {
	nu *usecases.PublishEventUseCase
}

func NewPublishEventController(nu *usecases.PublishEventUseCase) *PublishEventController {
	return &PublishEventController{nu: nu}
}

func (sn *PublishEventController) SendNotification(c *gin.Context) {
	var userAction *entities.UserAction

	if err := c.ShouldBindJSON(&userAction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := sn.nu.SendEvent(userAction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event sent successfully"})
}
