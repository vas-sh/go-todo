package mail

import (
	"bytes"
	"net/smtp"
)

const emailTemplate = `
		<html>
		<body style="margin:0; padding:0; background-color: #6a69e0;">
		  <table width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor=" #f4f4f4">
		    <tr>
		      <td align="center">
		        <table width="600" cellpadding="0" cellspacing="0" border="0" bgcolor="#ffffff" 
					   style="border-radius:8px; overflow:hidden;">
		          <tr>
		            <td align="center" bgcolor=" #222c41" 
						style="padding: 18px; color: #f8c345; font-family: Arial, sans-serif; font-size: 28px;">
		              <strong>Todo List</strong>
		            </td>
		          </tr>
		          <tr>
		            <td style="padding: 20px; font-family: Arial, sans-serif; font-size: 20px; color: #000000;">
		              <p>Hi! Thank you for choosing Todo List!</p>
		              <p>If it was not you, ignore the message.</p>
					  <p>To activate your account, click the button below.</p>
		              <p style="text-align: center; margin: 20px 0;">
		                <a href="{{.Link}}" 
						   style="background-color: #f8c345;
						   		  color: #000000; text-decoration: none; padding: 10px 20px;
						          border-radius: 20px; display: inline-block;">
		                  Finish verifying
		                </a>
		              </p>
		            </td>
		          </tr>
		          <tr>
		            <td bgcolor=" #48576d" style="padding: 10px;">
		            </td>
		          </tr>
		        </table>
		      </td>
		    </tr>
		  </table>
		</body>
		</html>
	`

func (s *srv) Send(to, subject, link string) error {
	var buf bytes.Buffer
	if err := s.templ.Execute(&buf, struct{ Link string }{link}); err != nil {
		return err
	}
	body := buf.String()
	msg := "From: " + s.login + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		body
	return smtp.SendMail(s.host+":"+s.port,
		smtp.PlainAuth("", s.login, s.password, s.host),
		s.login, []string{to}, []byte(msg))
}
