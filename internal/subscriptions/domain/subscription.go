package domain

type SubscriptionID string

type Subscription struct {
	ID          SubscriptionID
	Name        string
	Description string
}
