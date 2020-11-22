package infra_test

import (
	"encoding/base64"
	"math/big"
	"testing"

	"github.com/ademcaglin/authserver/infra"
	"github.com/ademcaglin/authserver/models"
)

func TestGetPrivateKey(t *testing.T) {
	pkey, _ := infra.GetPrivateKey(models.PrivateJwkKey{
		PublicJwkKey: models.PublicJwkKey{
			Crv: "P-256",
			X:   "IA67J_L3h-C5ZekNxza0OhQYIev5m-5v28abFiss2ro",
			Y:   "0Kzv3-fOwzAqBTRwLZbjVEFal4sFfd4qo1Nofcu36rY",
		},
		D: "geAOPDdt7WIWKxA6XaiMND_9AeVEUTKo_VRPtjkBLwg",
	})
	urlX := base64.URLEncoding.EncodeToString(pkey.X.Bytes())
	urlY := base64.URLEncoding.EncodeToString(pkey.Y.Bytes())
	x := base64.StdEncoding.EncodeToString(pkey.X.Bytes())

	t.Logf("urlX key is : %v", urlX)
	t.Logf("bit size : %v", pkey.Curve.Params().BitSize)
	t.Logf("X key is : %v", x)
	t.Logf("Y key is : %v", urlY)
	t.Logf("Private key is : %v", pkey.Curve.Params().Name)
}

func TestAbc(t *testing.T) {
	x, _ := base64.RawURLEncoding.DecodeString("IA67J_L3h-C5ZekNxza0OhQYIev5m-5v28abFiss2ro")
	xInt := new(big.Int).SetBytes(x)
	t.Logf("xInt is : %v", xInt)

	x2, _ := base64.URLEncoding.DecodeString("IA67J_L3h-C5ZekNxza0OhQYIev5m-5v28abFiss2ro")
	x2Int := new(big.Int).SetBytes(x2)
	t.Logf("x2Int is : %v", x2Int)
	//xBytes := newFixedSizeBuffer(xInt.Bytes(), 32)
	//xbytes := make([]byte, 32)

	xStr := base64.RawURLEncoding.EncodeToString(xInt.Bytes())
	t.Logf("key is : %v", xStr)

	yInt := new(big.Int).SetInt64(43657343623242454) //privkey.X
	yBase64 := base64.URLEncoding.EncodeToString(yInt.Bytes())
	t.Logf("yBase64 is : %v", yBase64)
	//privkey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	/*xInt := new(big.Int).SetInt64(436573436) //privkey.X
	xBase64 := base64.URLEncoding.EncodeToString(xInt.Bytes())
	xIntBytes, _ := base64.URLEncoding.DecodeString(xBase64)
	xIntNew := new(big.Int).SetBytes(xIntBytes)
	xBytes, _ := base64.RawURLEncoding.DecodeString("IA67J_L3h-C5ZekNxza0OhQYIev5m-5v28abFiss2ro")
	t.Logf("x is : %v", xInt)
	t.Logf("new x is : %v", xIntNew)
	t.Logf("x bytes is : %v", xBytes)*/
}

// Align left pads given byte array with zeros till it have at least bitSize length.
func newFixedSizeBuffer(data []byte, length int) *byteBuffer {
	if len(data) > length {
		panic("square/go-jose: invalid call to newFixedSizeBuffer (len(data) > length)")
	}
	pad := make([]byte, length-len(data))
	return newBuffer(append(pad, data...))
}

type byteBuffer struct {
	data []byte
}

func newBuffer(data []byte) *byteBuffer {
	if data == nil {
		return nil
	}
	return &byteBuffer{
		data: data,
	}
}
