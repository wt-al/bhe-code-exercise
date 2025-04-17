package sieve

import (
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

func NewSieve() Sieve {
	return &SieveImpl{}
}

type SieveImpl struct {
	
}

func  (sv *SieveImpl) NthPrime(n int64) int64 {
	upper := getUpLimit(n)
	primeList := make([]bool,  upper+1)
	for i:= int64(2); i <= upper; i++ {
		primeList[i] = true
	}
	for i := int64(2); i*i <= upper; i++ {
		if primeList[i] {
			for j := i * i; j <= upper; j += i {
				primeList[j] = false
			}
		}
	}
	count := int64(0)
	for i := int64(2); i <= upper; i++ {
		if primeList[i] {
			
			if count == n {
				return i
			}
			count++
		}
	}
	return -1
}


func getUpLimit(n int64) int64 {
	if n < 6 {
		return 15
	}
	nf := float64(n)
	logN := math.Log(nf)
	logLogN := math.Log(logN) 
	est := nf*logN + nf*logLogN
	return int64(est+0.5) 
}