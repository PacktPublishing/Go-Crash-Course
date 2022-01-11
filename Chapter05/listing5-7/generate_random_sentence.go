package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const (
	adjectivesFileName = "adjectives.txt"
	nounsFileName      = "nouns.txt"
	verbsFileName      = "verbs.txt"
	adverbsFileName    = "adverbs.txt"
)

type Phrases []string

func (p Phrases) GetRandomPhrase() string {
	rand.Seed(time.Now().Unix()) // so that we always get different random values
	// chooses a random number from 0-(len(p) - 1) and thus a random value from the phrases slice
	return p[rand.Intn(len(p))]
}

func NewPhrases(fileName string) (Phrases, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		// return nil slice and an error with proper error message
		return nil, fmt.Errorf("Error in reading phrases from file %s: %v", fileName, err)
	}
	dataString := string(data)
	// we split on newline since every phrase is in its own line
	return strings.Split(dataString, "\n"), nil
}

type Sentence struct {
	Adjective, Noun, Verb, Adverb string
}

func (s Sentence) Display() {
	fmt.Printf("The %s %s %s %s.\n", s.Adjective, s.Noun, s.Verb, s.Adverb)
}

func NewSentence(adjectivePhrase, nounPhrase, verbPhrase, adverbPhrase string) Sentence {
	return Sentence{
		Adjective: adjectivePhrase,
		Noun:      nounPhrase,
		Verb:      verbPhrase,
		Adverb:    adverbPhrase,
	}
}

func main() {
	adjectives, errAdj := NewPhrases(adjectivesFileName)
	if errAdj != nil {
		// errAdj.Error() gives string value of errAdj
		fmt.Printf("Unable to read from adjectives file: %s\n", errAdj.Error())
		return
	}

	nouns, errNouns := NewPhrases(nounsFileName)
	if errNouns != nil {
		fmt.Printf("Unable to read from nouns file: %s\n", errNouns.Error())
		return
	}

	verbs, errVerbs := NewPhrases(verbsFileName)
	if errVerbs != nil {
		fmt.Printf("Unable to read from verbs file: %s\n", errVerbs.Error())
		return
	}

	adverbs, errAdv := NewPhrases(adverbsFileName)
	if errAdv != nil {
		fmt.Printf("Unable to read from adverbs file: %s\n", errAdv.Error())
		return
	}

	sentence := NewSentence(adjectives.GetRandomPhrase(), nouns.GetRandomPhrase(), verbs.GetRandomPhrase(), adverbs.GetRandomPhrase())
	sentence.Display()
}
