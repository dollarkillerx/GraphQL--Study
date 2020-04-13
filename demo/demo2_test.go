package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestDemo2(t *testing.T) {
	decodeString, err := base64.StdEncoding.DecodeString("Y3Vyc29yMTA=")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeString))
}
