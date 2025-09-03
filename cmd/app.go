package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/JerryJeager/skelgo/cmd/config"
	"github.com/JerryJeager/skelgo/cmd/docs"
	"github.com/JerryJeager/skelgo/cmd/git"
	"github.com/JerryJeager/skelgo/cmd/internal"
	"github.com/JerryJeager/skelgo/cmd/internal/db"
	"github.com/JerryJeager/skelgo/cmd/internal/env"
	"github.com/JerryJeager/skelgo/cmd/internal/http"
	"github.com/JerryJeager/skelgo/cmd/internal/service"
	"github.com/JerryJeager/skelgo/cmd/internal/utils"
	"github.com/JerryJeager/skelgo/cmd/manualwire"
	"github.com/JerryJeager/skelgo/cmd/middleware"
)

func InitProject(projectName, modulePath string) error {
	projectPath := filepath.Join(projectName)
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	//create go mod
	cmd := exec.Command("go", "mod", "init", modulePath)
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	// handle config folder
	if err := config.InitConfig(projectName, modulePath); err != nil {
		return err
	}

	//handle docs folder
	if err := docs.InitDocs(projectName, modulePath); err != nil {
		return err
	}

	//handle models
	if err := internal.CreateModels(projectName); err != nil {
		return err
	}

	//handle utils
	if err := utils.HandleUtils(projectName); err != nil {
		return err
	}

	//handle migration file
	if err := db.CreateMigrationFile(projectName); err != nil {
		return err
	}

	//handle envs
	if err := env.HandleEnvs(projectName); err != nil {
		return err
	}

	//handle service
	if err := service.HanldeService(projectName, modulePath); err != nil {
		return err
	}

	//handle http
	if err := http.HandleHttp(projectName, modulePath); err != nil {
		return err
	}

	//handle manualwire
	if err := manualwire.HandleManualWire(projectName, modulePath); err != nil {
		return err
	}

	//handle middleware
	if err := middleware.HandleMiddleware(projectName, modulePath); err != nil {
		return err
	}

	//handle app
	if err := HandleApp(projectName, modulePath); err != nil {
		return err
	}

	//tidy go modules
	tidy := exec.Command("go", "mod", "tidy")
	tidy.Dir = projectPath
	tidy.Stdout = os.Stdout
	tidy.Stderr = os.Stderr
	err = tidy.Run()
	if err != nil {
		return err
	}

	//initialize git repository
	if err := git.GitInitProject(projectName); err != nil {
		return err
	}

	return nil
}

func HandleApp(projectName, modulePath string) error {
	cmdPath := filepath.Join(projectName, "cmd")
	err := os.MkdirAll(cmdPath, os.ModePerm)
	if err != nil {
		return err
	}

	appPath := filepath.Join(cmdPath, "app.go")
	appFileDest, err := os.Create(appPath)
	if err != nil {
		return err
	}
	defer appFileDest.Close()

	_, err = appFileDest.WriteString(GenerateAppFile(modulePath))
	if err != nil {
		return err
	}
	return nil
}

func GenerateAppFile(modulePath string) string {
	return fmt.Sprintf(`
package cmd

import (
	"log"
	"os"

	"%s/manualwire"
	"%s/middleware"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome",
		})
	})

	userController := manualwire.GetUserController()

	api := router.Group("/api/v1")
	users := api.Group("/users")

	users.POST("/signup", userController.CreateUser)
	users.POST("/verify-email", userController.VerifyUserEmail)
	users.POST("/login", userController.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panic("failed to run server")
	}
}

	`, modulePath, modulePath)
}
