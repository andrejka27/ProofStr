package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func generateString(secretKey, input string, length int) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(input))

	hashed := h.Sum(nil)

	hexString := hex.EncodeToString(hashed)

	var result string
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		index, _ := new(big.Int).SetString(hexString[i*2:i*2+2], 16)
		result += string(charset[int(index.Int64())%len(charset)])
	}

	return result
}

func printHelp() {
	fmt.Println("Usage: ./ProofStr [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  --secret string      Secret key to use for generating strings (required)")
	fmt.Println("  --num int            Number of strings to generate (default 5)")
	fmt.Println("  --length int         Length of the generated strings (default 7)")
	fmt.Println("  --help               Show this help menu")
}

func main() {
	secretKeyFlag := flag.String("secret", "", "Secret key to use for generating strings (required)")
	numStringsFlag := flag.Int("num", 5, "Number of strings to generate")
	lengthFlag := flag.Int("length", 7, "Length of the generated strings")
	helpFlag := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	if *secretKeyFlag == "" {
		fmt.Println("Error: Secret key is required.")
		printHelp()
		os.Exit(1)
	}

	for i := 1; i <= *numStringsFlag; i++ {
		str := generateString(*secretKeyFlag, strconv.Itoa(i), *lengthFlag)
		fmt.Println(str)
	}
}
