package models

import (
	"github.com/emdoyle/marvin/src/domain"
	"github.com/jinzhu/gorm"
)

//User is a DB model holding Slack user data
type User struct {
	gorm.Model
	SlackUserID string `gorm:"type:varchar(100)"`
}

//GORMUserDAO holds a GORM DB handle and does CRUD on Users in the DB
type GORMUserDAO struct {
	DB *gorm.DB
}

//GetUserByID gets a User from the DAO by ID
func (dao GORMUserDAO) GetUserByID(id uint) domain.UserResource {
	user := User{}
	dao.DB.First(&user, id)
	return domain.UserResource{
		ID:          user.ID,
		SlackUserID: user.SlackUserID,
	}
}

//GetUserBySlackID gets the first User with a matching SlackUserID
func (dao GORMUserDAO) GetUserBySlackID(slackID string) domain.UserResource {
	user := domain.UserResource{}
	dao.DB.Where(&User{SlackUserID: slackID}).First(&user)
	return domain.UserResource{
		ID:          user.ID,
		SlackUserID: user.SlackUserID,
	}
}

//CreateUser uses the DAO to persist a User
func (dao GORMUserDAO) CreateUser(userResource *domain.UserResource) bool {
	user := User{
		Model:       gorm.Model{ID: userResource.ID},
		SlackUserID: userResource.SlackUserID,
	}
	dao.DB.Create(&user)
	return true
}

//DeleteUserByID uses the DAO to delete a User by ID
func (dao GORMUserDAO) DeleteUserByID(id uint) bool {
	user := User{
		Model: gorm.Model{ID: id},
	}
	dao.DB.Delete(&user)
	return true
}

//DeleteUserBySlackID uses the DAO to delete a User by SlackID
func (dao GORMUserDAO) DeleteUserBySlackID(slackID string) bool {
	user := User{
		SlackUserID: slackID,
	}
	dao.DB.Where(&user).Delete(User{})
	return true
}

//UpdateUser uses the DAO to update a User
func (dao GORMUserDAO) UpdateUser(userResource *domain.UserResource) bool {
	user := User{
		Model:       gorm.Model{ID: userResource.ID},
		SlackUserID: userResource.SlackUserID,
	}
	dao.DB.Save(&user)
	return true
}
