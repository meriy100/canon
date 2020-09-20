package configs

import "os"

func GetSecretKey() []byte {
	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) > 1 {
		return []byte(secretKey)
	}
	return []byte("secret")
}

