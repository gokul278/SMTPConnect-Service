package authenticationservices

import (
	"net/http"
	accesstoken "smtpconnect/internal/Helper/AccessToken"
	audit "smtpconnect/internal/Helper/Audit"
	becrypt "smtpconnect/internal/Helper/Becrypt"
	logger "smtpconnect/internal/Helper/Logger"
	timeZone "smtpconnect/internal/Helper/TimeZone"
	authenticationmodel "smtpconnect/model/Authentication"
	authenticationquery "smtpconnect/query/Authentication"
	authenticationvalidate "smtpconnect/validate/Authentication"

	"gorm.io/gorm"
)

func SignInService(db *gorm.DB, reqVal authenticationvalidate.LoginReq) authenticationvalidate.LoginResponse {
	log := logger.InitLogger()

	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("ERROR: Failed to begin transaction: %v\n", tx.Error)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	defer func() {
		if r := recover(); r != nil {
			log.Error("ERROR: Recovered from panic, rolling back transaction:", r)
			tx.Rollback()
		}
	}()

	var VerifyEmailPhoneNo []authenticationmodel.VerifyEmailPhoneNo

	// Execute the raw SQL query with the username (phone number / mobile number)
	result := tx.Raw(authenticationquery.VerifyEmailPhoneNoSQL, reqVal.Username).Scan(&VerifyEmailPhoneNo)
	if result.Error != nil {
		log.Error("Error in SigninService while fetching user details: ", result.Error)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	// Check if any user found
	if len(VerifyEmailPhoneNo) == 0 {
		log.Info("No user found with the given username: ", reqVal.Username)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Invalid username or password",
			StatusCode: http.StatusUnauthorized,
		}
	}

	// Password verification
	user := VerifyEmailPhoneNo[0]
	match := becrypt.ComparePasswords(user.ADHashPass, reqVal.Password)

	if !match {
		log.Warn("LoginService Invalid Credentials(p) for Username: " + reqVal.Username)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Invalid username or password",
			StatusCode: http.StatusUnauthorized,
		}
	}

	//Insert Audit Logs
	AuditStatus := audit.AuditLogsInsert(tx, 1, "Signed Successfully", timeZone.GetPacificTime(), user.UserId, user.UserId)
	if !AuditStatus {
		tx.Rollback()
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("ERROR: Failed to commit transaction: %v\n", err)
		tx.Rollback()
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return authenticationvalidate.LoginResponse{
		Status:           true,
		Message:          "Successfully Logged In",
		RoleId:           VerifyEmailPhoneNo[0].RTId,
		PassChangeStatus: VerifyEmailPhoneNo[0].AHPassChangeStatus,
		Token:            accesstoken.CreateToken(user.UserId, user.RTId),
		StatusCode:       http.StatusOK,
	}
}

func SignUpService(db *gorm.DB, reqVal authenticationvalidate.SignupReq) authenticationvalidate.LoginResponse {
	log := logger.InitLogger()

	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("ERROR: Failed to begin transaction: %v\n", tx.Error)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	defer func() {
		if r := recover(); r != nil {
			log.Error("ERROR: Recovered from panic, rolling back transaction:", r)
			tx.Rollback()
		}
	}()

	//Hashpassword
	hashedPassword, err := becrypt.HashPassword(reqVal.Password)
	if err != nil {
		log.Error("Error in SignUpService while hashing password: ", err)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	// Check if email already exists or Create new user
	var SignupData authenticationmodel.SignupModel
	result := tx.Raw(
		authenticationquery.SignUpSQL, //query
		1,                             //$1
		reqVal.Name,                   //$2
		reqVal.Email,                  //$3
		hashedPassword,                //$4
		timeZone.GetPacificTime(),     //$5
	).Scan(&SignupData)
	if result.Error != nil {
		log.Error("Error in SignUpService while checking existing email: ", result.Error)
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("ERROR: Failed to commit transaction: %v\n", err)
		tx.Rollback()
		return authenticationvalidate.LoginResponse{
			Status:     false,
			Message:    "Something went wrong, please try again",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return authenticationvalidate.LoginResponse{
		Status:     SignupData.Status,
		Message:    SignupData.Message,
		StatusCode: SignupData.StatusCode,
	}
}
