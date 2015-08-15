package concurrentmap

import (
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func TestNewConcurrentMap(t *testing.T) {
	a := New()
	if a.Count() != 0 {
		t.Fatal()
	}
}

func TestConcurrentMapSet(t *testing.T) {
	total := 100
	a := New()
	for counter := 0; counter < total; counter++ {
		wg.Add(1)
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()

	if a.Count() != total { t.Fail() }
}

func TestConcurrentMapGet(t *testing.T) {
	total := 100
	a := New()
	for counter := 0; counter < total; counter++ {
		wg.Add(1)
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()

	if a.Count() != total { t.Fail() }
	fifty, _ := a.Get("50")
	if fifty != 50 { t.Fail() }
}

func TestConcurrentMapRemove(t *testing.T) {
	total := 100
	a := New()
	for counter := 0; counter < total; counter++ {
		wg.Add(1)
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()
	a.Remove("50")
	_, ok := a.Get("50")
	if ok { t.Fail() }
}

func TestConcurrentMapGetAndRemove(t *testing.T) {    total := 100
	a := New()
	for counter := 0; counter < total; counter++ {
		wg.Add(1)
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()
	val, ok := a.GetAndRemove("50")
	if val != 50 { t.Fail() }
	if !ok { t.Fail() }
}

func TestConcurrentMapCount(t *testing.T) {
	total := 10000
	a := New()
	for counter := 0; counter < total; counter++ {
		wg.Add(1)
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()

	if a.Count() != total { t.Fail() }
}

func TestConcurrentMapHas(t *testing.T) {
	total := 100
	a := New()
	for counter := 0; counter < total; counter++ {
		wg.Add(1)
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()

	if !a.Has("50") { t.Fail() }
}

