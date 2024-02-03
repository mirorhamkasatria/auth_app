package utils

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var mtx sync.Mutex

func ValidateStruct(v *validator.Validate, data interface{}) error {
	mtx.Lock()
	defer mtx.Unlock()
	trans := getTranslator(v)
	getJsonTags(v)
	if err := v.Struct(data); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(e.Translate(trans))
		}
	}

	return nil
}

func getJsonTags(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func getTranslator(v *validator.Validate) ut.Translator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Print("Error get Translator", err)
	}

	return trans
}
