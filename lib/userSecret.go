package lib

import "fmt"

var secrets struct {
	SomeSecret    string
	AnotherSecret string
}

func GetSecrets() string{
	return fmt.Sprintf("lib stuff - secret1: %s, secret2: %s", secrets.SomeSecret, secrets.AnotherSecret)
}
