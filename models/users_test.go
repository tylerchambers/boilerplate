package models

import (
	"testing"
	"time"
)

func TestUser_SetPassword(t *testing.T) {
	type fields struct {
		Username     string
		Email        string
		Password     string
		PasswordSet  time.Time
		Registered   time.Time
		Premium      bool
		PremiumSince time.Time
		PremiumUntil time.Time
	}
	type args struct {
		clearPw string
		now     time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "empty password",
			fields:  fields{},
			args:    args{now: time.Now(), clearPw: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Username:     tt.fields.Username,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
				PasswordSet:  tt.fields.PasswordSet,
				Registered:   tt.fields.Registered,
				Premium:      tt.fields.Premium,
				PremiumSince: tt.fields.PremiumSince,
				PremiumUntil: tt.fields.PremiumUntil,
			}
			if err := u.SetPassword(tt.args.now, tt.args.clearPw); (err != nil) != tt.wantErr {
				t.Errorf("User.SetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	u := User{}
	password := "hunter2"
	err := u.SetPassword(time.Now(), "hunter23456")
	if err != nil {
		t.Errorf("User.SetPassword() gave an unexpected error: %v", err)
	}
	if u.Password == password {
		t.Error("User.SetPassword() did not hash the password!")
	}
}

func TestUser_MakePremium(t *testing.T) {
	now := time.Date(2021, time.April, 20, 1, 1, 20, 0, time.UTC)
	until := now.Add(time.Hour * 24)
	u := User{}
	err := u.MakePremium(now, until)

	if err != nil {
		t.Errorf("User.MakePremium() gave an unexpected error: %v", err)
	}
	if !u.Premium {
		t.Error("User.MakePremium() did not make the user premium, but should have.")
	}
	if !u.PremiumSince.Equal(now) {
		t.Errorf("User.MakePremium() did not properly set u.PremiumSince. Want: %v, Got: %v", now, u.PremiumSince)
	}
	if !u.PremiumUntil.Equal(until) {
		t.Errorf("User.MakePremium() did not properly set u.PremiumUntil Want: %v, Got: %v", now, u.PremiumUntil)
	}

	u = User{Premium: true}
	err = u.MakePremium(now, until)
	if err == nil {
		t.Error("User.MakePremium() did not return an error when making an already premium user premium, but should have.")
	}
	if err.Error() != "already premium" {
		t.Errorf("User.MakePremium() returned an incorrect error for making an already premium user premium. Want: %v, Got: %v", err.Error(), "already premium")
	}
}

func TestUser_MakeNonPremium(t *testing.T) {
	now := time.Date(2021, time.April, 20, 1, 1, 20, 0, time.UTC)
	u := User{Premium: true, PremiumSince: now}
	u.MakeNonPremium()

	if u.Premium {
		t.Error("User.MakeNonPremium() did not demote the user.")
	}
	if !u.PremiumSince.Equal(time.Time{}) {
		t.Errorf("User.MakeNonPremium() did not set PremiumSince back to 0 (time.Time{})")
	}
}

func TestUser_Subscribe(t *testing.T) {
	now := time.Date(2021, time.April, 20, 1, 1, 20, 0, time.UTC)
	u := User{}
	u.Subscribe(now)
	if !u.Subscribed {
		t.Error("User.Subscribe() did not set u.Subscribed to true.")
	}
	if !u.SubscribedOn.Equal(now) {
		t.Errorf("User.Subscribe() did not properly set u.SubscribedOn. Want: %v, Got: %v", now, u.SubscribedOn)
	}
}

func TestUser_Unsubscribe(t *testing.T) {
	u := User{Subscribed: true}
	u.Unsubscribe()
	if u.Subscribed {
		t.Error("User.Unsubscribe did not properly unsubscribe the user.")
	}
}

func TestNewUser(t *testing.T) {
	now := time.Date(2021, time.April, 20, 1, 1, 20, 0, time.UTC)
	username := "test"
	email := "test@example.com"
	password := "hunter23456"
	u, err := NewUser(now, "test", email, password)
	if err != nil {
		t.Errorf("NewUser gave an unexpected error: %v", err)
	}
	if u.Username != "test" {
		t.Errorf("NewUser did not properly set username. Want: %s, Got: %s", username, u.Username)
	}
	if u.Email != email {
		t.Errorf("NewUser did not properly set username. Want: %s, Got: %s", email, u.Email)
	}
}
