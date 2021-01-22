package main

import (
	"encoding/json"
	"gitlab.com/cesarcedanov/social_expenses_tracker/controller"
	"gitlab.com/cesarcedanov/social_expenses_tracker/model"
	"log"
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to Get all the Expenses")
	expenses, err := controller.GetExpenses(db)
	if err != nil {
		log.Print(err.Error())
		return
	}

	json.NewEncoder(w).Encode(&expenses)
}

func CreateExpenses(w http.ResponseWriter, r *http.Request){
	log.Println("Request to Create an Expense")
	var expense model.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	tokenAuth, err := ExtractTokenMetadata(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	expense.UserID = uint(userId)
	newExpense, err := controller.CreateExpense(db, &expense)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newExpense)
}



func GetBalance(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting a balance" )
	tokenAuth, err := ExtractTokenMetadata(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	balance, err := controller.GetBalance(db, uint(userId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&balance)

}

func SignUp(w http.ResponseWriter, r *http.Request){
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		// TODO: Validate Email format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty Email - Check payload"))

		return
	}
	if len(user.Password) < 8{
		// Invalid Password - Password must have at least 8 characters
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Password - Password must have at least 8 characters"))

		return
	}
	newUser, err := controller.CreateUser(db, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	token, err := CreateToken(newUser.ID)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	saveErr := CreateAuth(user.ID, token)
	if saveErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(saveErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&token)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: Check user in the DB
	//compare the user from the request, with the one we defined:
	if 1 == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	saveErr := CreateAuth(user.ID, token)
	if saveErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(saveErr.Error())
		return
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tokens)
}


func Logout(w http.ResponseWriter, req *http.Request) {
	au, err := ExtractTokenMetadata(req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	deleted, delErr := DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Successfully logged out")
	return
}