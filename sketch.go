package main

type Sketch interface {
	Add(key string)
	Query(key string) uint
}
