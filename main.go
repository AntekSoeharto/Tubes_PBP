package main

import (
	"fmt"
	"log"
	"net/http"

	ctrl "github.com/Tubes/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", ctrl.RegisterMember).Methods("POST")
	
	// router.HandleFunc("/users", cuser.Authenticate(cuser.GetAllUsers, 0)).Methods("GET")
	// router.HandleFunc("/users", cuser.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{user_id}", cuser.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/login", cuser.LogIn).Methods("GET")
	// router.HandleFunc("/logout", cuser.Logout).Methods("GET")

	// router.HandleFunc("/products", cuser.GetAllProducts).Methods("GET")
	// router.HandleFunc("/products", cuser.InsertProduct).Methods("POST")
	// router.HandleFunc("/products", cuser.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products/{product_id}", cuser.DeleteProduct).Methods("DELETE")

	// router.HandleFunc("/transactions", cuser.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/transactions", cuser.InsertTransaction).Methods("POST")
	// router.HandleFunc("/transactions", cuser.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/transactions/{transaction_id}", cuser.DeleteTransction).Methods("DELETE")

	//cors
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	http.Handle("/", handler)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
