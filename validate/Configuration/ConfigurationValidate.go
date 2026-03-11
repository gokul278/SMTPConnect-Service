package configurationvalidate

import configurationmodel "smtpconnect/model/Configuration"

type ConfigReq struct {
	MailType     string `json:"mailType" validate:"required"`
	MailId       string `json:"mailId" validate:"required,email"`
	MailPassword string `json:"mailPassword" validate:"required"`
	MailHost     string `json:"mailHost" validate:"required"`
	MailPort     int    `json:"mailPort" validate:"required"`
}

type ConfigurationResponse struct {
	StatusCode int                                      `json:"statusCode"`
	Status     bool                                     `json:"status"`
	Message    string                                   `json:"message"`
	Data       []configurationmodel.ConfigurationModel `json:"data"`
}

type SingleConfigurationResponse struct {
	StatusCode int                                     `json:"statusCode"`
	Status     bool                                    `json: "status"`
	Message    string                                  `json: "message"`
	Data       configurationmodel.ConfigurationModel `json: "data"`
}
