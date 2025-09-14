package rbcl

/*
void randombytes(unsigned char * const buf, const unsigned long long buf_len);
void randombytes_buf_deterministic(void * const buf, const size_t size, const unsigned char *seed);
*/
import "C"
import (
	"unsafe"
)

// RandomBytes returns a slice of length n filled with random bytes
func RandomBytes(n int) []byte {
	buf := make([]byte, n)
	if n > 0 {
		C.randombytes(
			(*C.uchar)(unsafe.Pointer(&buf[0])),
			C.ulonglong(n),
		)
	}
	return buf
}

// RandomBytesDeterministic returns a slice of length n filled with random bytes derived
// deterministically from a seed
func RandomBytesDeterministic(n int, seed []byte) ([]byte, error) {
	if len(seed) != RandomBytesSeedBytes {
		return nil, ErrBadSeedLength
	}

	buf := make([]byte, n)
	if n > 0 {
		C.randombytes_buf_deterministic(
			unsafe.Pointer(&buf[0]),
			C.size_t(n),
			(*C.uchar)(unsafe.Pointer(&seed[0])),
		)
	}
	return buf, nil
}
