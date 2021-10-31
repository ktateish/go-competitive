package main

// vim:set ft=go:

// snip ------------------------------------------------------------------------

func IsPrime(n int) bool {
	assert(0 <= n && n < 4759123141)
	return intl_IsPrime(int32(n))
}

type Primes struct {
	p []bool
}

func NewPrimes(n int) Primes {
	m := n + 1
	p := Primes{
		p: make([]bool, m),
	}
	for i := 3; i < m; i += 2 {
		p.p[i] = true
	}
	p.p[2] = true
	i := 3
	for ; i*i < m; i += 2 {
		if !p.p[i] {
			continue
		}
		for j := i * i; j < m; j += i {
			p.p[j] = false
		}
	}
	return p
}

func (ps Primes) IsPrime(x int) bool {
	return ps.p[x]
}

type PrimesX struct {
	p []bool
	d []int
}

func NewPrimesX(n int) PrimesX {
	m := n + 1
	p := PrimesX{
		p: make([]bool, m),
		d: make([]int, m),
	}
	p.d[1] = 1
	for i := 3; i < m; i += 2 {
		p.p[i] = true
		p.d[i-1] = 2
	}
	if n%2 == 0 {
		p.d[n] = 2
	}
	p.p[2] = true
	i := 3
	for ; i*i < m; i += 2 {
		if !p.p[i] {
			continue
		}
		p.d[i] = i
		for j := i * i; j < m; j += i {
			p.p[j] = false
			if p.d[j] == 0 {
				p.d[j] = i
			}
		}
	}

	for i = i - 2; i < m; i += 2 {
		if p.d[i] == 0 {
			p.d[i] = i
		}
	}

	return p
}

func (ps PrimesX) IsPrime(x int) bool {
	return ps.p[x]
}

func (ps PrimesX) Divisor(x int) int {
	return ps.d[x]
}

func (ps PrimesX) Factorize(x int) [][2]int {
	res := make([][2]int, 0)
	if x == 1 {
		return res
	}
	v := x
	for v != 1 {
		d := ps.d[v]
		if d == x {
			break
		}
		if 0 < len(res) && res[len(res)-1][0] == d {
			res[len(res)-1][1]++
		} else {
			res = append(res, [2]int{d, 1})
		}
		v /= d
	}
	return res
}

func Factorize(n int) [][2]int {
	var res [][2]int
	for n%2 == 0 {
		if 0 < len(res) && res[len(res)-1][0] == 2 {
			res[len(res)-1][1]++
		} else {
			res = append(res, [2]int{2, 1})
		}
		n /= 2
	}

	for d := 3; d*d <= n; d += 2 {
		for n%d == 0 {
			if 0 < len(res) && res[len(res)-1][0] == d {
				res[len(res)-1][1]++
			} else {
				res = append(res, [2]int{d, 1})
			}
			n /= d
		}
	}
	if n != 1 {
		res = append(res, [2]int{n, 1})
	}

	return res
}

// NewFactorial(n, p int) returns 2 functions.
// 1. combination(x, r int) int: returns xCr modulo p (0 <= r <= x <= n)
// 2. permutation(x, r int) int: returns xPr modulo p (0 <= r <= x <= n)
// 3. factorial(x int) int: returns factorial x modulo p (0 <= x <= n)
func NewFactorial(n, p int) (combination func(int, int) int, permutation func(int, int) int, factorial func(int) int) {
	if n < 1 {
		panic("NewFactorial: n must be more than 0")
	}

	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = mul(fact[i-1], i, p)
	}

	ifact := make([]int, n+1)
	ifact[n] = inv(fact[n], p)
	for i := n - 1; 0 <= i; i-- {
		ifact[i] = mul(ifact[i+1], i+1, p)
	}

	combination = func(n, r int) int {
		return mul3(fact[n], ifact[n-r], ifact[r], p)
	}
	permutation = func(n, r int) int {
		return mul(fact[n], ifact[n-r], p)
	}
	factorial = func(i int) int {
		return fact[i]
	}
	return
}

func mul[T Integer](a, b, p T) T {
	return T(int64(a) * int64(b) % int64(p))
}

func mul3[T Integer](a, b, c, p T) T {
	return T(int64(a) * int64(b) % int64(p) * int64(c) % int64(p))
}

func div[T Integer](n, d, p T) T {
	return mul(n, inv(d, p), p)
}

func inv[T Integer](a, p T) T {
	return exp(a, p-2, p)
}

func exp[T Integer](n, e, p T) T {
	if e == 0 {
		return 1
	}
	if e&1 == 1 {
		return mul(n, exp(n, e-1, p), p)
	} else {
		return square(exp(n, e/2, p), p)
	}
}

func square[T Integer](a, p T) T {
	return mul(a, a, p)
}
