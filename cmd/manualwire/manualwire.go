package manualwire

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateWireFile(modulePath string) string {
	return fmt.Sprintf(`
package manualwire

import (
	"%s/config"
	"%s/internal/http"
	"%s/internal/service/users"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}

	`, modulePath, modulePath, modulePath)
}

func HandleManualWire(projectName, modulePath string) error {
	manualwirePath := filepath.Join(projectName, "manualwire")
	err := os.MkdirAll(manualwirePath, os.ModePerm)
	if err != nil {
		return err
	}

	wirePath := filepath.Join(manualwirePath, "wire.go")
	wireFileDest, err := os.Create(wirePath)
	if err != nil {
		return err
	}
	defer wireFileDest.Close()

	_, err = wireFileDest.WriteString(GenerateWireFile(modulePath))
	if err != nil {
		return err
	}

	return nil
}
