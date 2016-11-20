package main

import (
    "github.com/RadhiFadlillah/go-sastrawi"
	"io/ioutil"
	"fmt"
	"regexp"
)

func NewContext(filename string, translations []string) Context  {
    return Context {
        Filename: filename,
        Translations: translations,
        output: make(map[string]string),
    }
}

func (c *Context) Translation() map[string]string {
    for _, trans := range c.Translations {
        c.checkTranslation(trans)
    }
    return c.output
}

func ParseTranslation(file string) (output []string) {
    reg := regexp.MustCompile("msgstr \"(.+)\"")
	for _, t := range reg.FindAllStringSubmatch(string(file), -1) {
		output = append(output, t[1])
	}
	return    
}

func CheckGetText(filename string) {
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    translation := ParseTranslation(string(file))
    context := NewContext(filename, translation)
    errors := context.Translation()
    for k, v := range errors {
        fmt.Println(k + " => " + v)
    }
}

func (c *Context) checkTranslation(translation string) {
    tokenizer := sastrawi.NewTokenizer()
    kamus := sastrawi.DefaultDictionary
    stemmer := sastrawi.NewStemmer(kamus)
    
    token := tokenizer.Tokenize(translation)
    for _, t := range token {
        if kamus.Find(t) || len(t) <= 3 {
            continue
        }
        if stemmer.Stem(t) == t {
            c.output[translation] = t + " bukan merupakan kata baku!"
            break
        }
    }
}