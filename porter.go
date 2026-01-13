package porter

import "bytes"

func consonant(body []byte, offset int) bool {
	switch body[offset] {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return false
	case 'Y', 'y':
		if offset == 0 {
			return true
		}
		return offset > 0 && !consonant(body, offset-1)
	}
	return true
}

func vowel(body []byte, offset int) bool {
	return !consonant(body, offset)
}

const (
	stateVowel = iota
	stateConsonant
)

func measure(body []byte) int {
	meansure := 0
	if len(body) > 0 {
		var state int
		if vowel(body, 0) {
			state = stateVowel
		} else {
			state = stateConsonant
		}
		for i := range body {
			if vowel(body, i) && state == stateConsonant {
				state = stateVowel
			} else if consonant(body, i) && state == stateVowel {
				state = stateConsonant
				meansure++
			}
		}
	}
	return meansure
}

func hasVowel(body []byte) bool {
	for i := range body {
		if vowel(body, i) {
			return true
		}
	}
	return false
}

func one_a(body []byte) []byte {
	switch {
	case bytes.HasSuffix(body, []byte("sses")) || bytes.HasSuffix(body, []byte("ies")):
		return body[:len(body)-2]
	case bytes.HasSuffix(body, []byte("ss")):
		return body
	case bytes.HasSuffix(body, []byte("s")):
		return body[:len(body)-1]
	}
	return body
}

func star_o(body []byte) bool {
	size := len(body) - 1
	if size >= 2 && consonant(body, size-2) && vowel(body, size-1) && consonant(body, size) {
		return body[size] != 'w' && body[size] != 'x' && body[size] != 'y'
	}
	return false
}
func one_b_a(body []byte) []byte {
	size := len(body)
	switch {
	case bytes.HasSuffix(body, []byte("at")):
		fallthrough
	case bytes.HasSuffix(body, []byte("bl")):
		fallthrough
	case bytes.HasSuffix(body, []byte("iz")):
		return append(body, 'e')
	case consonant(body, size-1) && consonant(body, size-2) && body[size-1] == body[size-2]:
		if body[size-1] != 'l' && body[size-1] != 's' && body[size-1] != 'z' {
			return body[:size-1]
		}
	case star_o(body) && measure(body) == 1:
		return append(body, 'e')
	}
	return body
}

func one_b(body []byte) []byte {
	switch {
	case bytes.HasSuffix(body, []byte("eed")):
		if measure(body[:len(body)-3]) > 0 {
			return body[:len(body)-1]
		}
	case bytes.HasSuffix(body, []byte("ed")):
		if hasVowel(body[:len(body)-2]) {
			return one_b_a(body[:len(body)-2])
		}
	case bytes.HasSuffix(body, []byte("ing")):
		if hasVowel(body[:len(body)-3]) {
			return one_b_a(body[:len(body)-3])
		}
	}
	return body
}

func one_c(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("y")) && hasVowel(body[:len(body)-1]) {
		body[len(body)-1] = 'i'
	}
	return body
}

func two(body []byte) []byte {
	switch {
	case bytes.HasSuffix(body, []byte("ational")):
		if measure(body[:len(body)-7]) > 0 {
			return append(body[:len(body)-7], []byte("ate")...)
		}
	case bytes.HasSuffix(body, []byte("tional")):
		if measure(body[:len(body)-6]) > 0 {
			return body[:len(body)-2]
		}
	case bytes.HasSuffix(body, []byte("enci")), bytes.HasSuffix(body, []byte("anci")):
		if measure(body[:len(body)-4]) > 0 {
			return append(body[:len(body)-1], 'e')
		}
	case bytes.HasSuffix(body, []byte("izer")):
		if measure(body[:len(body)-4]) > 0 {
			return append(body[:len(body)-4], []byte("ize")...)
		}
	case bytes.HasSuffix(body, []byte("abli")):
		if measure(body[:len(body)-4]) > 0 {
			return append(body[:len(body)-4], []byte("able")...)
		}
	case bytes.HasSuffix(body, []byte("bli")):
		// This phrase deviates from the published algorithm.
		if measure(body[:len(body)-3]) > 0 {
			return append(body[:len(body)-1], 'e')
		}
	case bytes.HasSuffix(body, []byte("alli")):
		if measure(body[:len(body)-4]) > 0 {
			return append(body[:len(body)-4], []byte("al")...)
		}
	case bytes.HasSuffix(body, []byte("entli")):
		if measure(body[:len(body)-5]) > 0 {
			return append(body[:len(body)-5], []byte("ent")...)
		}
	case bytes.HasSuffix(body, []byte("eli")):
		if measure(body[:len(body)-3]) > 0 {
			return append(body[:len(body)-3], []byte("e")...)
		}
	case bytes.HasSuffix(body, []byte("ousli")):
		if measure(body[:len(body)-5]) > 0 {
			return append(body[:len(body)-5], []byte("ous")...)
		}
	case bytes.HasSuffix(body, []byte("ization")):
		if measure(body[:len(body)-7]) > 0 {
			return append(body[:len(body)-7], []byte("ize")...)
		}
	case bytes.HasSuffix(body, []byte("ation")):
		if measure(body[:len(body)-5]) > 0 {
			return append(body[:len(body)-5], []byte("ate")...)
		}
	case bytes.HasSuffix(body, []byte("ator")):
		if measure(body[:len(body)-4]) > 0 {
			return append(body[:len(body)-4], []byte("ate")...)
		}
	case bytes.HasSuffix(body, []byte("alism")):
		if measure(body[:len(body)-5]) > 0 {
			return append(body[:len(body)-5], []byte("al")...)
		}
	case bytes.HasSuffix(body, []byte("iveness")):
		if measure(body[:len(body)-7]) > 0 {
			return append(body[:len(body)-7], []byte("ive")...)
		}
	case bytes.HasSuffix(body, []byte("fulness")):
		if measure(body[:len(body)-7]) > 0 {
			return append(body[:len(body)-7], []byte("ful")...)
		}
	case bytes.HasSuffix(body, []byte("ousness")):
		if measure(body[:len(body)-7]) > 0 {
			return append(body[:len(body)-7], []byte("ous")...)
		}
	case bytes.HasSuffix(body, []byte("aliti")):
		if measure(body[:len(body)-5]) > 0 {
			return append(body[:len(body)-5], []byte("al")...)
		}
	case bytes.HasSuffix(body, []byte("iviti")):
		if measure(body[:len(body)-5]) > 0 {
			return append(body[:len(body)-5], []byte("ive")...)
		}
	case bytes.HasSuffix(body, []byte("biliti")):
		if measure(body[:len(body)-6]) > 0 {
			return append(body[:len(body)-6], []byte("ble")...)
		}
	case bytes.HasSuffix(body, []byte("logi")):
		// This phrase deviates from the published algorithm.
		if measure(body[:len(body)-4]) > 0 {
			return body[:len(body)-1]
		}
	}
	return body
}

func three(body []byte) []byte {
	switch {
	case bytes.HasSuffix(body, []byte("icate")):
		if measure(body[:len(body)-5]) > 0 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ative")):
		if measure(body[:len(body)-5]) > 0 {
			return body[:len(body)-5]
		}
	case bytes.HasSuffix(body, []byte("alize")):
		if measure(body[:len(body)-5]) > 0 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("iciti")):
		if measure(body[:len(body)-5]) > 0 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ical")):
		if measure(body[:len(body)-4]) > 0 {
			return body[:len(body)-2]
		}
	case bytes.HasSuffix(body, []byte("ful")):
		if measure(body[:len(body)-3]) > 0 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ness")):
		if measure(body[:len(body)-4]) > 0 {
			return body[:len(body)-4]
		}
	}
	return body
}

func four(body []byte) []byte {
	switch {
	case bytes.HasSuffix(body, []byte("al")):
		if measure(body[:len(body)-2]) > 1 {
			return body[:len(body)-2]
		}
	case bytes.HasSuffix(body, []byte("ance")):
		if measure(body[:len(body)-4]) > 1 {
			return body[:len(body)-4]
		}
	case bytes.HasSuffix(body, []byte("ence")):
		if measure(body[:len(body)-4]) > 1 {
			return body[:len(body)-4]
		}
	case bytes.HasSuffix(body, []byte("er")):
		if measure(body[:len(body)-2]) > 1 {
			return body[:len(body)-2]
		}
	case bytes.HasSuffix(body, []byte("ic")):
		if measure(body[:len(body)-2]) > 1 {
			return body[:len(body)-2]
		}
	case bytes.HasSuffix(body, []byte("able")):
		if measure(body[:len(body)-4]) > 1 {
			return body[:len(body)-4]
		}
	case bytes.HasSuffix(body, []byte("ible")):
		if measure(body[:len(body)-4]) > 1 {
			return body[:len(body)-4]
		}
	case bytes.HasSuffix(body, []byte("ant")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ement")):
		if measure(body[:len(body)-5]) > 1 {
			return body[:len(body)-5]
		}
	case bytes.HasSuffix(body, []byte("ment")):
		if measure(body[:len(body)-4]) > 1 {
			return body[:len(body)-4]
		}
	case bytes.HasSuffix(body, []byte("ent")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ion")):
		if measure(body[:len(body)-3]) > 1 {
			if len(body) > 4 && (body[len(body)-4] == 's' || body[len(body)-4] == 't') {
				return body[:len(body)-3]
			}
		}
	case bytes.HasSuffix(body, []byte("ou")):
		if measure(body[:len(body)-2]) > 1 {
			return body[:len(body)-2]
		}
	case bytes.HasSuffix(body, []byte("ism")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ate")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("iti")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ous")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ive")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	case bytes.HasSuffix(body, []byte("ize")):
		if measure(body[:len(body)-3]) > 1 {
			return body[:len(body)-3]
		}
	}
	return body
}

func five_a(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("e")) && measure(body[:len(body)-1]) > 1 {
		return body[:len(body)-1]
	} else if bytes.HasSuffix(body, []byte("e")) && measure(body[:len(body)-1]) == 1 && !star_o(body[:len(body)-1]) {
		return body[:len(body)-1]
	}
	return body
}

func five_b(body []byte) []byte {
	size := len(body)
	if measure(body) > 1 && consonant(body, size-1) && consonant(body, size-2) && body[size-1] == body[size-2] && body[size-1] == 'l' {
		return body[:len(body)-1]
	}
	return body
}

func Stem(word []byte) []byte {
	word = bytes.TrimSpace(bytes.ToLower(word))

	if len(word) > 2 {
		return five_b(five_a(four(three(two(one_c(one_b(one_a(word))))))))
	}
	return word
}

func StemStr(word string) string {
	return string(Stem([]byte(word)))
}
