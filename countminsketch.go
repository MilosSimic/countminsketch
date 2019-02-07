package main

import (
	"github.com/spaolacci/murmur3"
	"hash"
	"math"
	"time"
)

type CountMinSketch struct {
	k      uint          // number of hash functions
	n      uint          // size of array
	sketch [][]uint      // just sketch
	h      []hash.Hash32 // hash functions
}

func New(k, n uint) *CountMinSketch {
	return &CountMinSketch{
		k:      k,
		n:      n,
		sketch: initSketch(k, n),
		h:      fns(k),
	}
}

func NewWithEstiments(epsilon, delta float64) *CountMinSketch {
	// Column = e/ε size of array
	// Row = ln (1/δ) number of functions
	//ε (Epsilon) is “how much error is added to our counts with each item we add to the cm sketch”.
	//δ (Delta) is “with what probability do we want to allow the count estimate to be outside of our epsilon error rate”

	k := uint(math.Ceil(math.Log(math.E / delta)))
	n := uint(math.Ceil(math.E / epsilon))
	return &CountMinSketch{
		k:      k,
		n:      n,
		sketch: initSketch(k, n),
		h:      fns(k),
	}
}

func fns(k uint) []hash.Hash32 {
	h := []hash.Hash32{}
	ts := uint(time.Now().Unix())
	for i := uint(0); i < k; i++ {
		h = append(h, murmur3.New32WithSeed(uint32(ts+1)))
	}
	return h
}

func initSketch(k, n uint) [][]uint {
	sketch := make([][]uint, k)
	for col := range sketch {
		sketch[col] = make([]uint, n)
	}
	return sketch
}

func prepare(hfn hash.Hash32, key string, size uint) uint32 {
	hfn.Write([]byte(key))
	idx := hfn.Sum32() % uint32(size)
	hfn.Reset()
	return idx
}

func (c *CountMinSketch) Add(key string) {
	for ridx, fn := range c.h {
		cidx := prepare(fn, key, c.n)
		c.sketch[ridx][cidx]++
	}
}

func min(temp []uint) uint {
	min := temp[0]
	for _, val := range temp {
		if val < min {
			min = val
		}
	}
	return min
}

func (c *CountMinSketch) Query(key string) uint {
	temp := []uint{}
	for ridx, fn := range c.h {
		cidx := prepare(fn, key, c.n)
		val := c.sketch[ridx][cidx]
		if val != uint(0) {
			temp = append(temp, val)
		}
	}
	return min(temp)
}
