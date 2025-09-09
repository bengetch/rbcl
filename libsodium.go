package rbcl

/*
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/prebuilt/darwin-amd64/libsodium.a -framework Security -framework CoreFoundation

#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/prebuilt/darwin-arm64/libsodium.a -framework Security -framework CoreFoundation
*/
import "C"
