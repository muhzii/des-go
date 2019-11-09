package main

import (
	"fmt"
	"math"
	"strings"
)

var PC1_TABLE [56]int = [56]int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

var SHIMT_TABLE [16]int = [16]int{
	1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1,
}

var PC2_TABLE [48]int = [48]int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

var IP_TABLE [64]int = [64]int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

var INVERSE_IP_TABLE [64]int = [64]int{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

var EP_TABLE [48]int = [48]int{
	32, 1, 2, 3, 4, 5, 4, 5,
	6, 7, 8, 9, 8, 9, 10, 11,
	12, 13, 12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21, 20, 21,
	22, 23, 24, 25, 24, 25, 26, 27,
	28, 29, 28, 29, 30, 31, 32, 1,
}

var S1 [4][16]int = [4][16]int{
	{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
	{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
	{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
	{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
}

var S2 [4][16]int = [4][16]int{
	{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
	{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
	{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
	{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
}

var S3 [4][16]int = [4][16]int{
	{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
	{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
	{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
	{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
}

var S4 [4][16]int = [4][16]int{
	{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
	{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
	{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
	{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
}

var S5 [4][16]int = [4][16]int{
	{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
	{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
	{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
	{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
}

var S6 [4][16]int = [4][16]int{
	{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
	{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
	{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
	{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
}

var S7 [4][16]int = [4][16]int{
	{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
	{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
	{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
	{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
}

var S8 [4][16]int = [4][16]int{
	{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
	{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
	{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
	{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
}

var PBOX_TABLE [32]int = [32]int{
	16, 7, 20, 21,
	29, 12, 28, 17,
	1, 15, 23, 26,
	5, 18, 31, 10,
	2, 8, 24, 14,
	32, 27, 3, 9,
	19, 13, 30, 6,
	22, 11, 4, 25,
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCF(a *[28]int, n int) [28]int { // Left circular shift
	result := *a

	for c := 0; c < gcd(n, len(a)); c++ {
		tmp := a[c]

		i := c
		for {
			j := i + n
			if j >= len(a) {
				j -= len(a)
			}
			if j == c {
				break
			}

			a[i] = a[j]
			i = j
		}
		a[i] = tmp
	}
	return result
}

func PC1(key *[64]int) [56]int { // Permuted choice 1
	var pkey [56]int
	for i, idx := range PC1_TABLE {
		pkey[i] = key[idx-1]
	}
	return pkey
}

func PC2(lkey, rkey *[28]int) [48]int { // Permuted choice 2
	result := [48]int{}
	for i, idx := range PC2_TABLE {
		pval := 0
		if idx > 28 {
			pval = rkey[idx-29]
		} else {
			pval = lkey[idx-1]
		}
		result[i] = pval
	}
	return result
}

func generateRoundKeys(key *[64]int) [16][48]int {
	pkey := PC1(key)

	var lkey, rkey [28]int
	for i := 0; i < 28; i++ {
		lkey[i] = pkey[i]
		rkey[i] = pkey[i+28]
	}

	roundKeys := [16][48]int{}
	for i := 0; i < 16; i++ {
		lkey = LCF(&lkey, SHIMT_TABLE[i])
		rkey = LCF(&rkey, SHIMT_TABLE[i])

		roundKeys[i] = PC2(&lkey, &rkey)
	}
	return roundKeys
}

func IP(text *[64]int) [64]int { // Initial permutation
	var ptext [64]int
	for i, idx := range IP_TABLE {
		ptext[i] = text[idx-1]
	}
	return ptext
}

func EP(p *[32]int) [48]int { // Expansion permutation
	result := [48]int{}
	for i, idx := range EP_TABLE {
		result[i] = p[idx-1]
	}
	return result
}

func XOR(p1 []int, p2 []int) []int {
	if len(p1) != len(p2) {
		fmt.Println("Cannot XOR inputs of different length!")
		return nil
	}

	result := make([]int, len(p1))
	for i := 0; i < len(p1); i++ {
		result[i] = p1[i] ^ p2[i]
	}
	return result
}

func SBOX(p *[48]int) [32]int {
	result := [32]int{}
	for i := 0; i < 48; i += 6 {
		row := 2*p[i] + p[i+5]

		col := 0
		for j := 0; j < 4; j++ {
			col += p[i+j+1] * int(math.Pow(2, float64(3-j)))
		}

		var v int
		switch i / 6 {
		case 0:
			v = S1[row][col]
		case 1:
			v = S2[row][col]
		case 2:
			v = S3[row][col]
		case 3:
			v = S4[row][col]
		case 4:
			v = S5[row][col]
		case 5:
			v = S6[row][col]
		case 6:
			v = S7[row][col]
		case 7:
			v = S8[row][col]
		}

		for j := 0; j < 4; j++ {
			result[(i/6)*4+3-j] = (v >> uint(j)) & 1
		}
	}
	return result
}

func PBOX(a *[32]int) [32]int { // Permutation box
	result := [32]int{}
	for i, idx := range PBOX_TABLE {
		result[i] = a[idx-1]
	}
	return result
}

func inverseIP(text *[64]int) [64]int { // inverse of IP
	var ptext [64]int
	for i, idx := range INVERSE_IP_TABLE {
		ptext[i] = text[idx-1]
	}
	return ptext
}

func DESEncipher(text *[64]int, roundKeys *[16][48]int) string {
	ptext := IP(text)

	var lh, rh [32]int
	for i := 0; i < 32; i++ {
		lh[i] = ptext[i]
		rh[i] = ptext[i+32]
	}

	for i := 0; i < 16; i++ {
		expandedRh := EP(&rh)

		xorWithKey := [48]int{}
		copy(xorWithKey[:], XOR(expandedRh[:], roundKeys[i][:]))

		sboxRh := SBOX(&xorWithKey)
		mangledRh := PBOX(&sboxRh)

		tmph := rh
		copy(rh[:], XOR(lh[:], mangledRh[:]))
		lh = tmph
	}

	var final [64]int
	for i := 0; i < 32; i++ {
		final[i] = rh[i]
		final[i+32] = lh[i]
	}
	final = inverseIP(&final)

	return binToHex(final[:])
}

func encrypt(text *[64]int, key *[64]int) string {
	rkeys := generateRoundKeys(key)
	return DESEncipher(text, &rkeys)
}

func decrypt(text *[64]int, key *[64]int) string {
	rkeys := generateRoundKeys(key)
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		rkeys[i], rkeys[j] = rkeys[j], rkeys[i]
	}
	return DESEncipher(text, &rkeys)
}

func binToHex(a []int) string {
	if len(a)%4 != 0 {
		fmt.Println("Cannot convert inconsistent array to hex")
		return ""
	}

	result := ""
	for i := 0; i < len(a); i += 4 {
		v := 0
		for j := 0; j < 4; j++ {
			v += a[i+j] * int(math.Pow(2, float64(3-j)))
		}

		if v < 10 {
			result += string('0' + v)
		} else {
			result += string('A' + v - 10)
		}
	}
	return result
}

func hexToBin(s string) []int {
	res := make([]int, len(s)*4)
	for i, ch := range s {
		v := 0
		if ch > 47 && ch < 58 {
			v = int(ch) - '0'
		} else if ch > 64 && ch < 71 {
			v = int(ch) - 'A' + 10
		}

		for j := i*4 + 3; j >= i*4; j-- {
			res[j] = v & 1
			v = v >> 1
		}
	}
	return res
}

func main() {
	var n, mode int
	var textStr, keyStr string

	fmt.Print("Enter the key: ")
	fmt.Scanf("%s", &keyStr)
	if len(keyStr) != 16 {
		fmt.Println("Invalid key length")
		return
	}

	fmt.Print("Enter the plaintext: ")
	fmt.Scanf("%s", &textStr)
	if len(textStr) != 16 {
		fmt.Println("Invalid input length")
		return
	}

	fmt.Print("Enter the number of times to run the encryption: ")
	fmt.Scanf("%d", &n)

	fmt.Println("You wish to do:")
	fmt.Println("1. Encrpytion")
	fmt.Println("2. Decryption")
	fmt.Scanf("%d", &mode)

	var text [64]int
	var key [64]int
	copy(text[:], hexToBin(strings.ToUpper(textStr)))
	copy(key[:], hexToBin(strings.ToUpper(keyStr)))

	var result string
	if mode == 1 {
		result = encrypt(&text, &key)
	} else {
		result = decrypt(&text, &key)
	}
	for i := 0; i < n-1; i++ {
		var r [64]int
		copy(r[:], hexToBin(result))
		if mode == 1 {
			result = encrypt(&r, &key)
		} else {
			result = decrypt(&r, &key)
		}
	}

	fmt.Println(result)
}
