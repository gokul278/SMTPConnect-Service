package usermodel

type UserprofileModel struct {
	UserId     int    `json:"refUserId" gorm:"column:refUserId"`
	UserCustId string `json:"refUserCustId" gorm:"column:refUserCustId"`
	UserName   string `json:"refUserName" gorm:"column:refUserName"`
	RTId       int    `json:"refRTId" gorm:"column:refRTId"`
}
