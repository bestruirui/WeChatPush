package mail

import (
	"crypto/tls"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

var (
	from       mail.Address
	to         mail.Address
	smtpServer string
	smtpPort   string
	username   string
	password   string
	auth       smtp.Auth
	tlsConfig  *tls.Config
	smtpClient *smtp.Client
)

func init() {
	godotenv.Load(".env")

	from = mail.Address{
		Name:    os.Getenv("FROM_NAME"),
		Address: os.Getenv("FROM_ADDRESS"), // 发件人邮箱
	}
	to = mail.Address{
		Name:    os.Getenv("TO_NAME"),
		Address: os.Getenv("TO_ADDRESS"), // 收件人邮箱
	}
	smtpServer = os.Getenv("SMTP_SERVER") // SMTP 服务器地址
	smtpPort = os.Getenv("SMTP_PORT")     // SMTP 服务器端口号
	username = os.Getenv("USERNAME")      // 发件人邮箱用户名
	password = os.Getenv("PASSWORD")      // 发件人邮箱密码

}
func SendEmail(name string, content string) {

	auth = smtp.PlainAuth("", username, password, smtpServer)

	tlsConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	conn, _ := tls.Dial("tcp", smtpServer+":"+smtpPort, tlsConfig)
	smtpClient, _ = smtp.NewClient(conn, smtpServer)
	smtpClient.Auth(auth)
	smtpClient.Mail(from.Address)
	smtpClient.Rcpt(to.Address)

	msg := "From: " + from.String() + "\r\n" +
		"To: " + to.String() + "\r\n" +
		"Subject: " + name + ":" + content

	w, _ := smtpClient.Data()
	defer w.Close()

	w.Write([]byte(msg))
}
