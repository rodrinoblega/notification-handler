package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"net/http"
	"time"
)

type PublishEventController struct {
	nu *usecases.PublishEventUseCase
}

func NewPublishEventController(nu *usecases.PublishEventUseCase) *PublishEventController {
	return &PublishEventController{nu: nu}
}

func (sn *PublishEventController) SendNotification(c *gin.Context) {
	var request struct {
		UserID  string `json:"user_id"`
		Type    string `json:"type"`
		Content string `json:"content"`
		Status  string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	notification := entities.Notification{
		UserID:    request.UserID,
		Content:   request.Content,
		Status:    request.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := sn.nu.Publisher.Publish(notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notification)
}
