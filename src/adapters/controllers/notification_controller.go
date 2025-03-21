package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"net/http"
	"strconv"
	"time"
)

type NotificationController struct {
	sn *usecases.ShowNotificationsUseCase
}

func NewNotificationController(sn *usecases.ShowNotificationsUseCase) *NotificationController {
	return &NotificationController{sn: sn}
}

func (nc *NotificationController) GetNotificationsHandler(c *gin.Context) {
	userID := c.Query("user_id")
	notificationType := c.Query("type")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	// Parsear fechas
	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	// Parsear limit y offset
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	notifications, err := nc.sn.Search(userID, notificationType, startDate, endDate, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
