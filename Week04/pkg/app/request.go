package app

import (
	"service-notification/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErros(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error(err.Key,err.Message)

	}
}
