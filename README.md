# Caesar Cipher
A Go program that concurrently implements the Caesar Cipher algorithm on a list of strings.

## Introduction
The Caesar Cipher is a simple encryption technique in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet. This program takes a list of strings and applies the Caesar Cipher to each word concurrently, returning the resulting encrypted words.

## Usage
To use this program, simply run the following command:


`go run main.go`

The program will apply the Caesar Cipher to each word in the predefined messages list concurrently, with a shift of 2, and print the encrypted words to the console.

You can modify the messages list and the shift value to your liking.

## Example
For example, given the following input list of strings:


```
messages := []string{"Csi2520", "Csi2120", "3 Paradigms",
"Go is 1st", "Prolog is 2nd", "Scheme is 3rd",
"uottawa.ca", "csi/elg/ceg/seg", "800 King Edward"}
```
The program will output the following encrypted words:


```
EUK2742
EUK2328
3 RCFCTKIUO
IQ KU 1UV
Trqogn ku 2pf
Ukecog ku 3tf
WQVQCY.CC
euk/gni/egi/ugj
800 Mqpi Gfgigt
```

Note that the Caesar Cipher is applied concurrently to each word, thanks to the use of channels and goroutines in the CaesarCipherList function.

## Contributing
If you find a bug or have a suggestion for improvement, feel free to submit an issue or pull request.
