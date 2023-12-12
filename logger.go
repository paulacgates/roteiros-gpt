package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Configura o log da aplicação
func ConfiguraLog(destinoLog string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logrotate(destinoLog)
}

func logrotate(destinoLog string) {
	if destinoLog != "" {
		destinoLog += "/log/roteiros"
		dirName := destinoLog
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			if err := os.MkdirAll(dirName, 0666); err != nil {
				log.Fatalln("ERROR:", err)
			}
		}

		today := time.Now()
		fileName := fmt.Sprintf("%s-%s.log", dirName, today.Format("2006-01-02"))
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalln("ERROR:", err)
		}
		log.SetOutput(file)

		day := 24 * time.Hour
		week := 7 * day
		fileToRemove := fmt.Sprintf("%s-%s.log", dirName, today.Add(-week).Format("2006-01-02"))
		err = os.Remove(fileToRemove)
		if err != nil {
			log.Println("ATENÇÃO: Não foi possível remover", fileToRemove, ".", err)
		}
	}
}
