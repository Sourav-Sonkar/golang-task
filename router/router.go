package router

import (
	"assignment1/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", seveHome)
	r.HandleFunc("/createUser", controller.InsertUserController).Methods("POST")
	r.HandleFunc("/getUsers", controller.GetAllUsersController).Methods("GET")
	r.HandleFunc("/deleteUser/{id}", controller.DeleteUserController).Methods("DELETE")
	r.HandleFunc("/updateUser/{id}", controller.UpdateUserController).Methods("PUT")
	r.HandleFunc("/getUser/{id}", controller.GetUserController).Methods("GET")
	r.HandleFunc("/deleteAllUsers", controller.DeleteAllUsersController).Methods("DELETE")
	return r
}
func seveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Golang Task.Please go though api documention to use api(https://documenter.getpostman.com/view/19028114/UVXbuzHL)</h1>"))
}
