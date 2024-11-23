package middleware

import (
	"fmt"
	"reflect"
)

func validarCamposNoNulos(v interface{}) error {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr && val.IsNil() {
		return fmt.Errorf("json mal formado")
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		campo := val.Field(i)
		if campo.Kind() == reflect.Ptr && campo.IsNil() {
			return fmt.Errorf("el campo '%s' no esta definido", val.Type().Field(i).Name)
		}
	}

	return nil
}
