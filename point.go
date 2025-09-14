package rbcl

/*
int crypto_core_ristretto255_is_valid_point(const unsigned char *p);
void crypto_core_ristretto255_random(unsigned char *p);
int crypto_core_ristretto255_from_hash(unsigned char *p, const unsigned char *r);
int crypto_core_ristretto255_add(unsigned char *r, const unsigned char *p, const unsigned char *q);
int crypto_core_ristretto255_sub(unsigned char *r, const unsigned char *p, const unsigned char *q);
int crypto_scalarmult_ristretto255_base(unsigned char *q, const unsigned char *n);
int crypto_scalarmult_ristretto255(unsigned char *q, const unsigned char *n, const unsigned char *p);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// CryptoCoreRistretto255IsValidPoint returns true if p encodes a valid Ristretto255 point
func CryptoCoreRistretto255IsValidPoint(p []byte) bool {
	if len(p) != CryptoCoreRistretto255Bytes() {
		return false
	}

	rc := C.crypto_core_ristretto255_is_valid_point(
		(*C.uchar)(unsafe.Pointer(&p[0])),
	)
	return rc != 0
}

// CryptoCoreRistretto255Random returns a new random valid Ristretto255 point
func CryptoCoreRistretto255Random() []byte {
	buf := make([]byte, CryptoCoreRistretto255Bytes())
	C.crypto_core_ristretto255_random(
		(*C.uchar)(unsafe.Pointer(&buf[0])),
	)
	return buf
}

// CryptoCoreRistretto255FromHash maps a 64-byte hash to a Ristretto255 point
func CryptoCoreRistretto255FromHash(h []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255HashBytes()
	if len(h) != reqLen {
		return nil, ErrBadHashLength
	}

	out := make([]byte, CryptoCoreRistretto255Bytes())
	rc := C.crypto_core_ristretto255_from_hash(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&h[0])),
	)
	if rc != 0 {
		return nil, fmt.Errorf("unexpected nonzero return from libsodium: %d", int(rc))
	}
	return out, nil
}

// CryptoCoreRistretto255Add returns the sum of two Ristretto255 points
func CryptoCoreRistretto255Add(p []byte, q []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255Bytes()
	if (len(p) != reqLen) || (len(q) != reqLen) {
		return nil, ErrBadPointLength
	}

	out := make([]byte, CryptoCoreRistretto255Bytes())
	rc := C.crypto_core_ristretto255_add(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&p[0])),
		(*C.uchar)(unsafe.Pointer(&q[0])),
	)
	if rc != 0 {
		return nil, fmt.Errorf("unexpected nonzero return from libsodium: %d", int(rc))
	}
	return out, nil
}

// CryptoCoreRistretto255Sub returns the difference between two Ristretto255 points
func CryptoCoreRistretto255Sub(p []byte, q []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255Bytes()
	if (len(p) != reqLen) || (len(q) != reqLen) {
		return nil, ErrBadPointLength
	}

	out := make([]byte, CryptoCoreRistretto255Bytes())
	rc := C.crypto_core_ristretto255_sub(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&p[0])),
		(*C.uchar)(unsafe.Pointer(&q[0])),
	)
	if rc != 0 {
		return nil, fmt.Errorf("unexpected nonzero return from libsodium: %d", int(rc))
	}
	return out, nil
}

// CryptoScalarMultRistretto255Base returns the product of a standard group element and a scalar s
func CryptoScalarMultRistretto255Base(s []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if len(s) != reqLen {
		return nil, ErrBadScalarLength
	}

	out := make([]byte, CryptoScalarMultRistretto255Bytes())
	rc := C.crypto_scalarmult_ristretto255_base(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
	)
	if rc != 0 {
		return nil, fmt.Errorf(
			"input cannot be larger than the size of the group and cannot yield "+
				"the identity element when applied as an exponent: %d", int(rc),
		)
	}
	return out, nil
}

// CryptoScalarMultRistretto255 returns the product of a clamped scalar s and the provided point.
// The scalar is clamped, as done in the public key generation case, by setting to zero the bits
// in position [0, 1, 2, 255] and by setting to 1 the bit in position 254
func CryptoScalarMultRistretto255(s, p []byte) ([]byte, error) {
	reqLen := CryptoCoreRistretto255ScalarBytes()
	if len(s) != reqLen {
		return nil, ErrBadScalarLength
	}
	reqLen = CryptoCoreRistretto255Bytes()
	if len(p) != reqLen {
		return nil, ErrBadPointLength
	}

	out := make([]byte, CryptoScalarMultRistretto255Bytes())
	rc := C.crypto_scalarmult_ristretto255(
		(*C.uchar)(unsafe.Pointer(&out[0])),
		(*C.uchar)(unsafe.Pointer(&s[0])),
		(*C.uchar)(unsafe.Pointer(&p[0])),
	)
	if rc != 0 {
		return nil, fmt.Errorf(
			"input cannot be larger than the size of the group and cannot yield "+
				"the identity element when applied as an exponent: %d", int(rc),
		)
	}
	return out, nil
}
