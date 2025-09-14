package rbcl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoCoreRistretto255ScalarRandom(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarRandom success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, CryptoCoreRistretto255ScalarBytes, len(tt.s()))
		})
	}
}

func TestCryptoCoreScalarRistretto255ScalarReduce(t *testing.T) {
	tests := []struct {
		name string
		x    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreScalarRistretto255ScalarReduce success one",
			x: func() []byte {
				return RandomBytes(CryptoCoreRistretto255NonReducedScalarBytes)
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreScalarRistretto255ScalarReduce success two",
			x: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes)
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreScalarRistretto255ScalarReduce fail",
			x: func() []byte {
				return RandomBytes(CryptoCoreRistretto255NonReducedScalarBytes + 8)
			},
			err: ErrBadNonReducedScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := CryptoCoreScalarRistretto255ScalarReduce(tt.x())
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.NoError(t, err)
			}

			p := CryptoCoreRistretto255Random()
			masked, err := CryptoScalarMultRistretto255(s, p)
			assert.NoError(t, err)
			sInv, err := CryptoCoreRistretto255ScalarInvert(s)
			assert.NoError(t, err)
			unmasked, err := CryptoScalarMultRistretto255(sInv, masked)
			assert.NoError(t, err)
			assert.Equal(t, p, unmasked)
		})
	}
}

func TestCryptoCoreRistretto255ScalarNegate(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarNegate success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarNegate fail",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			err: ErrBadScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			sN, err := CryptoCoreRistretto255ScalarNegate(s)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.NoError(t, err)
			}

			z, err := CryptoCoreRistretto255ScalarAdd(s, sN)
			assert.NoError(t, err)
			assert.True(t, isZero(z))
		})
	}
}

func TestCryptoCoreRistretto255ScalarComplement(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarComplement success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarComplement fail one",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			err: ErrBadScalarLength,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarComplement fail two",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes - 8)
			},
			err: ErrBadScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			r, err := CryptoCoreRistretto255ScalarComplement(s)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.NoError(t, err)
			}

			// get "1" scalar by adding s and its complement
			one, err := CryptoCoreRistretto255ScalarAdd(s, r)
			assert.NoError(t, err)

			// generate a random point p, multiply it by "1", and test the output's equality with p
			p := CryptoCoreRistretto255Random()
			pOne, err := CryptoScalarMultRistretto255(one, p)
			assert.NoError(t, err)
			assert.Equal(t, p, pOne)
		})
	}
}

func TestCryptoCoreRistretto255ScalarInvert(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarInvert success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarInvert fail one",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			err: ErrBadScalarLength,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarInvert fail two",
			s: func() []byte {
				return make([]byte, CryptoCoreRistretto255ScalarBytes)
			},
			err: ErrZeroScalar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			sInv, err := CryptoCoreRistretto255ScalarInvert(s)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.NoError(t, err)
			}

			p := CryptoCoreRistretto255Random()
			masked, err := CryptoScalarMultRistretto255(s, p)
			assert.NoError(t, err)

			unmasked, err := CryptoScalarMultRistretto255(sInv, masked)
			assert.NoError(t, err)
			assert.Equal(t, p, unmasked)
		})
	}
}

func TestCryptoCoreRistretto255ScalarAdd(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		r    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarAdd success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			r: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarAdd fail one",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			r: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: ErrBadScalarLength,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarAdd fail two",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			r: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			err: ErrBadScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			r := tt.r()

			sr, err := CryptoCoreRistretto255ScalarAdd(s, r)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}

			rs, err := CryptoCoreRistretto255ScalarAdd(r, s)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, sr, rs)
			}
		})
	}
}

func TestCryptoCoreRistretto255ScalarSub(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		r    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarSub success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			r: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarSub fail one",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			r: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: ErrBadScalarLength,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarSub fail two",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			r: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			err: ErrBadScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			r := tt.r()

			sr, err := CryptoCoreRistretto255ScalarAdd(s, r)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			} else {
				assert.NoError(t, err)
			}

			ss, err := CryptoCoreRistretto255ScalarSub(sr, r)
			assert.NoError(t, err)
			assert.Equal(t, s, ss)
		})
	}
}

func TestCryptoCoreRistretto255ScalarMul(t *testing.T) {
	tests := []struct {
		name string
		s    func() []byte
		r    func() []byte
		err  error
	}{
		{
			name: "TestCryptoCoreRistretto255ScalarMul success",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			r: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: nil,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarMul fail one",
			s: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			r: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			err: ErrBadScalarLength,
		},
		{
			name: "TestCryptoCoreRistretto255ScalarMul fail two",
			s: func() []byte {
				return CryptoCoreRistretto255ScalarRandom()
			},
			r: func() []byte {
				return RandomBytes(CryptoCoreRistretto255ScalarBytes + 8)
			},
			err: ErrBadScalarLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s()
			r := tt.r()

			sr, err := CryptoCoreRistretto255ScalarMul(s, r)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}

			rs, err := CryptoCoreRistretto255ScalarMul(r, s)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, sr, rs)
			}
		})
	}
}
