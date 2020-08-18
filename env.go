package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Parse(v interface{}) []error {
	var errs []error

	r := reflect.ValueOf(v)
	if kind := r.Kind(); kind != reflect.Ptr {
		return []error{fmt.Errorf("env: expected %s but got %s", reflect.Ptr, kind)}
	}
	elem := r.Elem()
	rType := elem.Type()
	if kind := rType.Kind(); kind != reflect.Struct {
		return []error{fmt.Errorf("env: expected %s but got %s", reflect.Struct, kind)}
	}
	for i := 0; i < rType.NumField(); i++ {
		tag := rType.Field(i).Tag
		val, ok := os.LookupEnv(tag.Get("env"))
		if !ok {
			continue
		}
		err := setValue(elem.Field(i), val)
		if err != nil {
			errs = append(errs, fmt.Errorf("env: %s", err.Error()))
		}
	}
	return errs
}

func setValue(field reflect.Value, value string) error {
	switch kind := field.Kind(); kind {
		case reflect.String:
			field.SetString(value)
		case reflect.Int:
			return setInt(field, value, 0)
		case reflect.Int8:
			return setInt(field, value, 8)
		case reflect.Int16:
			return setInt(field, value, 16)
		case reflect.Int32:
			return setInt(field, value, 32)
		case reflect.Int64:
			return setInt(field, value, 64)
		case reflect.Float32:
			return setFloat(field, value, 32)
		case reflect.Float64:
			return setFloat(field, value, 64)
		default:
			return fmt.Errorf("%s is not a supported type", kind)
	}
	return nil
}

func setInt(field reflect.Value, value string, bitSize int) error {
	val, err := strconv.ParseInt(value, 10, bitSize)
	if err != nil {
		return err
	}
	field.SetInt(val)
	return nil
}

func setFloat(field reflect.Value, value string, bitSize int) error {
	val, err := strconv.ParseFloat(value, bitSize)
	if err != nil {
		return err
	}
	field.SetFloat(val)
	return nil
}
