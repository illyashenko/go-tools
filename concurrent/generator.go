package concurrent

import "context"

func Generator[T any](doneCh chan struct{}, data []T) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, el := range data {
			select {
			case <-doneCh:
				return
			case ch <- el:
			}
		}
	}()
	return ch
}

func GeneratorWithContext[T any](ctx context.Context, data []T) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, el := range data {
			select {
			case <-ctx.Done():
				return
			case ch <- el:
			}
		}
	}()
	return ch
}

func Process[T any](in chan T, functor func(param T)) {
	for i := range in {
		functor(i)
	}
}

func ProcessWithContext[T any](ctx context.Context, in chan T, functor func(param T)) {
	for i := range in {
		select {
		case <-ctx.Done():
			return
		default:
			functor(i)
		}
	}
}
