package logger

import "log"

func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}
