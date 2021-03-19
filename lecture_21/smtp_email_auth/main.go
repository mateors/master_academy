package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	fmt.Println(base64Encode("AnyTextYouWantBase64Encoded"))
}

func base64Encode(strtxt string) (encodedStr string) {

	//bnJqZ3pwcmlkcG14d3lreQ==
	encodedStr = base64.StdEncoding.EncodeToString([]byte(strtxt))
	return
}

func gmailPlainTextAuthCode() {

	//for generating gmail plain text authentication
	resp := []byte("\x00" + "yourgmailuser@gmail.com" + "\x00" + "gmailPassword")
	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)

}
