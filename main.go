package main

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
)

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var TWO = big.NewInt(2)
var FOUR = big.NewInt(4)

var CPU_COUNT = runtime.NumCPU()

type PrimeGenerator struct {
	state uint
}

func IsPrime(p uint) bool {
	for i := uint(2); i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func (pg *PrimeGenerator) Next() uint {
	var ret uint

	for !IsPrime(pg.state) {
		pg.state += 2
	}
	ret = pg.state
	pg.state += 2
	return ret
}

func NewPG() PrimeGenerator {
	return PrimeGenerator{3}
}

func LLT(p uint, leastSigMask, s, remainingBits *big.Int) bool {
	// leastSigMask = M = (1<<p) - 1 = 2**p - 1
	leastSigMask.Set(ONE)
	leastSigMask.Lsh(leastSigMask, p)
	leastSigMask.Sub(leastSigMask, ONE)

	s.Set(FOUR)
	rptCnt := p - 2

	for i := uint(0); i < rptCnt; i++ {
		// s = s * s - 2
		s.Exp(s, TWO, nil)
		s.Sub(s, TWO)

		// remainingBits = s>>p
		remainingBits.Set(s)
		remainingBits.Rsh(remainingBits, p)

		// while remainingBits > 0
		for remainingBits.Cmp(ZERO) == 1 {
			// s = remainingBits + (s & leastSigMask)
			s.Set(remainingBits.Add(remainingBits, s.And(s, leastSigMask)))

			// remainingBits = s>>p
			remainingBits.Rsh(remainingBits, p)
		}
	}

	return s.Cmp(leastSigMask) == 0
}

func workerLLT(input, output chan uint) {
	// Make each big.Int live for the lifetime of the worker
	leastSigMask := new(big.Int)
	s := new(big.Int)
	remainingBits := new(big.Int)

	for n := range input {
		if LLT(n, leastSigMask, s, remainingBits) {
			output <- n
		}
	}
}

func StoI(s string, r int) int {
	var ret int
	var b byte
	ax := r
	i := len(s)

	for i > 0 {
		i--
		b = s[i] ^ 0x30
		if b < 10 {
			break
		}
	}

	if i == 0 {
		return 0
	}

	ret = int(b)

	for i > 0 {
		i--
		b = s[i] ^ 0x30
		if b > 10 {
			if i == 0 && b == 29 {
				ret *= -1
				return ret
			}
			continue
		}
		ret += int(b) * ax
		ax *= r
	}

	return ret
}

func main() {
	var target int
	input := make(chan uint, CPU_COUNT*2)
	output := make(chan uint, CPU_COUNT)
	count := 0

	if len(os.Args) > 1 {
		target = StoI(os.Args[1], 10)
	} else {
		target = 100
	}

	for i := 0; i < CPU_COUNT; i++ {
		go workerLLT(input, output)
	}

	go func() {
		pg := NewPG()
		for count < target {
			input <- pg.Next()
		}
	}()

	for count < target {
		fmt.Println(<-output)
		count++
	}

	close(input)
	close(output)
}
