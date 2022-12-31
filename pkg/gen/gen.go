package gen

import (
	"strings"
	"crypto/rand"
	"math/big"
	"log"
	"github.com/smukm/go-genparoles/pkg/str"
)


func GenWords(iLines int) []string {

	// read random words from dictionary file
	s, err := getRandomLines("../../data/eng.txt", iLines*iLines)
	if err != nil {
		log.Fatal(err)
	}

	// make uppercase of random letter and put 4-letters words in slice
	parts := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		bz := str.GetPartOfWord(s[i], 4)
		parts[i] = str.UpperRandomLetter(strings.ToLower(bz[0]))
	}

	// put together parts of words
	words := make([]string, iLines)
	for i := 0; i < iLines; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(iLines*iLines)))
		ri, _ := rand.Int(rand.Reader, big.NewInt(int64(9)))
		words[i] = parts[r.Int64()] + ri.String()
	}

	return words
}