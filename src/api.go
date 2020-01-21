package main

import (
	"github.com/emdoyle/marvin/src/domain"
	"github.com/emdoyle/marvin/src/models"
	"log"
	"strings"
)

//UserDAO defines operations for persisting UserResources
type UserDAO interface {
	GetUserByID(id uint) domain.UserResource
	GetUserBySlackID(slackID string) domain.UserResource
	CreateUser(user *domain.UserResource) bool
	DeleteUserByID(id uint) bool
	DeleteUserBySlackID(slackID string) bool
	UpdateUser(user *domain.UserResource) bool
}

//GetUserByID gets a UserResource by its primary key ID
func GetUserByID(id uint, userDAO UserDAO) domain.UserResource {
	return userDAO.GetUserByID(id)
}

//GetUserBySlackID gets a UserResource by its Slack ID
func GetUserBySlackID(slackID string, userDAO UserDAO) domain.UserResource {
	return userDAO.GetUserBySlackID(slackID)
}

//CreateUser persists a UserResource
func CreateUser(user *domain.UserResource, userDAO UserDAO) bool {
	return userDAO.CreateUser(user)
}

//DeleteUserByID deletes a UserResource by its primary key ID
func DeleteUserByID(id uint, userDAO UserDAO) bool {
	return userDAO.DeleteUserByID(id)
}

//DeleteUserBySlackID deletes a UserResource by its Slack ID
func DeleteUserBySlackID(slackID string, userDAO UserDAO) bool {
	return userDAO.DeleteUserBySlackID(slackID)
}

//UpdateUser updates a persisted UserResource
func UpdateUser(user *domain.UserResource, userDAO UserDAO) bool {
	return userDAO.UpdateUser(user)
}

//HandleUserAPIRequest handles an incoming API request related to UserResources
func HandleUserAPIRequest(event Event) {
	dao := models.GORMUserDAO{DB: DB}
	if strings.Contains(event.Text, "remember me") {
		log.Printf("Request received to remember user: %s", event.User)
		CreateUser(&domain.UserResource{
			SlackUserID: event.User,
		}, dao)
		log.Printf("Created user with SlackUserID: %s", event.User)
	} else if strings.Contains(event.Text, "forget me") {
		log.Printf("Request received to forget user: %s", event.User)
		DeleteUserBySlackID(event.User, dao)
	} else {
		log.Print("Unknown request received.")
	}
}
