package tests

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"bankacc/pkg/entities"
	"bankacc/pkg/store"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
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
	User      entities.User
	Users     []entities.User
	Userstore store.UserModel
	Db        *sql.DB
}

func (suite *UserTestSuite) SetupTest() {
	var err error
	suite.Db = MySQLInit()
	user := store.NewUserStoreModel(suite.Db)
	suite.Users = []entities.User{
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
	for i, current := range suite.Users {
		suite.User, err = user.InsertUser(current.FullName, current.Email, current.PhoneNumber)
		if err != nil {
			suite.T().Fatal("Unable to run InsertUser store func")
		}
		suite.Users[i] = suite.User

	}
}

func (suite *UserTestSuite) TestGetUserById() {
	store := store.NewUserStoreModel(suite.Db)
	var err error
	var user []entities.User
	var expectedUser []entities.User
	for _, current := range suite.Users {
		user, err = store.GetUserById(current.Id)
		if err != nil {
			suite.T().Fatal("Unable to run GetUserById store func")
		}
		expectedUser = append(user, current)
		suite.Equal(expectedUser, user, "Checking Users" )
	}
}

func (suite *UserTestSuite) TestUpdateUser() {
	store := store.NewUserStoreModel(suite.Db)
	var err error
	var user entities.User
		time.Sleep(2 * time.Second)
	for _, current := range suite.Users {
		user, err = store.UpdateUser(current.Id, current.FullName, current.Email, current.PhoneNumber)
		if err != nil {
			suite.T().Fatal("Unable to run UpdateUser store func")
		}
		suite.Equal(current, user)
	}
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) TearDownTest() {
	store := store.NewUserStoreModel(suite.Db)
	var err error
	for _, current := range suite.Users {
		_, err = store.DeleteUser(current.Id)
		if err != nil {
			suite.T().Fatal("Unable to run DeleteUser store func")
		}
	}
}
