package rbcl

/*
int crypto_scalarmult_ristretto255_bytes(void);
int crypto_scalarmult_ristretto255_scalarbytes(void);
int crypto_core_ristretto255_bytes(void);
int crypto_core_ristretto255_hashbytes(void);
int crypto_core_ristretto255_nonreducedscalarbytes(void);
int crypto_core_ristretto255_scalarbytes(void);
int randombytes_seedbytes(void);
*/
import "C"

func CryptoScalarMultRistretto255Bytes() int {
	return int(C.crypto_scalarmult_ristretto255_bytes())
}

func CryptoScalarMultRistretto255ScalarBytes() int {
	return int(C.crypto_scalarmult_ristretto255_scalarbytes())
}

func CryptoCoreRistretto255Bytes() int {
	return int(C.crypto_core_ristretto255_bytes())
}

func CryptoCoreRistretto255HashBytes() int {
	return int(C.crypto_core_ristretto255_hashbytes())
}

func CryptoCoreRistretto255NonReducedScalarBytes() int {
	return int(C.crypto_core_ristretto255_nonreducedscalarbytes())
}

func CryptoCoreRistretto255ScalarBytes() int {
	return int(C.crypto_core_ristretto255_scalarbytes())
}

func RandomBytesSeedBytes() int {
	return int(C.randombytes_seedbytes())
}
