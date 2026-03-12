package configurationmodel

import "time"

type ConfigurationModel struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserId       int       `json:"refUserId" gorm:"column:refuserid"`
	MailType     string    `json:"mailType" gorm:"column:mailtype"`
	MailId       string    `json:"mailId" gorm:"column:mailid"`
	MailPassword string    `json:"mailPassword" gorm:"column:mailpassword"`
	MailHost     string    `json:"mailHost" gorm:"column:mailhost"`
	MailPort     int       `json:"mailPort" gorm:"column:mailport"`
	Status       bool      `json:"status" gorm:"column:status;default:false"`
	DeleteStatus bool      `json:"deleteStatus" gorm:"column:delete_status;default:false"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:createdat;autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updatedat;autoUpdateTime"`
}

func (ConfigurationModel) TableName() string {
	return "mail.configurations"
}
