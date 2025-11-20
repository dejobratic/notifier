package handlers

import (
	coreapp "notifier/internal/core/application"
	"notifier/internal/infrastructure/api/response"
	subapp "notifier/internal/subscriptions/application"
	"notifier/internal/subscriptions/domain"

	"github.com/gin-gonic/gin"
)

// @Summary Get a subscription by ID
// @Description Get a subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} domain.Subscription
// @Failure 500 {object} response.ErrorResponse
// @Router /subscriptions/{id} [get]
func GetSubscription(c *gin.Context) {
	queryHandler := subapp.GetSubscriptionQueryHandler{}
	query := subapp.GetSubscriptionQuery{
		ID: domain.SubscriptionID(c.Param("id")),
	}
	res, err := queryHandler.Handle(coreapp.BaseQueryContext{}, query)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.Ok(c, res)
}

// @Summary Create a subscription
// @Description Create a subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body subapp.CreateSubscriptionCommand true "Subscription"
// @Success 200 {object} response.MessageResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /subscriptions [post]
func CreateSubscription(c *gin.Context) {
	command := subapp.CreateSubscriptionCommand{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
	}
	commandHandler := subapp.CreateSubscriptionCommandHandler{}
	err := commandHandler.Handle(coreapp.BaseCommandContext{}, command)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.Ok(c, response.MessageResponse{Message: "Subscription created successfully"})
}
