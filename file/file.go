package file

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func WroteToFile(texts []string) {
	var f *os.File
	var err error

	if len(texts) == 0 {
		goto exit
	}

	f, err = os.Create(fmt.Sprintf("%s.txt", time.Now().String()))
	if err != nil {
		log.Error(err)
		goto exit
	}
	defer func() {
		_ = f.Close()
	}()

	for _, text := range texts {
		if _, err = f.WriteString(fmt.Sprintf("%s\n", text)); err != nil {
			log.Error(err)
			goto exit
		}
	}

	log.Printf("Successfully wrote to %s", f.Name())

exit:
}
