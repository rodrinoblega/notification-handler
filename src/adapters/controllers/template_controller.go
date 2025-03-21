package controllers

import (
	"github.com/rodrinoblega/notification_handler/src/adapters/repositories"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	Repo *repositories.PostgresTemplateRepository
}

func NewTemplateController(repo *repositories.PostgresTemplateRepository) *TemplateController {
	return &TemplateController{Repo: repo}
}

func (tc *TemplateController) CreateTemplate(ctx *gin.Context) {
	var notificationTemplate *entities.NotificationTemplate

	if err := ctx.ShouldBindJSON(&notificationTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.Repo.CreateTemplate(notificationTemplate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Template created successfully"})
}

func (tc *TemplateController) GetTemplates(ctx *gin.Context) {
	templates, err := tc.Repo.GetAllTemplates()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, templates)
}

func (tc *TemplateController) UpdateTemplate(ctx *gin.Context) {
	var notificationTemplate *entities.NotificationTemplate

	templateName := ctx.Param("type")

	if err := ctx.ShouldBindJSON(&notificationTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.Repo.UpdateTemplate(templateName, notificationTemplate.Template)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Template updated successfully"})
}

func (tc *TemplateController) DeleteTemplate(ctx *gin.Context) {
	templateName := ctx.Param("type")

	err := tc.Repo.DeleteTemplate(templateName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}
