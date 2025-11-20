package handlers

import (
	coreapp "notifier/internal/core/application"
	subapp "notifier/internal/subscriptions/application"
	"notifier/internal/subscriptions/domain"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

type MessageResponse struct {
	Message string `json:"message" example:"success message"`
}

// @Summary Get a subscription by ID
// @Description Get a subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} domain.Subscription
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [get]
func GetSubscription(c *gin.Context) {
	queryHandler := subapp.GetSubscriptionQueryHandler{}
	query := subapp.GetSubscriptionQuery{
		ID: domain.SubscriptionID(c.Param("id")),
	}
	response, err := queryHandler.Handle(coreapp.BaseQueryContext{}, query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}

// @Summary Create a subscription
// @Description Create a subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body subapp.CreateSubscriptionCommand true "Subscription"
// @Success 200 {object} MessageResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
func CreateSubscription(c *gin.Context) {
	command := subapp.CreateSubscriptionCommand{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
	}
	commandHandler := subapp.CreateSubscriptionCommandHandler{}
	err := commandHandler.Handle(coreapp.BaseCommandContext{}, command)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Subscription created successfully"})
}
