package rbcl

/*
#include "sodium/crypto_core_ristretto255.h"
*/
import "C"

func Ristretto255Bytes() int {
	return int(C.crypto_core_ristretto255_bytes())
}
