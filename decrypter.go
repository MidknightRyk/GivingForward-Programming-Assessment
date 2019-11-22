// Code is written by Marishka Magness
// For the purpose the GivingForward Programming Assessment
// Last Edited 23 November 2019

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Struct of encrypted text and Password
type encSet struct {
	Name     string
	Password string
	EncText  []string `json:"Encrypted Array"`
}

func main() {
	// Obfuscated Password Chunk Length
	const obsStep int = 3

	// Get input and output file name from args
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := ioutil.ReadFile(inputFile)
	check(err)

	//Prep array containing the encrypted texts and password
	var encArr []encSet

	err = json.Unmarshal([]byte(data), &encArr)
	check(err)

	/*
			f, err := os.Create("test.txt")
		    if err != nil {
		        fmt.Println(err)
		        return
		  }

	*/

	if _, err := os.Stat(outputFile); err == nil {
		fmt.Println("Output File Exists in directory, File will be replaced.")
		err = os.Remove(outputFile)
		check(err)
	}
	outputText, err := os.Create(outputFile)
	check(err)

	for _, set := range encArr {

		// Decrypt the string array
		decryptedPass := decryptPass(set.Password, obsStep)
		decryptedArr := decryptArr(set.EncText, []byte(decryptedPass))
		// Write decrypted array to file.
		fmt.Fprintf(outputText, "%s:\n\n %v\n\n", set.Name, decryptedArr)
	}

	// Close output file
	outputText.Close()
}

func decryptArr(encArr []string, pass []byte) []string {
	decryptedArray := []string{}
	decryptedNum := []int{}
	decryptedStrs := []string{}

	for i := 0; i < len(encArr); i++ {
		// Decrypt Strings
		decryptedStr := decryptStr(encArr[i], []byte(pass))

		// Convert Binary/Hex numbers to Decimals and append to num array
		if num, err := strconv.ParseInt(decryptedStr, 2, 64); err == nil {
			decryptedNum = append(decryptedNum, int(num))
		} else if num, err := strconv.ParseInt(decryptedStr, 0, 64); err == nil {
			decryptedNum = append(decryptedNum, int(num))
		} else {
			// Append decrypted string to the final Array
			decryptedStrs = append(decryptedStrs, strings.Fields(decryptedStr)...)
		}
	}

	// Sort the arrays and combine
	sort.Ints(decryptedNum)
	sort.Strings(decryptedStrs)
	for _, num := range decryptedNum {
		text := strconv.Itoa(num)
		decryptedArray = append(decryptedArray, text)
	}
	decryptedArray = append(decryptedArray, decryptedStrs...)
	return decryptedArray
}

// Function decrypts a given string and password of base64
// Returns the decrypted and decoded text as a string
func decryptStr(str string, pass []byte) string {
	// Convert the password to bytes
	text, _ := base64.StdEncoding.DecodeString(str)

	block, err := aes.NewCipher(pass)
	check(err)

	if len(text) < aes.BlockSize {
		panic("text is too short, improperly padded")
	}

	// Decrypt the text given
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	dec := cipher.NewCBCDecrypter(block, iv)
	dec.CryptBlocks(text, text)

	// Remove padding
	text = text[:len(text)-int(text[len(text)-1])]
	decodedText := fmt.Sprintf("%s", text)
	return decodedText
}

// Function takes the obfuscated password
// Returns the unobfuscated password
func decryptPass(encPass string, obsStep int) string {
	// Extract the obfuscated part of the password
	clipSize := len(encPass) / obsStep
	chunk := obsStep - 1
	clip := encPass[:clipSize]
	preClip := encPass[clipSize:]
	decryptedPass := ""

	// Replace the obfuscated clip to the correct positions
	for i := 0; i < clipSize; i++ {
		temp := preClip[:chunk]
		preClip = preClip[chunk:]
		temp2 := fmt.Sprintf("%c", clip[i])
		decryptedPass = decryptedPass + temp + temp2
	}

	//Append the trailing password to the final password
	decryptedPass = decryptedPass + preClip
	return decryptedPass
}

// Error Checking Function
func check(err error) {
	if err != nil {
		panic(err)
	}
}
