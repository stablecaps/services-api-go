package dbtools

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/tjarratt/babble"
)


func makeRandomName() string {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	return babbler.Babble()
}

func makeRandomDescription(wordCount int) string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = wordCount
	return babbler.Babble()
}

func MakeRandomService() []byte {
	randomName := makeRandomName()
	numWords := rand.Intn(10)
	radomDesc := makeRandomDescription(numWords)

	log.Printf("numWords is %d", numWords)
	log.Printf("randomName is %s", randomName)
	log.Printf("radomDesc is %s", radomDesc)

	body := []byte(fmt.Sprintf(`{
		"serviceName": "%s",
		"serviceDescription": "%s"
	}`, randomName, radomDesc) )

	return body
}