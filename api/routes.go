package main

import (
	"encoding/json"
	"fmt"
	"go_server/log"
	"go_server/models"
	"net/http"

	"github.com/gorilla/mux"
)

//Holds All the Different API Routes and Route Setup Functions

const API_VERSION = "v1"

type LocalRouter struct {
	*mux.Router
}

func setRoutes(router *LocalRouter) {

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/contact", ContactHandler).Methods("GET", "PUT")

}

// // // defining autentication struct
// type authenticationMiddleware struct {
// 	tokenUsers map[string]string
// }

// //initializing
// func (amw *authenticationMiddleware) Populate() {
// 	amw.tokenUsers["00000000"] = "user0"
// 	amw.tokenUsers["aaaaaaaa"] = "userA"
// 	amw.tokenUsers["05f717e5"] = "randomUser"
// 	amw.tokenUsers["deadbeef"] = "user0"
// }

// // Middleware function which will be called for each request
// func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
// 	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("X-Session-Token")

// 		if user, found := amw.tokenUsers[token]; found {
// 			// we found the token in our map
// 			log.Printf("Authenticated user &s\n", user)
// 			// pass down the request to the next middleware (or final handler)
// 			next.ServeHTTP(w, r)
// 		} else {
// 			// write an error and stop the handler chain
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 		}
// 	})
// }

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintln(w, "Contact page!")

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		//Will Eventually Add Some Authentication Business

		var c models.Contact
		selectedContacts, err := c.SelectAllContacts()
		if err != nil {
			//Write some http return code here usually gonna be some form of
			//500 in this case as it means we failed to go to the db

			// serving HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			log.Errorf("HTTP Server Error Return 500: %v", err)
		}
		json.NewEncoder(w).Encode(selectedContacts)

	case "PUT":

		var contacts []*models.Contact

		params := mux.Vars(r)
		for index, item := range contacts {
			if item.Name == params["name"] {
				contacts = append(contacts[:index], contacts[index+1:]...)

				var contact models.Contact
				_ = json.NewDecoder(r.Body).Decode(contact)
				contact.Name = append(contacts, contact)
				json.NewEncoder(w).Encode(&contact)

				return
			}
		}

		json.NewEncoder(w).Encode(contacts)

	} // end of switch method

} // end of ContactHandler

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to the home page")
	//w.Header().Set("Content-Type", "application/json")

	// probably not necessary here since it's just the home page, not the contact page

	// var c models.Contact
	// err := json.NewDecoder(r.Body).Decode(&c)

	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// json.NewEncoder(w).Encode(c)

}
