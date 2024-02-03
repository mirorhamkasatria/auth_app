package utils

import (
	"log"

	"github.com/jinzhu/copier"
)

func Copier(to interface{}, from interface{}) error {
	if err := copier.Copy(to, from); err != nil {
		log.Print("Error Copier", err)
		return nil
	}
	return nil
}
