package email

import (
	"net/smtp"

	util "github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

type EmailService struct {
	em   *email.Email
	auth smtp.Auth
	host string
}

var MailS EmailService

func NewEmailService() EmailService {
	SMTPHost := viper.GetString("emailservice.host")
	SMTPPort := viper.GetString("emailservice.port")
	SMTPUsername := viper.GetString("emailservice.username")
	SMTPPassword := viper.GetString("emailservice.password")
	em := email.NewEmail()
	em.From = SMTPUsername
	em.Subject = "FakeIns 邮件服务"

	ms := EmailService{
		em:   em,
		auth: smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost),
		host: SMTPHost + SMTPPort,
	}
	MailS = ms
	return MailS
}
func (s *EmailService) SendEmail(toEmail []string, subject string, content string) error {
	s.em.Text = []byte(content)
	s.em.Subject = subject
	s.em.To = toEmail

	err := s.em.Send(s.host, s.auth)
	return err
}
func (s *EmailService) SendValidCode(toEmail string) (error, string) {
	code := util.RandomStr(4)
	return s.SendEmail([]string{toEmail}, "FakeIns 邮箱验证码", "验证码："+code), code
}

func GetMailS() EmailService {
	return MailS
}
