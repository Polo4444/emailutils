package emailutils

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

// RandomNum generates random numbers
func RandomNum(nbOfNumbers int) string {

	sResult := ""

	// We init the random seed
	rSource, _ := uuid.Must(uuid.NewRandom()).Time().UnixTime()
	r := rand.New(rand.NewSource(rSource))
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// build the string
	for i := 1; i <= nbOfNumbers; i++ {
		sResult += fmt.Sprintf("%d", r.Intn(10))
	}
	return sResult
}

// RandomNumInt generates random numbers
func RandomNumInt(nbOfNumbers int) int {
	num, _ := strconv.Atoi(RandomNum(nbOfNumbers))
	return num
}
