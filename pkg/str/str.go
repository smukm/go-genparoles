package str

import (
	"crypto/rand"
	"math/big"
	"unicode"
)

func GetRandomChars(num int) []rune {

	allChars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	a := make([]rune, num)
	for i := 0; i < num; i++ {
		r,_ := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		a[i] = allChars[r.Int64()]
	}

	return a
}

func GetRandomStr(num int) string {
	return string(GetRandomChars(num))
}

func GetPartOfWord(word string, strlen int) []string {

	if len(word) <= strlen {
		return []string{word}
	}

	var part *string
	first := word[:strlen]
	part = &first

	return []string{*part}

}

// convert random letter to uppercase 
func UpperRandomLetter(s string) string {
	r := []byte(s)
	rs,_ := rand.Int(rand.Reader, big.NewInt(int64(len(r))));
	
	r[rs.Int64()] = byte(unicode.ToUpper(rune(r[rs.Int64()])))

	return string(r)
}
