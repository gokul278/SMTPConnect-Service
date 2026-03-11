package audit

import (
	logger "smtpconnect/internal/Helper/Logger"

	"gorm.io/gorm"
)

func AuditLogsInsert(db *gorm.DB, TransType int, Status string, TimeData string, UserId int, ActionBy int) bool {
	log := logger.InitLogger()

	var InsertAuditLogsSQL = `
INSERT INTO audit."TransHistory" ("transTypeId", "refTHData", "refTHTime", "refUserId", "refTHActionBy")
VALUES ($1, $2, $3, $4, $5);
`

	//Insert Audit Logs
	InsertAuditLogsErr := db.Exec(
		InsertAuditLogsSQL, //Query
		TransType,          //$1
		Status,             //$2
		TimeData,           //$3
		UserId,             //$4
		ActionBy,           //$5
	).Error
	if InsertAuditLogsErr != nil {
		log.Printf("ERROR: Failed to insert audit logs: %v\n", InsertAuditLogsErr)
		return false
	}

	return true
}
