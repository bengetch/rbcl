package rbcl

import (
	"errors"
	"fmt"
)

var ErrBadSeedLength = fmt.Errorf("invalid seed length, need %d", RandomBytesSeedBytes())
var ErrBadHashLength = fmt.Errorf("invalid hash length, need %d", CryptoCoreRistretto255HashBytes())
var ErrBadPointLength = fmt.Errorf("invalid point length, need %d", CryptoCoreRistretto255Bytes())
var ErrBadScalarLength = fmt.Errorf("invalid scalar length, need %d", CryptoCoreRistretto255ScalarBytes())
var ErrBadNonReducedScalarLength = fmt.Errorf("invalid non reduced scalar length, need <= %d", CryptoCoreRistretto255NonReducedScalarBytes())
var ErrZeroScalar = errors.New("scalar cannot be zero")
