# GivingForward-Programming-Assessment : A simple decryption program for the purpose of code assessment

## Overview [![GoDoc](https://godoc.org/github.com/MidknightRyk/GivingForward-Programming-Assessment?status.svg)](https://godoc.org/github.com/MidknightRyk/GivingForward-Programming-Assessment)

## Install

```
go get github.com/MidknightRyk/GivingForward-Programming-Assessment
```

## Usage

The program accepts a JSON file as input. The JSON file should contain an array of the input sets, each with the following fields:
  - Name              
    *The name of the set of strings being decrypted. Feel free to name it anything, this is just for easier reading of the outputs if there are many input sets many sets*
    
  - Password
    *This is the 16-byte AES key to be used for decryption*
    
  - Encrypted Array
    *The array of strings to be decrypted*
    
Compile with `go build` and run it with the following command
```
./[decrypter] [input.json] [output.txt]
```

- [decrypter] is the name of the program
- [input.json] is the name of the JSON input file
- [output.txt] is the name of the file you would like the program to print the decrypted information

## Author

By Marishka Magness
For GivingForward

## License

Unlicense.

