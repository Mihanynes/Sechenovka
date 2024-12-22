package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		hash := GenerateImageURL(i)
		fmt.Println("Encoded: ", hash)
		//decoded, _ := DecodeImageURL(hash)
		//fmt.Println("Decoded: ", decoded)
	}
}
