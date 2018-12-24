package controller

import (
	R "../../TQStruct/base"
	TQ "../../TQBase/base"
	TQC "../../TQBase/constant"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	url :=  TQC.LoginHost+ r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func Login(w http.ResponseWriter, r *http.Request)  {
	TQ.LogErr(json.NewEncoder(w).Encode(R.OK().OutLog()))
}

func Register(w http.ResponseWriter, r *http.Request)  {
	TQ.LogErr(json.NewEncoder(w).Encode(R.Error().OutLog()))
}