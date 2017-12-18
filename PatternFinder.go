package main

import (
	"regexp"
	"strings"
)

type StringFinder struct {
	haystack string
	needle   string
}

func normaliseString(needle string) string {
	// Normalise all football club representation
	commonPattern := []string{"FC", "F.C.", "F.C"}
	workerString := needle
	for _, val := range commonPattern {
		workerString = strings.Replace(workerString, val, "", -1)
	}
	return workerString
}

func regexMatchForShortForms(haystack string, needle string) bool {
	re := regexp.MustCompile("\\b[A-Z]")
	if re != nil {
		shortForm := re.FindAllString(haystack, -1)
		if strings.ToLower(strings.Replace(needle, " ", "", -1)) == strings.ToLower(strings.Join(shortForm, "")) {
			return true
		} else {
			shortForm = re.FindAllString(haystack, 2) //Atleast first 2 characters match in short form
			if strings.ContainsAny(haystack, needle) {
				if strings.HasPrefix(strings.ToLower(needle), strings.ToLower(strings.Join(shortForm, ""))) {
					return true
				}
			}
		}
	}
	return false
}

func (sf *StringFinder) MatchString() int {
	//Handle empty needle
	if len(strings.Replace(sf.needle, " ", "", -1)) == 0 {
		return 0
	}
	switch {
	case strings.Contains(strings.ToLower(sf.haystack), strings.ToLower(sf.needle)):
		if len(strings.Split(sf.needle, " ")) > 1 { //Normally more than one word matches will be a hit
			return 1
		}
	case strings.Contains(strings.ToLower(sf.haystack), strings.ToLower(normaliseString(sf.needle))):
		return 1
	case regexMatchForShortForms(sf.haystack, sf.needle):
		return 1
	}
	return 0
}
func main() {

	//fmt.Println("Testing pattern match")
}
