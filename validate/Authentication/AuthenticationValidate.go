package authenticationvalidate

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	StatusCode       int    `json:"statusCode"`
	Status           bool   `json:"status"`
	Message          string `json:"message"`
	RoleId           int    `json:"roleId"`
	Token            string `json:"token"`
	PassChangeStatus bool   `json:"passChangeStatus"`
}

type SignupReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     bool   `json:"status"`
	Message    string `json:"message"`
}
