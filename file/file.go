package file

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func WriteToFile(texts []string) error {
	var f *os.File
	var err error

	if len(texts) == 0 {
		return nil
	}

	f, err = os.Create(fmt.Sprintf("%s.txt", time.Now().String()))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	for _, text := range texts {
		if _, err = f.WriteString(fmt.Sprintf("%s\n", text)); err != nil {
			return err
		}
	}

	log.Printf("Successfully written to %s", f.Name())
	return nil
}
