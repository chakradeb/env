package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Parse(v interface{}) error {
	r := reflect.ValueOf(v)
	if kind := r.Kind(); kind != reflect.Ptr {
		return fmt.Errorf("env: expected %s but got %s", reflect.Ptr, kind)
	}
	elem := r.Elem()
	rType := elem.Type()
	if kind := rType.Kind(); kind != reflect.Struct {
		return fmt.Errorf("env: expected %s but got %s", reflect.Struct, kind)
	}
	for i := 0; i < rType.NumField(); i++ {
		tag := rType.Field(i).Tag
		val, ok := os.LookupEnv(tag.Get("env"))
		if !ok {
			continue
		}
		err := setValue(elem.Field(i), val)
		if err != nil {
			return fmt.Errorf("env: %s", err.Error())
		}
	}
	return nil
}

func setValue(field reflect.Value, value string) error {
	switch kind := field.Kind(); kind {
		case reflect.String:
			field.SetString(value)
		case reflect.Int:
			val, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				return err
			}
			field.SetInt(val)
		default:
			return fmt.Errorf("%s is not a supported type", kind)
	}
	return nil
}
