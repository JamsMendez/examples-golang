package main

import (
	"sync"
	"testing"
)

func TestTaskWithOneContext(t *testing.T) {
	wg := new(sync.WaitGroup)
	TaskOneContext(wg)

	wg.Wait()
}

func TestTaskWithTwoContext(t *testing.T) {
	wg := new(sync.WaitGroup)
	TaskTwoContext(wg)

	wg.Wait()
}

func TestTaskWithCustomContext(t *testing.T) {
	wg := new(sync.WaitGroup)
	TaskWithCustonContext(wg)

	wg.Wait()
}
