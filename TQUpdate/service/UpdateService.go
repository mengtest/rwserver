package service

import "../../TQBase/db"

func CheckVersion()  {
	db.DB.MustExec("")
}
