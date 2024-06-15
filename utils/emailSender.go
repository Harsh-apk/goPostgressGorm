package utils

import (
	"net/smtp"

	"github.com/Harsh-apk/notesPostgres/types"
)

func SendSMS(email *string, data *types.ENV_DATA, link *string) error {
	msg := ("Subject: Account Verification\nClick on the link below to verify your account\n" + *link)

	//msg := "Subject: Verification code\nYour code is 1232123.\n\tCheck out github.com/Harsh-apk/goLangMailSender"
	auth := smtp.PlainAuth("", data.USER, data.PASSWORD, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, data.USER, []string{*email}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
