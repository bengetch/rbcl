package rbcl

/*
void crypto_core_ristretto255_scalar_random(unsigned char *r);
void crypto_core_ristretto255_scalar_reduce(unsigned char *r, const unsigned char *s);
void crypto_core_ristretto255_scalar_negate(unsigned char *neg, const unsigned char *s);
void crypto_core_ristretto255_scalar_complement(unsigned char *comp, const unsigned char *s);
int crypto_core_ristretto255_scalar_invert(unsigned char *recip, const unsigned char *s);
void crypto_core_ristretto255_scalar_add(unsigned char *z, const unsigned char *x, const unsigned char *y);
void crypto_core_ristretto255_scalar_sub(unsigned char *z, const unsigned char *x, const unsigned char *y);
void crypto_core_ristretto255_scalar_mul(unsigned char *z, const unsigned char *x, const unsigned char *y);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// CryptoCoreRistretto255ScalarRandom a new random valid Ristretto255 scalar
func CryptoCoreRistretto255ScalarRandom() []byte {
	buf := make([]byte, CryptoCoreRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_random(
		(*C.uchar)(unsafe.Pointer(&buf[0])),
	)
	return buf
}

// CryptoCoreScalarRistretto255ScalarReduce returns the reduced representation s modulo L of a
// scalar, where L is the order of the main subgroup
func CryptoCoreScalarRistretto255ScalarReduce(s []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255NonReducedScalarBytes()
	if len(s) > reqLen {
		return nil, ErrBadNonReducedScalarLength
	}

	out := make([]byte, CryptoCoreRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_reduce(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
	)
	return out, nil
}

// CryptoCoreRistretto255ScalarNegate returns the additive inverse of the scalar s modulo L, i.e.,
// a scalar t such that s + t == 0 modulo L, where L is the order of the main subgroup
func CryptoCoreRistretto255ScalarNegate(s []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if len(s) != reqLen {
		return nil, ErrBadScalarLength
	}

	out := make([]byte, CryptoCoreRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_negate(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
	)
	return out, nil
}

// CryptoCoreRistretto255ScalarComplement returns the additive complement of the scalar s modulo L, i.e.,
// a scalar t such that s + t == 1 modulo L, where L is the order of the main subgroup
func CryptoCoreRistretto255ScalarComplement(s []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if len(s) != reqLen {
		return nil, ErrBadScalarLength
	}

	out := make([]byte, CryptoCoreRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_complement(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
	)
	return out, nil
}

func isZero(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}

// CryptoCoreRistretto255ScalarInvert returns the multiplicative inverse of the scalar s modulo L,
// i.e., an integer t such that s * t == 1 modulo L, where L is the order of the main subgroup
func CryptoCoreRistretto255ScalarInvert(s []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if len(s) != reqLen {
		return nil, ErrBadScalarLength
	}
	if isZero(s) {
		return nil, ErrZeroScalar
	}

	out := make([]byte, CryptoCoreRistretto255ScalarBytes())
	rc := C.crypto_core_ristretto255_scalar_invert(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
	)
	if rc != 0 {
		return nil, fmt.Errorf("unexpected nonzero return from libsodium: %d", int(rc))
	}
	return out, nil
}

// CryptoCoreRistretto255ScalarAdd returns the addition of two scalars s and t modulo L, where
// L is the order of the main subgroup
func CryptoCoreRistretto255ScalarAdd(s, t []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if (len(s) != reqLen) || (len(t) != reqLen) {
		return nil, ErrBadScalarLength
	}

	out := make([]byte, CryptoCoreRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_add(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
		(*C.uchar)(unsafe.Pointer(&t[0])),
	)
	return out, nil
}

// CryptoCoreRistretto255ScalarSub returns the difference between two scalars s and t modulo L,
// where L is the order of the main subgroup
func CryptoCoreRistretto255ScalarSub(s, t []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if (len(s) != reqLen) || (len(t) != reqLen) {
		return nil, ErrBadScalarLength
	}

	out := make([]byte, CryptoCoreRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_sub(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
		(*C.uchar)(unsafe.Pointer(&t[0])),
	)
	return out, nil
}

// CryptoCoreRistretto255ScalarMul returns the result of multiplying two scalars s and t modulo
// L, where L is the order of the main subgroup
func CryptoCoreRistretto255ScalarMul(s, t []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if (len(s) != reqLen) || (len(t) != reqLen) {
		return nil, ErrBadScalarLength
	}

	out := make([]byte, CryptoScalarMultRistretto255ScalarBytes())
	C.crypto_core_ristretto255_scalar_mul(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
		(*C.uchar)(unsafe.Pointer(&t[0])),
	)
	return out, nil
}
