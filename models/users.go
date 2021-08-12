package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user of our service.
type User struct {
	gorm.Model
	ID           uuid.UUID
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	PasswordSet  time.Time
	Registered   time.Time
	Premium      bool
	PremiumSince time.Time
	PremiumUntil time.Time
	Subscribed   bool
	SubscribedOn time.Time
}

func (u *User) SetID() {
	u.ID = uuid.New()
}

func (u *User) SetUsername(username string) {
	// TODO: check if username is already taken, etc.
	u.Username = username
}

func (u *User) SetFirstName(first string) {
	u.FirstName = first
}

func (u *User) SetLastName(last string) {
	u.LastName = last
}

func (u *User) SetEmail(email string) {
	// TODO: check if email is already taken, do validation.
	u.Email = email
}

// SetPassword hashes, then set's a user's password.
func (u *User) SetPassword(now time.Time, clearPw string) error {
	if len(clearPw) <= 8 {
		return errors.New("password must be 8 characters or longer")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(clearPw), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	u.PasswordSet = now
	return nil
}

// CheckPassword verifies a cleartext password against a user's password.
func (u *User) CheckPassword(clearPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(clearPw))
	return err == nil
}

// MakePremium marks a user's account as premium.
func (u *User) MakePremium(now time.Time, until time.Time) error {
	if u.Premium {
		return errors.New("already premium")
	}
	u.Premium = true
	if u.PremiumSince.IsZero() {
		u.PremiumSince = now
	}
	u.PremiumUntil = until
	return nil
}

// MakeNonPremium demotes a user from premium / paid down to a regular user.
func (u *User) MakeNonPremium() {
	u.Premium = false
	// Set PremiumSince back to 0
	u.PremiumSince = time.Time{}
}

// Subscribe marks a user as "subscribed", used for recurring billing.
func (u *User) Subscribe(now time.Time) {
	u.Subscribed = true
	u.SubscribedOn = now
}

// Unsubscribe unsusbscribes a user from our service.
func (u *User) Unsubscribe() {
	u.Subscribed = false
}

// NewUser sets up a new user account.
func NewUser(now time.Time, username, email, password string) (*User, error) {
	u := new(User)
	u.SetUsername(username)
	u.SetEmail(email)
	u.SetID()
	err := u.SetPassword(now, password)
	if err != nil {
		return nil, errors.New("could not register new user")
	}
	u.Registered = now
	return u, nil
}
