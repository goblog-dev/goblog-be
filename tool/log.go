package tool

import "log"

func PrintLog(title string, err error) error {
	log.Println(title, "::", err.Error())
	return err
}
