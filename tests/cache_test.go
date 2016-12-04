package tests

//go:generate gotemplate "../cache" "StringIntCache(string, int)"

import (
	"strconv"
	"sync"
	"testing"
)

func TestNewStringIntCache(t *testing.T) {
	if a := NewStringIntCache(); a.Count() != 0 {
		t.Fatal()
	}
}

func TestCacheSet(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntCache()
	go func() {
		for counter := 0; counter < total; counter++ {
			go func(counter int) {
				a.Set(strconv.Itoa(counter), counter, 10)
				wg.Done()
			}(counter)
		}
	}()
	wg.Wait()

	if a.Count() != total {
		t.Fail()
	}
}

func TestCacheGet(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntCache()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter, 10)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	if a.Count() != total {
		t.Fail()
	}
	if fifty, ok := a.Get("50"); !ok || fifty != 50 {
		t.Fail()
	}
}

func TestCacheCount(t *testing.T) {
	var wg sync.WaitGroup
	total := 10000
	wg.Add(total)
	a := NewStringIntCache()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter, 10)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	if a.Count() != total {
		t.Fail()
	}
}
