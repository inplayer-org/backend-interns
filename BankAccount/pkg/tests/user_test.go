package tests

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"bankacc/pkg/entities"
)


func MySQLInit() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Welcome1!"
	dbName := "BankAccount?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println(err)
	}
	return db
}
type TearDownTestSuite interface {
	TearDownTest()
}
type UserTestSuite struct {
	suite.Suite
	User      *entities.User
	UserS     []entities.User
	UserPs    []*entities.User
	Userstore store.UserModel
	Db        *sql.DB
}
func (suite *UserTestSuite) SetupTest() {
	var err error
	suite.Db = MySQLInit()
	user := store.NewUserStoreModel(suite.Db)
	suite.UserS = []entities.User{
		{
			FullName:    "Filip Krs",
			Email:       "krs@",
			PhoneNumber: "1",
		},
		{
			FullName:    "Viktor Pat",
			Email:       "pat@",
			PhoneNumber: "2",
		},
		{
			FullName:    "Darko Dja",
			Email:       "dja@",
			PhoneNumber: "3",
		},
	}
	for _, value := range suite.UserS {
		suite.User, err = user.InsertUser(value.FullName, value.Email, value.PhoneNumber)
		if err != nil {
			suite.T().Fatal("Unable to run InsertUser store func")
		}
		suite.UserPs = append(suite.UserPs, suite.User)
	}
}
func (suite *UserTestSuite) TestGetUserById() {
	store := store.NewUserStoreModel(suite.Db)
	var err error
	var user []*entities.User
	for _, value := range suite.UserPs {
		user, err = store.GetUserById(value.Id)
		if err != nil {
			suite.T().Fatal("Unable to run GetUserById store func")
		}
	}
	for _, value := range suite.UserPs {
		for i := range user {
			if value.Id == user[i].Id {
				suite.Equal(value.Id, user[i].Id)
				suite.Equal(value.FullName, user[i].FullName)
				suite.Equal(value.Email, user[i].Email)
				suite.Equal(value.PhoneNumber, user[i].PhoneNumber)
			}
		}
	}
}
func (suite *UserTestSuite) TestUpdateUser(){
	store := store.NewUserStoreModel(suite.Db)
	var err error
	var user *entities.User
	var users []*entities.User
	time.Sleep(2*time.Second)
	for _, value := range suite.UserPs{
		user, err = store.UpdateUser(value.Id, value.FullName, value.Email, value.PhoneNumber)
		if err != nil{
			suite.T().Fatal("Unable to run UpdateUser store func")
		}
		users = append(users, user)
	}
	for _, value := range suite.UserPs{
		for i := range users {
			if value.Id == users[i].Id{
				suite.Equal(value.Id, users[i].Id)
				suite.Equal(value.FullName, users[i].FullName)
				suite.Equal(value.Email, users[i].Email)
				suite.Equal(value.PhoneNumber, users[i].PhoneNumber)
			}
		}
	}
}
func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
func (suite *UserTestSuite) TearDownTest() {
	store := store.NewUserStoreModel(suite.Db)
	var err error
	for _, value := range suite.UserPs{
		_, err = store.DeleteUser(value.Id)
		if err != nil{
			suite.T().Fatal("Unable to run DeleteUser store func")
		}
	}
}

