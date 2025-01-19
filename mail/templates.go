package mail

import (
	"fmt"
	"salimon/nexus/db"
	"salimon/nexus/types"
)

func SendRegisterVerificationEmail(user *types.User) {

	verification, err := db.InsertRegisterEmailVerification(user.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	verifyBaseUrl := "https://salimon.net/auth/register/verify"
	verifyUrl := fmt.Sprintf("%s?token=%s&email=%s", verifyBaseUrl, verification.Token, user.Email)

	body := fmt.Sprintf("Hello %s,\nwelcome to the salimon network.\n\nI am Nexus, your proxy to the network and entities. please verify you email in order to complete registeration.\n%s\n\nyou can click on this link to proceed. if you have not registered in salimon network. please ignore this email.", user.Username, verifyUrl)
	subject := "Registeration in Salimon"
	to := user.Email

	job := SendEmailJob{
		Body:    body,
		Subject: subject,
		To:      to,
	}
	EmailQueue <- job
}

func SendPasswordResetEmail(user *types.User) {
	verification, err := db.InsertPasswordResetVerification(user.Id)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	verifyBaseUrl := "http://salimon.net/auth/password-reset/verify"
	verifyUrl := fmt.Sprintf("%s/%s", verifyBaseUrl, verification.Token)

	subject := "Password Reset in Salimon"
	body := fmt.Sprintf("Hello %s,\nis seems you have requested to reset your password. please proceed your password reset progress from this link:\n%s\n\nif you have not requested for a password reset. you can safely ignore this email.", user.Username, verifyUrl)
	to := user.Email

	job := SendEmailJob{
		Body:    body,
		Subject: subject,
		To:      to,
	}

	EmailQueue <- job
}
