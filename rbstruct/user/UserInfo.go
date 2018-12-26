package user

//用户结构体
type UserInfo struct {
    lUserId uint32     //用户ID
	strNickName string //用户昵称
	strRealName string //用户真实姓名
	strPwd string      //用户密码
	roles[] RoleInfo   //角色列表
}
