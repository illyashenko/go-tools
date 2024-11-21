package concurrent

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
