package rbcl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomBytes(t *testing.T) {
	tests := []struct {
		name string
		len  int
	}{
		{
			name: "TestRandomBytesZero",
			len:  0,
		},
		{
			name: "TestRandomBytesEight",
			len:  8,
		},
		{
			name: "TestRandomBytesSixteen",
			len:  16,
		},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%sLength", tt.name),
			func(t *testing.T) {
				b := RandomBytes(tt.len)
				assert.Equal(t, tt.len, len(b))
			},
		)
		t.Run(
			fmt.Sprintf("%sNotAllZero", tt.name),
			func(t *testing.T) {
				if tt.len == 0 {
					return
				}
				b := RandomBytes(tt.len)
				allZero := true
				for _, v := range b {
					if v != 0 {
						allZero = false
						break
					}
				}
				assert.False(t, allZero)
			},
		)
		t.Run(
			fmt.Sprintf("%sNotEqual", tt.name),
			func(t *testing.T) {
				if tt.len == 0 {
					return
				}
				b := RandomBytes(tt.len)
				c := RandomBytes(tt.len)
				assert.NotEqual(t, b, c)
			},
		)
	}
}

func TestRandomBytesDeterministic(t *testing.T) {
	tests := []struct {
		name string
		len  int
		seed func() []byte
		err  error
	}{
		{
			name: "TestRandomBytesDeterministicZero",
			len:  0,
			seed: func() []byte {
				return RandomBytes(RandomBytesSeedBytes)
			},
			err: nil,
		},
		{
			name: "TestRandomBytesDeterministicEight",
			len:  8,
			seed: func() []byte {
				return RandomBytes(RandomBytesSeedBytes)
			},
			err: nil,
		},
		{
			name: "TestRandomBytesDeterministicSixteen",
			len:  16,
			seed: func() []byte {
				return RandomBytes(RandomBytesSeedBytes)
			},
			err: nil,
		},
		{
			name: "TestRandomBytesDeterministicErr",
			len:  16,
			seed: func() []byte {
				return []byte("0123")
			},
			err: ErrBadSeedLength,
		},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%sLength", tt.name),
			func(t *testing.T) {
				b, err := RandomBytesDeterministic(tt.len, tt.seed())

				if tt.err != nil {
					assert.EqualError(t, err, tt.err.Error())
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.len, len(b))
				}
			},
		)
		t.Run(
			fmt.Sprintf("%sNotAllZero", tt.name),
			func(t *testing.T) {
				if tt.len == 0 {
					return
				}
				b, err := RandomBytesDeterministic(tt.len, tt.seed())

				if tt.err != nil {
					assert.EqualError(t, err, tt.err.Error())
				} else {
					assert.NoError(t, err)
					allZero := true
					for _, v := range b {
						if v != 0 {
							allZero = false
							break
						}
					}
					assert.False(t, allZero)
				}
			},
		)
		t.Run(
			fmt.Sprintf("%sEqual", tt.name),
			func(t *testing.T) {
				if tt.len == 0 {
					return
				}
				seed := tt.seed()
				b, errB := RandomBytesDeterministic(tt.len, seed)
				c, errC := RandomBytesDeterministic(tt.len, seed)

				if tt.err != nil {
					assert.EqualError(t, errB, tt.err.Error())
					assert.EqualError(t, errC, tt.err.Error())
				} else {
					assert.NoError(t, errB)
					assert.NoError(t, errC)
					assert.Equal(t, b, c)
				}
			},
		)
	}
}
