package utils

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed otp.txt
var otpTemplate string

//go:embed token.txt
var tokenTemplate string

//go:embed emailVerification.txt
var emailVerificationTemplate string

func HandleUtils(projectName string) error {
	utilsPath := filepath.Join(projectName, "internal", "utils")
	err := os.MkdirAll(utilsPath, os.ModePerm)
	if err != nil {
		return err
	}

	emailsPath := filepath.Join(projectName, "internal", "utils", "emails")
	err = os.Mkdir(emailsPath, os.ModePerm)
	if err != nil {
		return err
	}

	otpPath := filepath.Join(utilsPath, "otp.go")
	if err := os.WriteFile(otpPath, []byte(otpTemplate), 0644); err != nil {
		return err
	}

	tokenPath := filepath.Join(utilsPath, "token.go")
	if err := os.WriteFile(tokenPath, []byte(tokenTemplate), 0644); err != nil {
		return err
	}

	sendEmailTxt := CreateSendEmailFile(projectName)

	emailVerificationPath := filepath.Join(emailsPath, "emailverification.go")
	if err := os.WriteFile(emailVerificationPath, []byte(emailVerificationTemplate), 0644); err != nil {
		return err
	}

	sendEmailPath := filepath.Join(emailsPath, "sendEmail.go")
	if err := os.WriteFile(sendEmailPath, []byte(sendEmailTxt), 0644); err != nil {
		return err
	}

	return nil
}
