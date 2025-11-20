package application

import "notifier/internal/core/application"

type CreateSubscriptionCommand struct {
	Name        string
	Description string
}

func (c CreateSubscriptionCommand) CommandName() string {
	return "CreateSubscriptionCommand"
}

type CreateSubscriptionCommandHandler struct {
}

func (h *CreateSubscriptionCommandHandler) Handle(ctx application.CommandContext, cmd CreateSubscriptionCommand) error {
	return nil
}
