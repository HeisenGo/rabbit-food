package commands

import (
	"client/models"
	"client/services"
	"errors"
	"fmt"
)

type DisplayProfileCommand struct {
	service services.Service
}

func (c *DisplayProfileCommand) Execute(data any) error {
	userID, ok := data.(uint)
	if !ok {
		return errors.New("data type isn't uint")
	}
	profile, err := c.service.DisplayProfile(userID)
	if err != nil {
		return err
	}
	fmt.Println("Profile Information:")
	fmt.Printf("ID: %d\nName: %s\nEmail: %s\nPhone: %s\n", profile.ID, profile.FirstName+" "+profile.LastName, profile.Email, profile.Phone)
	return nil
}

type EditProfileCommand struct {
	service services.Service
}

func (c *EditProfileCommand) Execute(data any) error {
	user, ok := data.(*models.User)
	if !ok {
		return errors.New("data type isn't user")
	}
	updatedProfile, err := c.service.EditProfile(user)
	if err != nil {
		return err
	}
	fmt.Println("Profile updated successfully.")
	fmt.Printf("Updated Profile Information: %+v\n", updatedProfile)
	return nil
}

func NewDisplayProfileCommand(service services.Service) *DisplayProfileCommand {
	return &DisplayProfileCommand{service: service}
}

func NewEditProfileCommand(service services.Service) *EditProfileCommand {
	return &EditProfileCommand{service: service}
}
