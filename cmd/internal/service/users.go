package service

import (
	"os"
	"path/filepath"
)

func HanldeService(projectName, modulePath string) error {
	serviceDirPath := filepath.Join(projectName, "internal", "service", "users")
	err := os.MkdirAll(serviceDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	storePath := filepath.Join(serviceDirPath, "store.go")
	storeFileDest, err := os.Create(storePath)
	if err != nil{
		return err
	}
	defer storeFileDest.Close()

	_, err = storeFileDest.WriteString(GenerateStore(modulePath))
	if err != nil{
		return err
	}

	servicePath := filepath.Join(serviceDirPath, "service.go")
	serviceFileDest, err := os.Create(servicePath)
	if err != nil{
		return err
	}
	defer serviceFileDest.Close()

	_, err = serviceFileDest.WriteString(GenerateService(modulePath))
	if err != nil{
		return err
	}

	return nil
}
