package main

import (
	"hygya-api/database"
	"hygya-api/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()

	// e := echo.New()

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	// 	AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	// }))

	// router := mux.NewRouter()
	// router.HandleFunc("/", testHandler).Methods("GET")
	// http.ListenAndServe(":5000",
	// 	handlers.CORS(
	// 		handlers.AllowedOrigins([]string{"*"}),
	// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	)(router))

}

// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "Hello World!"}`))
// }
