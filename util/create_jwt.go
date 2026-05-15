package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type PayLoad struct {
	Sub         int `json:"sub"` //id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string,data PayLoad) (string ,error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteHeaderArray, err := json.Marshal(header)
	if err!=nil{
		return "",err
	}
	headerB64:=base64UrlEncode(byteHeaderArray)
    byteDataArray,err:=json.Marshal(data)
	if err!=nil{
		return "",err
	}
	payloadB64:=base64UrlEncode(byteDataArray)

	message:=headerB64+"."+payloadB64
    
	byteSecret:=[]byte(secret)
	byteMessage:=[]byte(message)

	h:=hmac.New(sha256.New,byteSecret)
	h.Write(byteMessage)

	signature:=h.Sum(nil)

	signatureB64:=base64UrlEncode(signature)

	jwt:=headerB64+"."+payloadB64+"."+signatureB64

	return jwt,nil


}
// base64 Encoder 
func base64UrlEncode(data []byte) string{
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
