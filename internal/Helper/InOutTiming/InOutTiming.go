package inouttiming

import (
	logger "smtpconnect/internal/Helper/Logger"
)

func InOutTiming(inTime string, outTime string, path string) {
	log := logger.InitLogger()
	log.Info("API Timing - ", "inTime: ", inTime, ", outTime: ", outTime, ", path: ", path)
}
