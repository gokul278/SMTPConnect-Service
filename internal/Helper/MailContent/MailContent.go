package mailcontent

import (
	timeZone "smtpconnect/internal/Helper/TimeZone"
	"html"
	"os"
	"strconv"
	"time"
)

func RegistrationMailContent(userName, patientID, gmail, password string) string {
	return `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .credentials {
        background-color: #fff;
        padding: 15px;
        border-radius: 5px;
        margin: 20px auto;
        width: fit-content;
        text-align: left;
        font-family: monospace;
        border: 1px solid #ccc;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Welcome, ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
        <p>
          You have successfully been registered on <strong>Ease Disease</strong>.
        </p>
        <p>Your login credentials are as follows:</p>
        <div class="credentials">
          <p>
            <strong>ID:</strong> ` +
		html.EscapeString(patientID) + `
          </p>
          <p>
            <strong>Email:</strong> ` + html.EscapeString(gmail) + `
          </p>
          <p><strong>Password:</strong> ` + html.EscapeString(password) + `</p>
        </div>
        <a href="` + os.Getenv("ACCESSURL") + `" class="button"
          >Login Now</a
        >
        <p style="text-align: justify;margin-top: 20px">This email is intended only for the individual or entity to which it is addressed, and may contain information that is privileged, confidential, and exempt from disclosure under applicable law. If the reader of this message is not the intended recipient, or the employee or agent responsible for delivering the message to the intended recipient, you are hereby informed that any use, disclosure, distribution, or copying of this communication is strictly prohibited. If you have received this communication in error, please notify us immediately by telephone and delete the original email message.</p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Ease Disease. All rights reserved.
      </div>
    </div>
  </body>
</html>
`
}

func ForgetPasswordMailContent(userName string, otp int) string {
	return `
  <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .credentials {
        background-color: #fff;
        padding: 15px;
        border-radius: 5px;
        margin: 20px auto;
        width: fit-content;
        text-align: left;
        font-family: monospace;
        border: 1px solid #ccc;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Welcome, ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
        <p>
          Please use the <b>Ease Disease</b> verification code below to reset your password. This code will remain valid for 10 minutes:
        </p>
         <h1>` + strconv.Itoa(otp) + `</h1>
        <p style="text-align: justify;margin-top: 20px">This email is intended only for the individual or entity to which it is addressed, and may contain information that is privileged, confidential, and exempt from disclosure under applicable law. If the reader of this message is not the intended recipient, or the employee or agent responsible for delivering the message to the intended recipient, you are hereby informed that any use, disclosure, distribution, or copying of this communication is strictly prohibited. If you have received this communication in error, please notify us immediately by telephone and delete the original email message.</p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Ease Disease. All rights reserved.
      </div>
    </div>
  </body>
</html>
  `
}

func EnrollCourseMailContent(userName string, courseName string, urlLink string) string {
	return `
  <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .credentials {
        background-color: #fff;
        padding: 15px;
        border-radius: 5px;
        margin: 20px auto;
        width: fit-content;
        text-align: left;
        font-family: monospace;
        border: 1px solid #ccc;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Welcome, ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
          <p>Greetings from WellthGreen Foundation.</p>
        <p>
          Thank you for enrolling in our course on <b>Ease Disease</b>. Your course access has been activated.
        </p>
        <div class="credentials">
          <p><strong>Course Name:</strong> ` + html.EscapeString(courseName) + `</p>
        </div>
        <a href="` + urlLink + `" class="button"
          >View Course</a
        >
        <p style="text-align: justify;margin-top: 20px">This email is intended only for the individual or entity to which it is addressed, and may contain information that is privileged, confidential, and exempt from disclosure under applicable law. If the reader of this message is not the intended recipient, or the employee or agent responsible for delivering the message to the intended recipient, you are hereby informed that any use, disclosure, distribution, or copying of this communication is strictly prohibited. If you have received this communication in error, please notify us immediately by telephone and delete the original email message.</p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Ease Disease. All rights reserved.
      </div>
    </div>
  </body>
</html>
  `
}
