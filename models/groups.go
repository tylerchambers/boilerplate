package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Group represents a group of users.
type Group struct {
	ID      uuid.UUID
	Name    string
	Members []*User
	Created time.Time
}

func NewGroup(now time.Time, name string, members []*User) *Group {
	g := new(Group)
	g.SetID()
	g.AddMembers()
	g.Created = now
	g.AddMembers(members...)
}

// SetID sets the ID of the group to a new UUID.
func (g *Group) SetID() {
	g.ID = uuid.New()
}

// SetName sets the name of the group.
func (g *Group) SetName(name string) error {
	g.Name = name
	return nil
}

// DeleteMember removes a member from the group.
func (g *Group) DeleteMember(user *User) error {
	for i, v := range g.Members {
		if v.ID == user.ID {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			return nil
		}
	}
	return errors.New("member could not be removed because it was not in the group")
}

// DeleteMemberByID takes a UUID representing a user and deletes them from the group.
// Returns an error if the member is not in the group to begin with.
func (g *Group) DeleteMemberByID(id uuid.UUID) error {
	for i, v := range g.Members {
		if v.ID == id {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			return nil
		}
	}
	return errors.New("member could not be removed because it was not in the group")
}

// AddMember adds a member to the group.
func (g *Group) AddMember(user *User) error {
	g.Members = append(g.Members, user)
	return nil
}

// AddMembers adds multiple members to the group.
func (g *Group) AddMembers(members ...*User) {
	g.Members = append(g.Members, members...)
}

// Contains returns true if a user is in the group.
func (g *Group) Contains(user *User) bool {
	for _, v := range g.Members {
		if v.ID == user.ID {
			return true
		}
	}
	return false
}
