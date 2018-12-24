package controller

import (
	R "../../TQStruct/base"
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome!")
}

func Login(w http.ResponseWriter, r *http.Request)  {
	err := json.NewEncoder(w).Encode(R.OK())
	if  err != nil{
		panic(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request)  {
	err := json.NewEncoder(w).Encode(R.Error())
	if  err != nil{
		panic(err)
	}

}