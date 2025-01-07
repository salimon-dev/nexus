package mail

import (
	"fmt"
	"salimon/proxy/db"
	"salimon/proxy/types"
)

func SendRegisterVerificationEmail(user *types.User) {

	verification, err := db.InsertRegisterEmailVerification(user.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	verifyBaseUrl := "https://salimon.net/auth/verify"
	verifyUrl := fmt.Sprintf("%s/%s", verifyBaseUrl, verification.Token)

	body := fmt.Sprintf("Hello %s,\nwelcome to the salimon network.\n\nI am Nexus, your proxy to the network and entities. please verify you email in order to complete registeration.\n\n%s\n\nyou can click on this link to proceed. if you have not registered in salimon network. please ignore this email.", user.Username, verifyUrl)
	subject := "welcome to salimon"
	to := user.Email

	job := SendEmailJob{
		Body:    body,
		Subject: subject,
		To:      to,
	}
	EmailQueue <- job
}
