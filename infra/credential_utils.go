package infra

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/ademcaglin/authserver/models"
)

func MapJWKToPrivateKey(jwk models.PrivateJwkKey) (ecdsa.PrivateKey, error) {
	var curve elliptic.Curve
	switch jwk.Crv {
	case "P-256":
		curve = elliptic.P256()
	case "P-384":
		curve = elliptic.P384()
	}
	xInt, _ := base64.URLEncoding.DecodeString(jwk.X)
	yInt, _ := base64.URLEncoding.DecodeString(jwk.Y)
	dInt, _ := base64.URLEncoding.DecodeString(jwk.D)
	return ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     new(big.Int).SetBytes(xInt),
			Y:     new(big.Int).SetBytes(yInt),
		},
		D: new(big.Int).SetBytes(dInt),
	}, nil
}

func MapJWKToPublicKey(jwk models.PublicJwkKey) (ecdsa.PublicKey, error) {
	var curve elliptic.Curve
	switch jwk.Crv {
	case "P-256":
		curve = elliptic.P256()
	case "P-384":
		curve = elliptic.P384()
	}
	xInt, _ := base64.URLEncoding.DecodeString(jwk.X)
	yInt, _ := base64.URLEncoding.DecodeString(jwk.Y)
	return ecdsa.PublicKey{
		Curve: curve,
		X:     new(big.Int).SetBytes(xInt),
		Y:     new(big.Int).SetBytes(yInt),
	}, nil
}

func MapPrivateKeyToJWK(key ecdsa.PrivateKey, kid string) models.PrivateJwkKey {
	return models.PrivateJwkKey{
		PublicJwkKey: models.PublicJwkKey{
			Alg: "ES256",
			Kty: "EC",
			Crv: "P-256",
			Kid: kid,
			Use: "sig",
			X:   base64.URLEncoding.EncodeToString(key.X.Bytes()),
			Y:   base64.URLEncoding.EncodeToString(key.Y.Bytes()),
		},
		D: base64.URLEncoding.EncodeToString(key.D.Bytes()),
	}
}

func GenerateKey(kid string) (models.PrivateJwkKey, error) {
	curve := elliptic.P256()
	privateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	return MapPrivateKeyToJWK(*privateKey, kid), nil
}
