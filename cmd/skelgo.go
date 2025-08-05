package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func Init(projectName, modulePath string) error {
	s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Initializing %s=> ", projectName)
	s.Start()

	if err := InitProject(projectName, modulePath); err != nil {
		return err
	}

	s.Stop()
	return nil
}
