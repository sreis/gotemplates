package tests

//go:generate gotemplate concurrentmap "StringIntMap(string, int)"

import (
	"sync"
	"testing"
)

func TestNewStringIntMapConcurrentMap(t *testing.T) {
	a := NewStringIntMap()
	if a.Count() != 0 {
		t.Fatal()
	}
}

func TestConcurrentMapSet(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	go func() {
                for counter := 0; counter < total; counter++ {
                        go func() {
                                a.Set(string(counter), counter)
                                wg.Done()
                        }()
                }
        }()

	wg.Wait()

	println(a.Count())
	if a.Count() != total { t.Fail() }
}

func TestConcurrentMapGet(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()

	//if a.Count() != total { t.Fail() }
	fifty, ok := a.Get("50")
	if !ok || fifty != 50 { t.Fail() }
}

func TestConcurrentMapRemove(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()
	a.Remove("50")
	removed, ok := a.Get("50")
	if !ok || removed != 50 { t.Fail() }
}

func TestConcurrentMapGetAndRemove(t *testing.T) {
	var wg sync.WaitGroup
	total := 100
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
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
	var wg sync.WaitGroup
	total := 10000
	wg.Add(total)
	a := NewStringIntMap()
	for counter := 0; counter < total; counter++ {
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
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
		go func() {
			a.Set(string(counter), counter)
			wg.Done()
		}()
	}

	wg.Wait()

	if !a.Has("50") { t.Fail() }
}

