package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`
____ ____    _  _ ____ ____ _  _ ____ ____
| __ |  |    |__| |__| [__  |__| |___ |__/
|__] |__|    |  | |  | ___] |  | |___ |  \

`)
	fmt.Print("Raw Input: ")
	scanner.Scan()
	rawString := scanner.Text()
	fmt.Println("What type of hash do you want to use?")
	fmt.Println("[1] MD5")
	fmt.Println("[2] SHA-256")
	fmt.Println("[3] SHA-512")
	fmt.Print("-> ")
	scanner.Scan()
	hashType := scanner.Text()
	if hashType == "1" {
		fmt.Println("Hash (MD5): " + MD5(rawString))
		end()
	} else if hashType == "2" {
		fmt.Println("Hash (SHA-256): " + SHA256(rawString))
		end()
	} else if hashType == "3" {
		fmt.Println("Hash (SHA-512): " + SHA512(rawString))
		end()
	} else {
		fmt.Println("Invalid option")
		end()
	}
}

func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func SHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Kind of unnecessary but it minimizes the code
func end() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
