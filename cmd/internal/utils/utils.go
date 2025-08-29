package utils 

import (
	"io"
	"os"
	"path/filepath"
)

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

	otpSource := "./cmd/internal/utils/otp.txt"
	tokenSource := "./cmd/internal/utils/token.txt"
	emailVerificationSource := "./cmd/internal/utils/emailVerification.txt"

	otpFile, err := os.Open(otpSource)
	if err != nil {
		return err
	}
	defer otpFile.Close()

	tokenFile, err := os.Open(tokenSource)
	if err != nil {
		return err
	}
	defer tokenFile.Close()

	emailVerificationFile, err := os.Open(emailVerificationSource)
	if err != nil {
		return err
	}
	defer emailVerificationFile.Close()

	otpPath := filepath.Join(utilsPath, "otp.go")
	otpFileDest, err := os.Create(otpPath)
	if err != nil {
		return err
	}
	defer otpFileDest.Close()

	_, err = io.Copy(otpFileDest, otpFile)
	if err != nil {
		return err
	}

	tokenPath := filepath.Join(utilsPath, "token.go")
	tokenFileDest, err := os.Create(tokenPath)
	if err != nil {
		return err
	}
	defer tokenFileDest.Close()

	_, err = io.Copy(tokenFileDest, tokenFile)
	if err != nil{
		return nil
	}

	sendEmailTxt := CreateSendEmailFile(projectName)

	emailVerificationPath := filepath.Join(emailsPath, "emailverification.go")
	emailVerificationDest, err := os.Create(emailVerificationPath)
	if err != nil{
		return err
	}
	defer emailVerificationDest.Close()

	_, err = io.Copy(emailVerificationDest, emailVerificationFile)
	if err != nil{
		return err
	}

	sendEmailPath := filepath.Join(emailsPath, "sendEmail.go")
	sendEmailDest, err := os.Create(sendEmailPath)
	if err != nil{
		return err 
	}
	defer sendEmailDest.Close()
	_, err = sendEmailDest.WriteString(sendEmailTxt)
	if err != nil{
		return err
	}

	return nil
}
