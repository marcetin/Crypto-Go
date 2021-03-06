package api

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/mitsukomegumi/Crypto-Go/src/accounts"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SetupAccountRoutes - setup necessary routes for accout database
func SetupAccountRoutes(db *mgo.Database) (*fasthttprouter.Router, error) {
	router, pErr := setPosts(db)

	if pErr != nil {
		return router, pErr
	}

	_, err := setGets(router, db)

	if err != nil {
		return router, err
	}

	return router, nil
}

func setGets(initRouter *fasthttprouter.Router, db *mgo.Database) (*fasthttprouter.Router, error) {
	req, err := NewRequestServer(":username", "/api/accounts", "GET", db, db, "username")

	if err != nil {
		return nil, err
	}

	router, err := req.AttemptToServeRequestsWithRouter(initRouter)

	if err != nil {
		return nil, err
	}

	return router, nil
}

func setPosts(db *mgo.Database) (*fasthttprouter.Router, error) {
	postReq, rErr := NewRequestServer("POST", "/api/accounts", "POST", nil, db, "/:username/:email/:pass")

	if rErr != nil {
		return nil, rErr
	}

	router, pErr := postReq.AttemptToServeRequests()

	if pErr != nil {
		panic(rErr)
	}

	return router, nil
}

func addAccount(database *mgo.Database, account *accounts.Account) error {

	_, err := findAccount(database, account.Username)

	if err != nil {
		c := database.C("accounts")

		iErr := c.Insert(account)

		if iErr != nil {
			return iErr
		}

		return nil
	}
	return nil
}

func findAccount(database *mgo.Database, username string) (*accounts.Account, error) {
	c := database.C("accounts")

	result := accounts.Account{}

	err := c.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
