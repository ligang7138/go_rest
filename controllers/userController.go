package controllers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"go_rest/service"
	"net/http"
)
type User struct {

}

func New() *User {
	return &User{}
}

func (u *User) GetUserName(w rest.ResponseWriter, req *rest.Request) {
	//name,err := service.GetUserName(req.PathParam("id"))
	name,err := service.GetUserBySql(req.PathParam("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.WriteJson("nil")
		return
	}
	w.WriteJson(name)
}
