// package routes

// import (
// 	"net/http"

// 	"github.com/duduBTW/karaokev2backend/controllers"
// 	middlewares "github.com/duduBTW/karaokev2backend/handlers"
// 	"github.com/gorilla/mux"
// )

// // Routes -> define endpoints
// func Routes() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/person", controllers.CreatePersonEndpoint).Methods("POST")
// 	router.HandleFunc("/auth", controllers.Auths).Methods("GET")
// 	router.HandleFunc("/people", middlewares.IsAuthorized(controllers.GetPeopleEndpoint)).Methods("GET")
// 	router.HandleFunc("/person/{id}", controllers.GetPersonEndpoint).Methods("GET")
// 	router.HandleFunc("/person/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
// 	router.HandleFunc("/person/{id}", controllers.UpdatePersonEndpoint).Methods("PUT")
// 	router.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")
// 	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploaded/"))))
// 	return router
// }

package routes

import (
	"net/http"

	"github.com/duduBTW/karaokev2backend/controllers"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/music", controllers.CreateMusicEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/musics/{section_id}", controllers.GetMusicsSectionEndpoint).Methods(http.MethodGet)

	return router
}
