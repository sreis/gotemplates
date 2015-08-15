package tests

//go:generate gotemplate "../concurrentmap" "StringIntMap(string, int)"

import (
	"strconv"
	"sync"
	"testing"
)

func TestNewStringIntMapConcurrentMap(t *testing.T) {
	if a := NewStringIntMap(); a.Count() != 0 { t.Fatal() }
}

func TestConcurrentMapSet(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	go func() {
                for counter := 0; counter < total; counter++ {
                        go func(counter int) {
                                a.Set(strconv.Itoa(counter), counter)
                                wg.Done()
                        }(counter)
                }
        }()
	wg.Wait()

	if a.Count() != total { t.Fail() }
}

func TestConcurrentMapGet(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	if a.Count() != total { t.Fail() }
	if fifty, ok := a.Get("50"); !ok || fifty != 50 { t.Fail() }
}

func TestConcurrentMapRemove(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	a.Remove("50")
	if _, ok := a.Get("50"); ok { t.Fail() }
}

func TestConcurrentMapGetAndRemove(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	if fifty, ok := a.GetAndRemove("50"); fifty != 50 || !ok { t.Fail() }
}

func TestConcurrentMapCount(t *testing.T) {
	var wg sync.WaitGroup
	total := 10000
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	if a.Count() != total { t.Fail() }
}

func TestConcurrentMapHas(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func(counter int) {
			a.Set(strconv.Itoa(counter), counter)
			wg.Done()
		}(counter)
	}
	wg.Wait()

	if !a.Has("50") { t.Fail() }
}

