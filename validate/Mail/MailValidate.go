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
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	ConfigId     int    `json:"configId"`
	SenderEmail  string `json:"senderEmail"`
	Recipient    string `json:"recipient"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
	SentAt       string `json:"sentAt"`
}

type MailHistoryResponse struct {
	Status     bool                `json:"status"`
	Message    string              `json:"message"`
	StatusCode int                 `json:"statusCode"`
	Data       []MailHistoryDetail `json:"data"`
}
