package mailservice

import (
	logger "smtpconnect/internal/Helper/Logger"
	configurationmodel "smtpconnect/model/Configuration"
	mailmodel "smtpconnect/model/Mail"
	mailvalidate "smtpconnect/validate/Mail"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SendMailService(db *gorm.DB, userId int, req mailvalidate.SendMailReq) mailvalidate.SendMailResponse {
	log := logger.InitLogger()

	// 1. Fetch SMTP Configuration
	var config configurationmodel.ConfigurationModel
	if err := db.Where("id = ? AND refuserid = ?", req.ConfigId, userId).First(&config).Error; err != nil {
		return mailvalidate.SendMailResponse{
			Status:     false,
			Message:    "SMTP Configuration not found",
			StatusCode: 404,
		}
	}

	// 2. Prepare Email
	m := gomail.NewMessage()
	m.SetHeader("From", config.MailId)
	m.SetHeader("To", req.Recipient)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", req.Content)

	d := gomail.NewDialer(config.MailHost, config.MailPort, config.MailId, config.MailPassword)
	d.TLSConfig = nil // gomail handles TLS/STARTTLS

	// 3. Send and Record History
	history := mailmodel.MailHistoryModel{
		UserId:    userId,
		ConfigId:  req.ConfigId,
		Recipient: req.Recipient,
		Subject:   req.Subject,
		Content:   req.Content,
	}

	err := d.DialAndSend(m)
	if err != nil {
		log.Errorf("Failed to send mail: %v", err)
		history.Status = "failed"
		history.ErrorMessage = err.Error()
	} else {
		history.Status = "sent"
	}

	// 4. Save History to DB
	if err := db.Create(&history).Error; err != nil {
		log.Errorf("Failed to save mail history: %v", err)
	}

	if history.Status == "failed" {
		return mailvalidate.SendMailResponse{
			Status:     false,
			Message:    "Failed to send email: " + history.ErrorMessage,
			StatusCode: 500,
			Data:       history,
		}
	}

	return mailvalidate.SendMailResponse{
		Status:     true,
		Message:    "Email sent successfully",
		StatusCode: 200,
		Data:       history,
	}
}

func GetMailHistoryService(db *gorm.DB, userId int) mailvalidate.MailHistoryResponse {
	log := logger.InitLogger()
	var history []mailvalidate.MailHistoryDetail

	err := db.Table("mail.mail_history").
		Select("mail.mail_history.*, mail.configurations.mailid as sender_email").
		Joins("inner join mail.configurations on mail.configurations.id = mail.mail_history.refconfigid").
		Where("mail.mail_history.refuserid = ?", userId).
		Order("mail.mail_history.sentat desc").
		Scan(&history).Error

	if err != nil {
		log.Error(err)
		return mailvalidate.MailHistoryResponse{
			Status:     false,
			Message:    "Failed to fetch mail history",
			StatusCode: 500,
		}
	}

	return mailvalidate.MailHistoryResponse{
		Status:     true,
		Message:    "Mail history fetched successfully",
		StatusCode: 200,
		Data:       history,
	}
}
