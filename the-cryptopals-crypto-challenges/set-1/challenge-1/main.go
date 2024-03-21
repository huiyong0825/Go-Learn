package main 

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
)

func main(){
	const hexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const answer string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// decoding hex string to byte
	hexToByte, err := hex.DecodeString(hexString)
	
	if(err != nil){
		fmt.Println("Err:", err)
		return
	}

	// encoding byte to base64 string
	byteToB64 := base64.StdEncoding.EncodeToString(hexToByte)

	if(byteToB64 != answer){
		fmt.Println("Not Match!")
	} else{
		fmt.Println("Match!")
	}
}