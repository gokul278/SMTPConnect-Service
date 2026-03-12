package mailvalidate

import mailmodel "smtpconnect/model/Mail"

type SendMailReq struct {
	ConfigId  int    `json:"configId" validate:"required"`
	Recipient string `json:"recipient" validate:"required,email"`
	Subject   string `json:"subject" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

type SendMailResponse struct {
	Status     bool                 `json:"status"`
	Message    string               `json:"message"`
	StatusCode int                  `json:"statusCode"`
	Data       mailmodel.MailHistoryModel `json:"data"`
}

type MailHistoryDetail struct {
	Id           int    `json:"id" gorm:"column:id"`
	UserId       int    `json:"userId" gorm:"column:refuserid"`
	ConfigId     int    `json:"configId" gorm:"column:refconfigid"`
	SenderEmail  string `json:"senderEmail" gorm:"column:sender_email"`
	Recipient    string `json:"recipient" gorm:"column:recipient"`
	Subject      string `json:"subject" gorm:"column:subject"`
	Content      string `json:"content" gorm:"column:content"`
	Status       string `json:"status" gorm:"column:status"`
	ErrorMessage string `json:"errorMessage" gorm:"column:errormessage"`
	SentAt       string `json:"sentAt" gorm:"column:sentat"`
}

type MailHistoryResponse struct {
	Status     bool                `json:"status"`
	Message    string              `json:"message"`
	StatusCode int                 `json:"statusCode"`
	Data       []MailHistoryDetail `json:"data"`
}
