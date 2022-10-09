package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/husamettinarabaci/tinyurl/util"
)

var validUserType validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if userType, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedUserType(userType)
	}
	return false
}
