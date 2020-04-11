package model

type LoginReq struct {
	Username string `from:"username"`
	Password string `from:"password"`
}
