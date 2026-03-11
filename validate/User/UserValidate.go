package uservalidate

import usermodel "smtpconnect/model/User"

type GetProfileResponse struct {
	StatusCode int                        `json:"statusCode"`
	Status     bool                       `json:"status"`
	Message    string                     `json:"message"`
	Data       usermodel.UserprofileModel `json:"data"`
}
