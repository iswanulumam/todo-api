package auth

type LoginRequestFormat struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Token string `json:"token"`
}
