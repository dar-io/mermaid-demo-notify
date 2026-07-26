// Package db holds the persisted models.
package db

import "time"

type Recipient struct {
	ID        string
	Email     string
	PushToken string
	CreatedAt time.Time
}

type Notification struct {
	ID          int64
	RecipientID string
	Channel     string
	Body        string
	Delivered   bool
	Attempts    int
	CreatedAt   time.Time
}

type Preference struct {
	RecipientID string
	Channel     string
	Enabled     bool
}
