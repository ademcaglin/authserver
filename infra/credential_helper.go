package infra

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"math/big"

	"github.com/ademcaglin/authserver/models"
)

func GetPrivateKey(jwk models.PrivateJwkKey) (ecdsa.PrivateKey, error) {
	var curve elliptic.Curve
	switch jwk.Crv {
	case "P-256":
		curve = elliptic.P256()
	case "P-384":
		curve = elliptic.P384()
	}
	xInt, _ := base64.RawURLEncoding.DecodeString(jwk.X)
	yInt, _ := base64.RawURLEncoding.DecodeString(jwk.Y)
	dInt, _ := base64.RawURLEncoding.DecodeString(jwk.D)
	return ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     new(big.Int).SetBytes(xInt),
			Y:     new(big.Int).SetBytes(yInt),
		},
		D: new(big.Int).SetBytes(dInt),
	}, nil
}
