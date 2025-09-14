package rbcl

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoCoreRistretto255IsValidPoint(t *testing.T) {
	tests := []struct {
		name string
		p    func() []byte
		exp  bool
	}{
		{
			name: "TestCryptoCoreRistretto255IsValidPoint success",
			p: func() []byte {
				return CryptoCoreRistretto255Random()
			},
			exp: true,
		},
		{
			name: "TestCryptoCoreRistretto255IsValidPoint fail",
			p: func() []byte {
				return RandomBytes(8)
			},
			exp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, CryptoCoreRistretto255IsValidPoint(tt.p()))
		})
	}
}

func TestCryptoCoreRistretto255Random(t *testing.T) {
	tests := []struct {
		name string
		p    func() []byte
		exp  bool
	}{
		{
			name: "TestCryptoCoreRistretto255Random success",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			exp:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, CryptoCoreRistretto255IsValidPoint(tt.p()))
		})
	}
}

func TestCryptoCoreRistretto255FromHash(t *testing.T) {
	tests := []struct {
		name string
		h    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255FromHash success",
			h:    func() []byte { return bytes.Repeat([]byte{0x70}, CryptoCoreRistretto255HashBytes) },
			err:  nil,
		},
		{
			name: "TestCryptoCoreRistretto255FromHash fail",
			h:    func() []byte { return bytes.Repeat([]byte{0x70}, CryptoCoreRistretto255HashBytes-8) },
			err:  ErrBadHashLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := CryptoCoreRistretto255FromHash(tt.h())
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.True(t, CryptoCoreRistretto255IsValidPoint(p))
			}
		})
	}
}

func TestCryptoCoreRistretto255Add(t *testing.T) {
	tests := []struct {
		name string
		p    func() []byte
		q    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255Add success one",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			q:    func() []byte { return CryptoCoreRistretto255Random() },
			err:  nil,
		},
		{
			name: "TestCryptoCoreRistretto255Add success two",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			q: func() []byte {
				q, _ := CryptoCoreRistretto255FromHash(
					bytes.Repeat([]byte{0x70}, CryptoCoreRistretto255HashBytes),
				)
				return q
			},
		},
		{
			name: "TestCryptoCoreRistretto255Add fail one",
			p:    func() []byte { return RandomBytes(CryptoCoreRistretto255Bytes - 8) },
			q:    func() []byte { return CryptoCoreRistretto255Random() },
			err:  ErrBadPointLength,
		},
		{
			name: "TestCryptoCoreRistretto255Add fail two",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			q:    func() []byte { return RandomBytes(CryptoCoreRistretto255Bytes - 8) },
			err:  ErrBadPointLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p()
			q := tt.q()

			pq, err := CryptoCoreRistretto255Add(p, q)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.True(t, CryptoCoreRistretto255IsValidPoint(pq))
			}

			qp, err := CryptoCoreRistretto255Add(p, q)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.True(t, CryptoCoreRistretto255IsValidPoint(qp))
				assert.Equal(t, pq, qp) // test commutativity
			}
		})
	}
}

func TestCryptoCoreRistretto255Sub(t *testing.T) {
	tests := []struct {
		name string
		p    func() []byte
		q    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255Sub success one",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			q:    func() []byte { return CryptoCoreRistretto255Random() },
			err:  nil,
		},
		{
			name: "TestCryptoCoreRistretto255Sub success two",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			q: func() []byte {
				q, err := CryptoCoreRistretto255FromHash(
					bytes.Repeat([]byte{0x70}, CryptoCoreRistretto255HashBytes),
				)
				if err != nil {
					panic(err)
				}
				return q
			},
		},
		{
			name: "TestCryptoCoreRistretto255Sub fail one",
			p:    func() []byte { return RandomBytes(CryptoCoreRistretto255Bytes - 8) },
			q:    func() []byte { return CryptoCoreRistretto255Random() },
			err:  ErrBadPointLength,
		},
		{
			name: "TestCryptoCoreRistretto255Sub fail two",
			p:    func() []byte { return CryptoCoreRistretto255Random() },
			q:    func() []byte { return RandomBytes(CryptoCoreRistretto255Bytes - 8) },
			err:  ErrBadPointLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p()
			q := tt.q()

			masked, err := CryptoCoreRistretto255Add(p, q)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.True(t, CryptoCoreRistretto255IsValidPoint(masked))
			}

			unmasked, err := CryptoCoreRistretto255Sub(masked, q)
			assert.NoError(t, err)
			assert.Equal(t, p, unmasked)
		})
	}
}

func TestCryptoScalarMultRistretto255Base(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		err  error
	}{
		{
			name: "TestCryptoScalarMultRistretto255Base success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoScalarMultRistretto255Base fail",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes - 8)
			},
			err: ErrBadScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := CryptoScalarMultRistretto255Base(tt.s())
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.True(t, CryptoCoreRistretto255IsValidPoint(p))
			}
		})
	}
}

func TestCryptoScalarMultRistretto255(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		p    func() []byte
		err  error
	}{
		{
			name: "TestCryptoScalarMultRistretto255 success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			p: func() []byte {
				return CryptoCoreRistretto255Random()
			},
			err: nil,
		},
		{
			name: "TestCryptoScalarMultRistretto255 fail one",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes - 8)
			},
			p: func() []byte {
				return CryptoCoreRistretto255Random()
			},
			err: ErrBadScalarLength,
		},
		{
			name: "TestCryptoScalarMultRistretto255 fail two",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			p: func() []byte {
				return RandomBytes(CryptoCoreRistretto255Bytes - 8)
			},
			err: ErrBadPointLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			p := tt.p()

			masked, err := CryptoScalarMultRistretto255(s, p)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.True(t, CryptoCoreRistretto255IsValidPoint(masked))
			}

			sInv, err := CryptoCoreRistretto255ScalarInvert(s)
			assert.NoError(t, err)
			unmasked, err := CryptoScalarMultRistretto255(sInv, masked)
			assert.NoError(t, err)
			assert.Equal(t, unmasked, p)
		})
	}
}
