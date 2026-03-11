package authenticationmodel

type VerifyEmailPhoneNo struct {
	UserId             int    `json:"refUserId" gorm:"column:refUserId"`
	UserCustId         string `json:"refUserCustId" gorm:"column:refUserCustId"`
	UserName           string `json:"refUserName" gorm:"column:refUserName"`
	RTId               int    `json:"refRTId" gorm:"column:refRTId"`
	CODOEmail          string `json:"refUCMail" gorm:"column:refUCMail"`
	ADHashPass         string `json:"refUAPass" gorm:"column:refUAPass"`
	AHPassChangeStatus bool   `json:"refUAResetPassStatus" gorm:"column:refUAResetPassStatus"`
}

type SignupModel struct {
	Status     bool   `json:"status" gorm:"column:status"`
	Message    string `json:"message" gorm:"column:message"`
	StatusCode int    `json:"statuscode" gorm:"column:statuscode"`
}
