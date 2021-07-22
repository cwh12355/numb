package ratutils

import (
	"log"
	"math/big"
)

var ZEROINT = big.NewInt(0)
var ONE = big.NewRat(1, 1)
var TEN = big.NewRat(10, 1)

func ExpInt(x *big.Rat, n int) *big.Rat {
	if n == 0 {
		return ONE
	}
	r := new(big.Rat).Set(x)
	for i := 1; i < n; i++ {
		r.Mul(r, x)
	}
	if n < 0 {
		r.Quo(ONE, r)
	}
	return r
}

func Exp(b, n int64) *big.Rat {
	if n == 0 {
		return ONE
	}
	int := big.NewInt(b)
	rat := big.NewRat(1, 1)
	if n > 0 {
		rat.Num().Set(int.Exp(int, big.NewInt(n), nil))
	} else {
		rat.Denom().Set(int.Exp(int, big.NewInt(-n), nil))
	}
	return rat
}

func MulInt(a *big.Rat, n int64) *big.Rat {
	x := new(big.Rat).Set(a)
	x.Mul(x, big.NewRat(n, 1))
	return x
}

func DivRat(a, x *big.Rat) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, x)
	return num
}

func DivInt(a *big.Rat, x int64) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, big.NewRat(x, 1))
	return num
}

func ModRat(a, n *big.Rat) *big.Rat {
	// r = a - n * trunc(a / n)
	return new(big.Rat).Sub(a, new(big.Rat).Mul(n, Trunc(new(big.Rat).Quo(a, n))))
}

func Trunc(x *big.Rat) *big.Rat {
	i := new(big.Int).Div(x.Num(), x.Denom())
	r := new(big.Rat)
	r.Num().Set(i)
	return r
}

func ToInt(x *big.Rat) *big.Int {
	if IsInt(x) {
		return x.Num()
	}
	// TODO: quorem vs. divmod
	quo := new(big.Int)
	rem := new(big.Int)
	quo.QuoRem(x.Num(), x.Denom(), rem)
	return quo
}

func IsInt(x *big.Rat) bool {
	return x.Denom().Cmp(ONE.Num()) == 0
}

func Num(x string) *big.Rat {
	rat, ok := new(big.Rat).SetString(x)
	if !ok {
		log.Fatal("unable to parse: " + x)
	}
	return rat
}
