package configurationservices

import (
	logger "smtpconnect/internal/Helper/Logger"
	configurationmodel "smtpconnect/model/Configuration"
	configurationvalidate "smtpconnect/validate/Configuration"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func CreateConfigurationService(db *gorm.DB, userId int, req configurationvalidate.ConfigReq) configurationvalidate.SingleConfigurationResponse {
	log := logger.InitLogger()

	// 1. Verify SMTP Connection first
	if !VerifySMTP(req.MailHost, req.MailPort, req.MailId, req.MailPassword) {
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "SMTP Verification Failed. Please check your credentials.",
			StatusCode: 400,
		}
	}

	// 2. Save to Database
	config := configurationmodel.ConfigurationModel{
		UserId:       userId,
		MailType:     req.MailType,
		MailId:       req.MailId,
		MailPassword: req.MailPassword,
		MailHost:     req.MailHost,
		MailPort:     req.MailPort,
		Status:       true,
	}

	if err := db.Create(&config).Error; err != nil {
		log.Error(err)
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "Error saving configuration to database",
			StatusCode: 500,
		}
	}

	return configurationvalidate.SingleConfigurationResponse{
		Status:     true,
		Message:    "Configuration added and verified successfully",
		StatusCode: 200,
		Data:       config,
	}
}

func GetAllConfigurationsService(db *gorm.DB, userId int) configurationvalidate.ConfigurationResponse {
	log := logger.InitLogger()

	var configs []configurationmodel.ConfigurationModel
	if err := db.Where("refuserid = ? AND delete_status = ?", userId, false).Find(&configs).Error; err != nil {
		log.Error(err)
		return configurationvalidate.ConfigurationResponse{
			Status:     false,
			Message:    "Error fetching configurations",
			StatusCode: 500,
		}
	}

	return configurationvalidate.ConfigurationResponse{
		Status:     true,
		Message:    "Configurations fetched successfully",
		StatusCode: 200,
		Data:       configs,
	}
}

func UpdateConfigurationService(db *gorm.DB, userId int, configId int, req configurationvalidate.ConfigReq) configurationvalidate.SingleConfigurationResponse {
	log := logger.InitLogger()

	// 1. Verify SMTP Connection first
	if !VerifySMTP(req.MailHost, req.MailPort, req.MailId, req.MailPassword) {
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "SMTP Verification Failed. Update cancelled.",
			StatusCode: 400,
		}
	}

	// 2. Find and Update
	var config configurationmodel.ConfigurationModel
	if err := db.Where("id = ? AND refuserid = ?", configId, userId).First(&config).Error; err != nil {
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "Configuration not found",
			StatusCode: 404,
		}
	}

	config.MailType = req.MailType
	config.MailId = req.MailId
	config.MailPassword = req.MailPassword
	config.MailHost = req.MailHost
	config.MailPort = req.MailPort
	config.Status = true

	if err := db.Save(&config).Error; err != nil {
		log.Error(err)
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "Error updating configuration",
			StatusCode: 500,
		}
	}

	return configurationvalidate.SingleConfigurationResponse{
		Status:     true,
		Message:    "Configuration updated and verified successfully",
		StatusCode: 200,
		Data:       config,
	}
}

func DeleteConfigurationService(db *gorm.DB, userId int, configId int) configurationvalidate.SingleConfigurationResponse {
	log := logger.InitLogger()

	var config configurationmodel.ConfigurationModel
	if err := db.Where("id = ? AND refuserid = ? AND delete_status = ?", configId, userId, false).First(&config).Error; err != nil {
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "Configuration not found",
			StatusCode: 404,
		}
	}

	config.DeleteStatus = true
	config.Status = false // also disable it

	if err := db.Save(&config).Error; err != nil {
		log.Error(err)
		return configurationvalidate.SingleConfigurationResponse{
			Status:     false,
			Message:    "Error deleting configuration",
			StatusCode: 500,
		}
	}

	return configurationvalidate.SingleConfigurationResponse{
		Status:     true,
		Message:    "Configuration deleted successfully",
		StatusCode: 200,
	}
}

func VerifySMTP(host string, port int, user string, password string) bool {
	log := logger.InitLogger()
	d := gomail.NewDialer(host, port, user, password)
	d.TLSConfig = nil // Let gomail handle STARTTLS or SSL/TLS based on port

	closer, err := d.Dial()
	if err != nil {
		log.Errorf("SMTP Verification Error for %s: %v", user, err)
		return false
	}
	closer.Close()
	return true
}
