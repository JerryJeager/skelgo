package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func Init(projectName, modulePath string) error {
	s := spinner.New(spinner.CharSets[40], 100*time.Millisecond) 
	s.Suffix = fmt.Sprintf(" Initializing project: %s", projectName)
	s.Start() 

	if err := InitProject(projectName, modulePath); err != nil {
		return err
	}
	time.Sleep(5 * time.Second) 

	s.Stop() 
	return nil
}

