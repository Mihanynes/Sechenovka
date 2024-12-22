package main

import (
	"encoding/base64"
	"strconv"
)

func GenerateImageURL(questionId int) string {
	encodedId := base64.URLEncoding.EncodeToString([]byte("question:" + strconv.Itoa(questionId)))

	return encodedId
}

//func DecodeImageURL(imageURL string) (int, error) {
//	decodedBytes, err := base64.URLEncoding.DecodeString(imageURL)
//	if err != nil {
//		return 0, err
//	}
//	return strconv.Atoi(string(decodedBytes))
//
//}
