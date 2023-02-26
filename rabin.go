package gorabin

import (
	"crypto/rand"
	"io"
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

func exp(p *big.Int) (s, d *big.Int) {
	d = new(big.Int).Sub(p, one)
	s = big.NewInt(0)

	tmp := new(big.Int)
	for {
		s.Add(s, one)
		d.Rsh(d, 1)
		if tmp.And(d, one).Cmp(one) == 0 {
			return
		}
	}
}

func isPrime(random io.Reader, p *big.Int, count int) (bool, error) {
	p = new(big.Int).Set(p)
	if p.Cmp(one) == 0 {
		return false, nil
	}
	if p.Cmp(two) == 0 {
		return true, nil
	}
	if new(big.Int).And(p, one).Cmp(zero) == 0 {
		// p%2==0(p&1==0)
		return false, nil
	}
	// p is odd number
	s, d := exp(p)
	for i := 0; i < count; i++ {
		ok, err := isPrimeOnce(random, p, s, d)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func isPrimeOnce(random io.Reader, p, s, d *big.Int) (bool, error) {
	p = new(big.Int).Set(p)
	max := new(big.Int).Sub(p, one)

	a := new(big.Int).Set(zero)
	for a.Cmp(zero) == 0 {
		var err error
		a, err = rand.Int(random, p)
		if err != nil {
			return false, err
		}
	}
	{
		// not prime if a^d!=1(mod p)
		atmp := new(big.Int).Set(a)
		atmp = a.Exp(atmp, d, p)
		if atmp.Cmp(one) == 0 {
			return true, nil
		}
	}
	aa := new(big.Int)
	pw := new(big.Int)
	for r := big.NewInt(0); r.Cmp(s) < 0; r.Add(r, one) {
		aa.Set(a)
		pw.Exp(two, r, nil).Mul(pw, d)
		aa.Exp(aa, pw, p)
		if aa.Cmp(max) == 0 {
			return true, nil
		}
	}
	return false, nil
}
