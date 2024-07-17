package cryptography

import (
	"errors"
)

// returns greatest common divisor of two integers
func gcd(a, b int) (gcd int, err error) {
	gcd, _, _, err = eea(a, b)
	return gcd, err
}

// performs the Extended Euclidean Algorithm on the integers a and b
// returns gcd, s, and t that satisfy a*s+b*t=gcd(a,b).
func eea(a, b int) (gcd, s, t int, err error) {

	if a == 0 && b == 0 {
		return 0, 0, 0, errors.New("can't find gcd with two 0 terms")
	}

	// ensure a >= b
	if b > a {
		a, b = b, a
	}

	if b == 0 {
		return a, 1, 0, nil // t can, in fact, be any value since b is zero
	}

	// core algorithm
	s_1, s_2, t_1, t_2 := 1, 0, 0, 1
	for {
		r := a % b
		if r == 0 {
			break
		}
		q := a / b
		s_1, s_2 = -q*s_1+s_2, s_1
		t_1, t_2 = -q*t_1+t_2, t_1
		a, b = b, r
	}

	return b, s_1, t_1, nil
}
