package crypto

import (
	"math/big"
	"math/rand"
	"time"
)

func gcd(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

func modInverse(a, m *big.Int) *big.Int {
	m0, x0, x1 := new(big.Int).Set(m), big.NewInt(0), big.NewInt(1)
	if m.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(0)
	}
	for a.Cmp(big.NewInt(1)) > 0 {
		q := new(big.Int).Div(m, a)
		m, a = a, new(big.Int).Mod(m, a)
		x0, x1 = x1.Sub(x1, new(big.Int).Mul(q, x0)), x0
	}
	if x1.Cmp(big.NewInt(0)) < 0 {
		x1 = x1.Add(x1, m0)
	}
	return x1
}

func GenerateRSAKeys(bits int) (*big.Int, *big.Int, *big.Int, *big.Int) {
	rand.Seed(time.Now().UnixNano())
	p := big.NewInt(0) // Generate prime p
	q := big.NewInt(0) // Generate prime q
	// Use a prime generation method here
	n := new(big.Int).Mul(p, q)
	phi := new(big.Int).Sub(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	e := big.NewInt(65537)
	d := modInverse(e, phi)
	return e, d, n, phi
}

func EncryptRSA(publicKey, n *big.Int, m *big.Int) *big.Int {
	return new(big.Int).Exp(m, publicKey, n)
}

func DecryptRSA(privateKey, n *big.Int, c *big.Int) *big.Int {
	return new(big.Int).Exp(c, privateKey, n)
}
