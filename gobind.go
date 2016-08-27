package gobind

import (
	"errors"
	"fmt"
	"reflect"

	"strconv"
)

func getArg(index string, args ...interface{}) (interface{}, error) {
	i, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		return nil, err
	}
	if int64(len(args)) <= i {
		return nil, fmt.Errorf("invalid args number, want tag %d, but args number is only %d", i, len(args))
	}
	return args[i], nil
}

func Bind(target interface{}, args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("invalid args")
	}
	v := reflect.ValueOf(target).Elem()

	// if !v.CanSet() {
	// return errors.New("target should be passed with a pointer")
	// }

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("tag")
		obj, _ := getArg(tag, args...)
		field := v.Field(i)

		objValue := reflect.ValueOf(obj)
		val := objValue.FieldByName(t.Field(i).Name)
		copyValue(field, val)
	}

	return nil
}

func copyValue(dest reflect.Value, src reflect.Value) error {
	if dest.Kind() != dest.Kind() {
		return errors.New("wrong reflect Kind")
	}

	if !dest.CanSet() {
		return errors.New("target can't set...")
	}

	switch dest.Kind() {
	case reflect.String:
		dest.SetString(src.String())
	default:
		return errors.New("unsupport reflect kind..")
	}

	return nil
}
