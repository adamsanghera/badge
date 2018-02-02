package badge

import (
	"time"

	"github.com/adamsanghera/auth"
	"github.com/adamsanghera/category"
)

// Badge is an interface for services providing identity insurance.
// Badges are a passport, certifying a recent check of citizenship.
type Badge interface {
	// Commands
	Revoke() error

	// Queries
	BelongsTo() category.Category
	Create() (uid interface{}, bid interface{}, badgeScanner auth.Authenticator, timeToLive time.Duration, err error)
	Refresh() (uid interface{}, bid interface{}, badgeScanner auth.Authenticator, timeToLive time.Duration, err error)
	Validate(uid interface{}, bid interface{}, badgeScanner auth.Authenticator) (bool, error)
}
