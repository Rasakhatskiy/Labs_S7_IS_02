package main

import (
	"OWL/rule_engine"
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/logger"
)

// can be part of user serice and a separate directory
type User struct {
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Age         int     `json:"age"`
	Gender      string  `json:"gender"`
	Income      float64 `json:"income"`
	HasCar      bool    `json:"has_car"`
	HasChildren bool    `json:"has_children"`
	LovesTravel bool    `json:"loves_travel"`
	CarType     string  `json:"has_sport_car"`
}

// can be moved to offer directory
type OfferService interface {
	CheckOfferForUser(user User) string
}

type OfferServiceClient struct {
	ruleEngineSvc *rule_engine.RuleEngineSvc
}

func NewOfferService(ruleEngineSvc *rule_engine.RuleEngineSvc) OfferService {
	return &OfferServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

func (svc OfferServiceClient) CheckOfferForUser(user User) string {
	offerCard := rule_engine.NewUserOfferContext()
	offerCard.UserOfferInput = &rule_engine.UserOfferInput{
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Gender:      user.Gender,
		Age:         user.Age,
		Income:      user.Income,
		HasCar:      user.HasCar,
		HasChildren: user.HasChildren,
		LovesTravel: user.LovesTravel,
		CarType:     user.CarType,
	}

	err := svc.ruleEngineSvc.Execute(offerCard)
	if err != nil {
		logger.Log.Error("get user offer rule engine failed", err)
	}

	return offerCard.UserOfferOutput.OfferType
}

func main() {
	ruleEngineSvc := rule_engine.NewRuleEngineSvc()
	offerSvc := NewOfferService(ruleEngineSvc)

	users := []User{
		{
			Name:        "Maksym Rasakhatskyi",
			Username:    "maksym",
			Email:       "me@google.com",
			Age:         30,
			Gender:      "male",
			Income:      10000,
			HasCar:      true,
			HasChildren: false,
			LovesTravel: false,
			CarType:     "slavuta",
		},
		{
			Name:        "Ivan",
			Username:    "pj",
			Email:       "pj@abc.com",
			Age:         50,
			Gender:      "male",
			Income:      500000,
			HasCar:      false,
			HasChildren: false,
			LovesTravel: false,
			CarType:     "",
		},
		{
			Name:        "Andrii",
			Username:    "aa",
			Email:       "aa@abc.com",
			Age:         18,
			Gender:      "male",
			Income:      500000,
			HasCar:      true,
			HasChildren: false,
			LovesTravel: false,
			CarType:     "sport",
		},
		{
			Name:        "Olga Kravets",
			Username:    "OK",
			Email:       "Ok@abc.com",
			Age:         50,
			Gender:      "female",
			Income:      5000,
			HasCar:      true,
			HasChildren: false,
			LovesTravel: false,
			CarType:     "",
		},
		{
			Name:        "Olena Biloboka",
			Username:    "OK",
			Email:       "Ok@abc.com",
			Age:         50,
			Gender:      "female",
			Income:      5000,
			HasCar:      false,
			HasChildren: true,
			LovesTravel: false,
			CarType:     "",
		},
	}

	for _, user := range users {
		fmt.Printf("offer for user %s: %s\n", user.Name, offerSvc.CheckOfferForUser(user))
	}
}
