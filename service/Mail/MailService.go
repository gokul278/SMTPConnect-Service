package mailservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	logger "smtpconnect/internal/Helper/Logger"
	configurationmodel "smtpconnect/model/Configuration"
	mailmodel "smtpconnect/model/Mail"
	mailvalidate "smtpconnect/validate/Mail"

	"gorm.io/gorm"
)

func SendMailService(db *gorm.DB, userId int, req mailvalidate.SendMailReq) mailvalidate.SendMailResponse {
	log := logger.InitLogger()

	// 1. Fetch SMTP Configuration
	var config configurationmodel.ConfigurationModel
	if err := db.Where("id = ? AND refuserid = ?", req.ConfigId, userId).First(&config).Error; err != nil {
		return mailvalidate.SendMailResponse{
			Status: false, Message: "SMTP Configuration not found", StatusCode: 404,
		}
	}

	// 2. Prepare Payload for Cloudflare Worker
	proxyURL := "https://smtpconnectservice.gokulhk278.workers.dev"
	payload := map[string]interface{}{
		"to":       req.Recipient,
		"from":     config.MailId,
		"subject":  req.Subject,
		"content":  req.Content,
		"host":     config.MailHost,
		"port":     config.MailPort,
		"password": config.MailPassword, // Note: Worker must handle this
	}

	jsonData, _ := json.Marshal(payload)

	// 3. Send via HTTP POST (Bypasses Render's SMTP Block)
	resp, err := http.Post(proxyURL, "application/json", bytes.NewBuffer(jsonData))

	history := mailmodel.MailHistoryModel{
		UserId: userId, ConfigId: req.ConfigId, Recipient: req.Recipient,
		Subject: req.Subject, Content: req.Content,
	}

	if err != nil || resp.StatusCode != 200 {
		errorMessage := "Proxy connection failed"
		if err != nil {
			errorMessage = err.Error()
		}
		log.Errorf("Failed to send mail via proxy: %v", errorMessage)
		history.Status = "failed"
		history.ErrorMessage = errorMessage
	} else {
		history.Status = "sent"
	}

	// 4. Save History and Return
	db.Create(&history)

	if history.Status == "failed" {
		return mailvalidate.SendMailResponse{
			Status: false, Message: "Failed via Proxy: " + history.ErrorMessage, StatusCode: 500,
		}
	}

	return mailvalidate.SendMailResponse{
		Status: true, Message: "Email sent via Cloudflare Proxy", StatusCode: 200,
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
