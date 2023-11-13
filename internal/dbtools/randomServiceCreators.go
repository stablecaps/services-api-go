package dbtools

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/tjarratt/babble"
)


func MakeRandomName() string {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	cleanedName := strings.Replace(babbler.Babble(), "'", -1)
	return cleanedName
}

func MakeRandomDescription(wordCount int) string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = wordCount
	return babbler.Babble()
}

func MakeRandomService() []byte {
	randomName := MakeRandomName()
	numWords := rand.Intn(10)
	radomDesc := MakeRandomDescription(numWords)

	log.Printf("numWords is %d", numWords)
	log.Printf("service randomName is %s", randomName)
	log.Printf("service radomDesc is %s", radomDesc)

	body := []byte(fmt.Sprintf(`{
		"serviceName": "%s",
		"serviceDescription": "%s"
	}`, randomName, radomDesc) )

	return body
}