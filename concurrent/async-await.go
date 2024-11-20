package concurrent

import (
	"errors"
	"reflect"
)

type Future struct {
	Error error
	Data  []any
}

func Promise(task any, args ...any) chan Future {
	ch := make(chan Future)
	go func() {
		defer close(ch)
		result, err := Invoke(task, args...)
		ch <- Future{
			Data:  result,
			Error: err,
		}
	}()
	return ch
}

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
