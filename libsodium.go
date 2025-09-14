package rbcl

/*
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/prebuilt/x86_64-apple-darwin/libsodium.a -framework Security -framework CoreFoundation
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/prebuilt/arm64-apple-darwin/libsodium.a -framework Security -framework CoreFoundation
#cgo linux,amd64  LDFLAGS: ${SRCDIR}/prebuilt/x86_64-unknown-linux-gnu/libsodium.a -lm
#cgo linux,arm64  LDFLAGS: ${SRCDIR}/prebuilt/aarch64-unknown-linux-gnu/libsodium.a -lm

int crypto_scalarmult_ristretto255_bytes(void);
int crypto_scalarmult_ristretto255_scalarbytes(void);
int crypto_core_ristretto255_bytes(void);
int crypto_core_ristretto255_hashbytes(void);
int crypto_core_ristretto255_nonreducedscalarbytes(void);
int crypto_core_ristretto255_scalarbytes(void);
int randombytes_seedbytes(void);
int sodium_init(void);
*/
import "C"
import "fmt"

var (
	CryptoScalarMultRistretto255Bytes           int
	CryptoScalarMultRistretto255ScalarBytes     int
	CryptoCoreRistretto255Bytes                 int
	CryptoCoreRistretto255HashBytes             int
	CryptoCoreRistretto255NonReducedScalarBytes int
	CryptoCoreRistretto255ScalarBytes           int
	RandomBytesSeedBytes                        int
)

func init() {
	rc := C.sodium_init()
	if rc < 0 {
		panic(fmt.Errorf("libsodium initialization failed"))
	}

	CryptoScalarMultRistretto255Bytes = int(C.crypto_scalarmult_ristretto255_bytes())
	CryptoScalarMultRistretto255ScalarBytes = int(C.crypto_scalarmult_ristretto255_scalarbytes())
	CryptoCoreRistretto255Bytes = int(C.crypto_core_ristretto255_bytes())
	CryptoCoreRistretto255HashBytes = int(C.crypto_core_ristretto255_hashbytes())
	CryptoCoreRistretto255NonReducedScalarBytes = int(C.crypto_core_ristretto255_nonreducedscalarbytes())
	CryptoCoreRistretto255ScalarBytes = int(C.crypto_core_ristretto255_scalarbytes())
	RandomBytesSeedBytes = int(C.randombytes_seedbytes())
}
