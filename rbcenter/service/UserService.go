package service

//定义用户结构体
type User struct {

}

func GetUser(mobile string) (User, int, string) {
	user:=User{}
   return user,0,""
}


func Login(mobile string,pwd string) (User, int, string) {
	user:=User{}
	return user,0,""
}