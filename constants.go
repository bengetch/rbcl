package rbcl

/*
int crypto_core_ristretto255_bytes(void);
*/
import "C"

func Ristretto255Bytes() int {
	return int(C.crypto_core_ristretto255_bytes())
}
