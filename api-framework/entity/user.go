package entity

import (
	entitysubtables "chat-server/api-framework/entity-sub-tables"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type User struct {
	ApiBase
	FirstName         string        `json:"first_name" gorm:"unique;index"`
	LastName          string        `json:"last_name"`
	DateOfBirth       time.Time     `json:"date_of_birth"`
	EmailId           string        `json:"email_id" gorm:"unique;uniqueIndex" validate:"email"`
	Status            CurrentStatus `json:"user_status" gorm:"default:1"`
	UserAuthDetailsID uint
	UserAuthDetails   entitysubtables.UserAuthDetails `gorm:"constraint:OnDelete:SET NULL;"`
}

// api entity methods
func (user *User) Authorize(httpMethod string) (bool, error) {
	switch httpMethod {
	case fiber.MethodPost:
		return true, nil
	case fiber.MethodGet:
		return true, nil
	}

	return false, fiber.NewError(fiber.StatusUnauthorized)
}
func (user *User) ID() uint {
	return user.ApiBase.ID
}

func (user *User) Validate(httpMethod string) error {
	return nil
}

// function option methods
func (user *User) FillDefaults() error {
	log.Debug("reached method")
	// user.LastUpdatedTime = time.Now()
	return nil
}

func (user *User) RemoveInternalFields() error {
	log.Debug("reached method")
	return nil
}
