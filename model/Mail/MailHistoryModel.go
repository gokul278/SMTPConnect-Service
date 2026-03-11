package mailmodel

import (
	"time"
)

type MailHistoryModel struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserId       int       `json:"refUserId" gorm:"column:refuserid"`
	ConfigId     int       `json:"refConfigId" gorm:"column:refconfigid"`
	Recipient    string    `json:"recipient" gorm:"column:recipient"`
	Subject      string    `json:"subject" gorm:"column:subject"`
	Content      string    `json:"content" gorm:"column:content"`
	Status       string    `json:"status" gorm:"column:status"` // 'sent', 'failed'
	ErrorMessage string    `json:"errorMessage" gorm:"column:errormessage"`
	SentAt       time.Time `json:"sentAt" gorm:"column:sentat;autoCreateTime"`
}

func (MailHistoryModel) TableName() string {
	return "mail.mail_history"
}
