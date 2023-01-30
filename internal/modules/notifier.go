package modules

import "github.com/russtone/sonar/internal/database/models"

type Notification struct {
	user    *models.User
	payload *models.Payload
	event   *models.Event
}

// Notifier must be implemented by all modules, which are going to notify
// users about payload events.
type Notifier interface {

	// Notify is called every time payload event happens.
	Notify(Notification) error
}
