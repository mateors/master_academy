package main

import (
	"container/list"
	"context"
	"fmt"
	"time"
)

type Future[T comparable] interface {
	Get() Result[T]
	Cancel()
}

type Result[S any] interface {
	Success() S
	Failure() error
}

type result[S any] struct {
	success S
	failure error
}

func (this *result[S]) Success() S {
	return this.success
}

func (this *result[S]) Failure() error {
	return this.failure
}

func (this *result[S]) String() string {

	if this.failure != nil {
		return fmt.Sprintf("%v", this.failure)
	} else {
		return fmt.Sprintf("%v", this.success)
	}
}

func main() {

	list.New()
	//f func() (T, error)
	f1 := NewFuture(func() (string, error) {
		time.Sleep(1000)
		return "F1", nil
	})

	//waiting for the result
	fmt.Printf("ready with %v \n", f1.Get())

	//Discarding the Future in case it's not needed:
	f3 := NewFuture(func() (int, error) {
		time.Sleep(100)
		fmt.Println("I'm done!")
		return 100, nil
	})
	//f3.Cancel()
	fmt.Println(f3.Get())

}

type future[T any] struct {
	result    *result[T]
	completed bool
	wait      chan bool
	ctx       context.Context
	cancel    func()
}

func (this *future[T]) Get() Result[T] {
	if this.completed {
		return this.result
	} else {
		fmt.Println("Need to wait...")
		select {
		case <-this.wait:
			return this.result
		case <-this.ctx.Done():
			return &result[T]{failure: this.ctx.Err()}
		}
	}
}

func (this *future[T]) Cancel() {
	this.cancel()
}

func NewFuture[T comparable](f func() (T, error)) Future[T] {

	fut := &future[T]{
		wait: make(chan bool),
	}
	fut.ctx, fut.cancel = context.WithCancel(context.Background())
	//...
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("future recovered: ", r)
			}
		}()
		success, failure := f()

		fut.result = &result[T]{success, failure}
		fut.completed = true
		fut.wait <- true
		close(fut.wait)
	}()

	return fut
}
