package main

import (
	"github.com/gorilla/mux"
	"github.com/didip/tollbooth"
)

func addRoutes(r *mux.Router){

	// API Endpoints
	r.Handle("/expense", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil),
		TokenAuthMiddleware(GetExpenses))).Methods("GET")
	r.Handle("/expense", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil),
		TokenAuthMiddleware(CreateExpenses))).Methods("POST")
	r.Handle("/user", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil),
		TokenAuthMiddleware(GetExpenses))).Methods("GET")
	r.Handle("/user/balance", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil),
		TokenAuthMiddleware(GetBalance))).Methods("GET")


	// Authentication Endpoints
	r.HandleFunc("/signup", SignUp).Methods("POST")
	r.HandleFunc("/login", Login).Methods("PUT")
	r.HandleFunc("/logout", Logout).Methods("PUT")
}