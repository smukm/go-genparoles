package gen

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/smukm/go-genparoles/pkg/str"
)


func GenWords(iLines int, passLen int) []string {
	iLines *= 2
	passLen = passLenValidate(passLen)

	_, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }

	// read random words from dictionary file
	s, err := getRandomLines(filepath.Join(filepath.Dir(filename),"data/eng.txt"), iLines*iLines)

	if err != nil {
		log.Fatal(err)
	}

	partLen := math.Ceil(float64(passLen/2)) - 1
	// make uppercase of random letter and put 4-letters words in slice
	parts := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		bz := str.GetPartOfWord(s[i], int(partLen))
		parts[i] = str.UpperRandomLetter(strings.ToLower(bz[0]))
	}

	// put together parts of words
	words := make([]string, iLines)
	for i := 0; i < iLines; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(iLines*iLines)))
		ri, _ := rand.Int(rand.Reader, big.NewInt(int64(9)))
		words[i] = parts[r.Int64()] + ri.String()
	}

	var ret []string
	for i := 0; i<iLines; i=i+2 {
		ret = append(ret, words[i] + words[i+1])
	}

	return ret
}

func passLenValidate(passLen int) int {
	if(passLen < 6) {
		passLen = 6
	}
	if(passLen > 12) {
		passLen = 12
	}
	return passLen
}