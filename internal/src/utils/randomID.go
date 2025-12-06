package utils

import "math/rand"

//creates random ID
func RandomIDGenerate(NAME string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Convert := make([]byte, 10)

	for i := range Convert {
		Convert[i] = chars[rand.Intn(len(chars))]
	}
	DATA := NAME + string(Convert)
	return string(DATA)
}
