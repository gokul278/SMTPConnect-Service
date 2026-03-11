package dashboardservice

import (
	dashboardvalidate "smtpconnect/validate/Dashboard"
	"time"
	"gorm.io/gorm"
)

func GetDashboardStats(db *gorm.DB, userId int) dashboardvalidate.DashboardResponse {

	var stats dashboardvalidate.DashboardStats

	// 1. Total, Sent, Failed Counts
	db.Table("mail.mail_history").Where("refuserid = ?", userId).Count(&stats.TotalEmails)
	db.Table("mail.mail_history").Where("refuserid = ? AND status = ?", userId, "sent").Count(&stats.SentEmails)
	db.Table("mail.mail_history").Where("refuserid = ? AND status = ?", userId, "failed").Count(&stats.FailedEmails)

	// 2. Success Rate
	if stats.TotalEmails > 0 {
		stats.SuccessRate = (float64(stats.SentEmails) / float64(stats.TotalEmails)) * 100
	}

	// 3. Active Configs
	db.Table("mail.configurations").Where("refuserid = ? AND status = ?", userId, true).Count(&stats.ActiveConfigs)

	// 4. Recent Activity (Last 7 Days)
	// We'll calculate the last 7 days manually to ensure every day is represented even with 0 emails
	var activity []dashboardvalidate.DailyStats
	now := time.Now()
	for i := 6; i >= 0; i-- {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		var sentCount int64
		var failedCount int64

		db.Table("mail.mail_history").
			Where("refuserid = ? AND status = ? AND DATE(sentat) = ?", userId, "sent", date).
			Count(&sentCount)
		
		db.Table("mail.mail_history").
			Where("refuserid = ? AND status = ? AND DATE(sentat) = ?", userId, "failed", date).
			Count(&failedCount)

		activity = append(activity, dashboardvalidate.DailyStats{
			Date:  now.AddDate(0, 0, -i).Format("02 Jan"), // User friendly format for chart
			Sent:  sentCount,
			Fixed: failedCount, // Using Fixed as failure count for chart mapping
		})
	}
	stats.RecentActivity = activity

	return dashboardvalidate.DashboardResponse{
		Status:     true,
		Message:    "Dashboard stats fetched successfully",
		StatusCode: 200,
		Data:       stats,
	}
}
