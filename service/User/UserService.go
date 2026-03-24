package userservices

import (
	logger "mailstitch/internal/Helper/Logger"
	usermodel "mailstitch/model/User"
	userquery "mailstitch/query/User"
	uservalidate "mailstitch/validate/User"

	"gorm.io/gorm"
)

func UserProfileService(db *gorm.DB, idVal int) uservalidate.GetProfileResponse {
	log := logger.InitLogger()

	var GetUserProfileData usermodel.UserprofileModel
	GetUserProfileDataErr := db.Raw(
		userquery.GetUserProfileQuery, //query
		idVal,                         //$1
	).Scan(&GetUserProfileData).Error

	if GetUserProfileDataErr != nil {
		log.Error(GetUserProfileDataErr)
		return uservalidate.GetProfileResponse{
			Status:     false,
			Message:    "Error fetching user profile",
			StatusCode: 500,
		}
	}

	return uservalidate.GetProfileResponse{
		Status:     true,
		Message:    "User profile fetched successfully",
		StatusCode: 200,
		Data:       GetUserProfileData,
	}
}
