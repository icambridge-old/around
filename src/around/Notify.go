package around


import (
	"log"
	"net/smtp"
)

/**
 * Some people don't have access to campfire but may find some of the
 * notifications useful. So we'll notify them some how.
 */
func NotifyOtherPeople(message string) {


	// Set up authentication information.


	auth := smtp.PlainAuth(
		"",
		"",
		"",
		"smtp.sendgrid.net",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.

	body := "To: " + "test2@icambridge.me" + "\r\nFrom:" + "<iain.cambridge@workstars.com>" + "\r\nSubject: " + "shit is all broken" + "\r\n\r\n" + message

	err := smtp.SendMail(
		"smtp.sendgrid.net:25",
		auth,
		"<iain.cambridge@workstars.com>",
		[]string{"test2@icambridge.me"},
		[]byte(body),
	)
	if err != nil {
		log.Fatal(err)
	}
}
