package concurrent

import (
	"errors"
	"reflect"
)

func Invoke(fn interface{}, args ...any) ([]any, error) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return nil, errors.New("not FUNC type")
	}
	if v.Type().NumIn() != len(args) {
		return nil, errors.New("wrong argument count")
	}
	agv := make([]reflect.Value, len(args))
	for i, a := range args {
		agv[i] = reflect.ValueOf(a)
	}
	var values []any
	for _, val := range v.Call(agv) {
		values = append(values, val.Interface())
	}
	return values, nil
}
