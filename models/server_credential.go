package models

type PublicJwkKey struct {
	Kid string `json:"kid"`
	X   string `json:"x"`
	Y   string `json:"y"`
	Kty string `json:"kty"`
	Crv string `json:"crv"`
	Alg string `json:"alg"`
	Use string `json:"use"`
}

type PrivateJwkKey struct {
	D string `json:"d"`
	PublicJwkKey
}

type ServerCredentialStore interface {
	GetPublicKey(kid string) (PublicJwkKey, error)
	GetPrivateKey(kid string) (PrivateJwkKey, error)
	Save(jwk PrivateJwkKey) error
}

/*func test() {
	x := PrivateJwkKey{
		PublicJwkKey: PublicJwkKey{
			X:   "",
			Y:   "",
			Kty: "",
			Crv: "",
			Alg: "",
			Use: "",
		},
		D: "",
	}
	x.Alg = "aa"
	print(x)
}*/
