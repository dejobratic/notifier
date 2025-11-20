package application

import (
	"notifier/internal/core/application"
	"notifier/internal/subscriptions/domain"
)

type GetSubscriptionQuery struct {
	ID domain.SubscriptionID
}

func (q GetSubscriptionQuery) QueryName() string {
	return "GetSubscriptionQuery"
}

type GetSubscriptionQueryHandler struct {
}

func (h *GetSubscriptionQueryHandler) Handle(ctx application.QueryContext, query GetSubscriptionQuery) (domain.Subscription, error) {
	return domain.Subscription{
		ID:          query.ID,
		Name:        "Test Subscription",
		Description: "This is a test subscription",
	}, nil
}
