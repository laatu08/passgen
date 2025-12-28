package generator

import (
	"crypto/rand"
	"math/big"
	"errors"
)

// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	digitsSet = "0123456789"
	symbolsSet = "!@#$%^&*()-_=+[]{}<>?/|"
)

func Generate(length int, upper, lower, digits, symbols bool)(string, error){
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	
	charset := ""

	if upper {
		charset += uppercase
	}
	if lower {
		charset += lowercase
	}
	if digits {
		charset += digitsSet
	}
	if symbols {
		charset += symbolsSet
	}

	if charset == "" {
		charset="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
		// return "", errors.New("select at least one character type")
	}

	password:=make([]byte, length)

	for i := 0; i < length; i++ {
		n,err:=rand.Int(rand.Reader,big.NewInt(int64(len(charset))))
		if err!=nil{
			return "",err
		}
		password[i]=charset[n.Int64()]
	}

	return string(password),nil
}