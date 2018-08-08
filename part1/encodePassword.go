package main
import (    "crypto/sha512"
            "encoding/base64"
        )

func encodePassword(password string) string {   //returns a base64 encoded version of the sha512 encrypted string.
    encrypted := sha512.Sum512([]byte(password))            //sha512 encrypt
    return base64.StdEncoding.EncodeToString(encrypted[:])  //base64 encode
}
