package cryptography

import (
	"errors"
	"testing"
)

func TestGcd_0(t *testing.T) {
	a, b := 5, 1
	want := 1
	gcd, err := gcd(a, b)

	if err != nil {
		t.Fatal(err)
	}

	if gcd != want {
		t.Fatalf("have gcd(%d, %d)=%d, but want %d", a, b, gcd, want)
	}
}

func TestGcd_1(t *testing.T) {
	a, b := 3, 6
	want := 3
	gcd, err := gcd(a, b)

	if err != nil {
		t.Fatal(err)
	}

	if gcd != want {
		t.Fatalf("have gcd(%d, %d)=%d, but want %d", a, b, gcd, want)
	}
}

func TestGcd_2(t *testing.T) {
	a, b := 0, 12345
	want := 12345
	gcd, err := gcd(a, b)

	if err != nil {
		t.Fatal(err)
	}

	if gcd != want {
		t.Fatalf("have gcd(%d, %d)=%d, but want %d", a, b, gcd, want)
	}
}

func TestGcd_3(t *testing.T) {
	_, err := gcd(0, 0)
	if err == nil {
		t.Fatal(errors.New("this should have thrown an error"))
	}
}

func TestEEA_0(t *testing.T) {
	gcd, ca, cb, err := eea(5, 11)

	if err != nil {
		t.Fatal(err)
	}

	if gcd != 1 || ca != -2 || cb != 1 {
		t.Fatalf("have (%d, %d, %d), but want (1, -2, 1)", gcd, ca, cb)
	}
}

func TestEEA_1(t *testing.T) {
	gcd, ca, cb, err := eea(482, 1180)

	if err != nil {
		t.Fatal(err)
	}

	if gcd != 2 || ca != 71 || cb != -29 {
		t.Fatalf("have (%d, %d, %d), but want (2, 71, -29)", gcd, ca, cb)
	}
}

func TestEEA_2(t *testing.T) {
	gcd, ca, cb, err := eea(11111, 12345)

	if err != nil {
		t.Fatal(err)
	}

	if gcd != 1 || ca != 2471 || cb != -2224 {
		t.Fatalf("have (%d, %d, %d), but want (1, 2471, -2224)", gcd, ca, cb)
	}
}

func TestEEA_3(t *testing.T) {
	_, _, _, err := eea(0, 0)
	if err == nil {
		t.Fatal(errors.New("this should have thrown an error"))
	}
}
