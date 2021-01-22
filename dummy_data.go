package main

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/cesarcedanov/social_expenses_tracker/model"
	"time"
)

var db *gorm.DB

var err error



var (

	users = []model.User{

		{Email: "alice@dummy.go", Password: "TEST"},
		{Email: "bob@dummy.go", Password: "TEST"},
		{ Email: "cheryl@dummy.go", Password: "TEST"},
	}

	expenses = []model.Expense{
		{
			UserID: 1,
			Amount: 100,
			Timestamp: time.Date(2021,time.January,21,0,0,0,0,time.UTC),
		},
		{
			UserID: 2,
			Amount: 50,
			Timestamp: time.Date(2021,time.January,21,0,0,0,0,time.UTC),
		},
		{
			UserID: 3,
			Amount: 250,
			Timestamp: time.Date(2021,time.January,21,0,0,0,0,time.UTC),
		},
	}

)
