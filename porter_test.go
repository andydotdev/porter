package porter

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestConsonant(t *testing.T) {
	word := []byte("TOY")
	cmp(t, true, consonant(word, 0), "T")
	cmp(t, false, consonant(word, 1), "O")
	cmp(t, true, consonant(word, 2), "Y")
	word = []byte("SYZYGY")
	cmp(t, true, consonant(word, 0), "S")
	cmp(t, false, consonant(word, 1), "Y")
	cmp(t, true, consonant(word, 2), "Z")
	cmp(t, false, consonant(word, 3), "Y")
	cmp(t, true, consonant(word, 4), "G")
	cmp(t, false, consonant(word, 5), "Y")
	word = []byte("yoke")
	cmp(t, true, consonant(word, 0), "YOKE")
}

func TestMeasure(t *testing.T) {
	cmp(t, 0, measure([]byte("TR")))
	cmp(t, 0, measure([]byte("EE")))
	cmp(t, 0, measure([]byte("TREE")))
	cmp(t, 0, measure([]byte("Y")))
	cmp(t, 0, measure([]byte("BY")))
	cmp(t, 1, measure([]byte("TROUBLE")))
	cmp(t, 1, measure([]byte("OATS")))
	cmp(t, 1, measure([]byte("TREES")))
	cmp(t, 1, measure([]byte("IVY")))
	cmp(t, 2, measure([]byte("TROUBLES")))
	cmp(t, 2, measure([]byte("PRIVATE")))
	cmp(t, 2, measure([]byte("OATEN")))
	cmp(t, 2, measure([]byte("ORRERY")))
}

func Test1A(t *testing.T) {
	cmp(t, "caress", string(one_a([]byte("caresses"))))
	cmp(t, "poni", string(one_a([]byte("ponies"))))
	cmp(t, "ti", string(one_a([]byte("ties"))))
	cmp(t, "caress", string(one_a([]byte("caress"))))
	cmp(t, "cat", string(one_a([]byte("cats"))))
}

func Test1B(t *testing.T) {
	cmp(t, "feed", string(one_b([]byte("feed"))))
	cmp(t, "agree", string(one_b([]byte("agreed"))))
	cmp(t, "plaster", string(one_b([]byte("plastered"))))
	cmp(t, "bled", string(one_b([]byte("bled"))))
	cmp(t, "motor", string(one_b([]byte("motoring"))))
	cmp(t, "sing", string(one_b([]byte("sing"))))
	cmp(t, "motor", string(one_b([]byte("motoring"))))
	cmp(t, "conflate", string(one_b([]byte("conflated"))))
	cmp(t, "trouble", string(one_b([]byte("troubled"))))
	cmp(t, "size", string(one_b([]byte("sized"))))
	cmp(t, "hop", string(one_b([]byte("hopping"))))
	cmp(t, "tan", string(one_b([]byte("tanned"))))
	cmp(t, "fail", string(one_b([]byte("failing"))))
	cmp(t, "file", string(one_b([]byte("filing"))))
}

func Test1C(t *testing.T) {
	cmp(t, "sky", string(one_c([]byte("sky"))))
	cmp(t, "happi", string(one_c([]byte("happy"))))

}

func Test2(t *testing.T) {
	cmp(t, "relate", string(two([]byte("relational"))))
	cmp(t, "condition", string(two([]byte("conditional"))))
	cmp(t, "rational", string(two([]byte("rational"))))
	cmp(t, "valence", string(two([]byte("valenci"))))
	cmp(t, "hesitance", string(two([]byte("hesitanci"))))
	cmp(t, "digitize", string(two([]byte("digitizer"))))
	cmp(t, "conformable", string(two([]byte("conformabli"))))
	cmp(t, "radical", string(two([]byte("radicalli"))))
	cmp(t, "different", string(two([]byte("differentli"))))
	cmp(t, "vile", string(two([]byte("vileli"))))
	cmp(t, "analogous", string(two([]byte("analogousli"))))
	cmp(t, "vietnamize", string(two([]byte("vietnamization"))))
	cmp(t, "predicate", string(two([]byte("predication"))))
	cmp(t, "operate", string(two([]byte("operator"))))
	cmp(t, "feudal", string(two([]byte("feudalism"))))
	cmp(t, "decisive", string(two([]byte("decisiveness"))))
	cmp(t, "hopeful", string(two([]byte("hopefulness"))))
	cmp(t, "callous", string(two([]byte("callousness"))))
	cmp(t, "formal", string(two([]byte("formaliti"))))
	cmp(t, "sensitive", string(two([]byte("sensitiviti"))))
	cmp(t, "sensible", string(two([]byte("sensibiliti"))))
}

func Test3(t *testing.T) {
	cmp(t, "triplic", string(three([]byte("triplicate"))))
	cmp(t, "form", string(three([]byte("formative"))))
	cmp(t, "formal", string(three([]byte("formalize"))))
	cmp(t, "electric", string(three([]byte("electriciti"))))
	cmp(t, "electric", string(three([]byte("electrical"))))
	cmp(t, "hope", string(three([]byte("hopeful"))))
	cmp(t, "good", string(three([]byte("goodness"))))
}

func Test4(t *testing.T) {
	cmp(t, "reviv", string(four([]byte("revival"))))
	cmp(t, "allow", string(four([]byte("allowance"))))
	cmp(t, "infer", string(four([]byte("inference"))))
	cmp(t, "airlin", string(four([]byte("airliner"))))
	cmp(t, "gyroscop", string(four([]byte("gyroscopic"))))
	cmp(t, "adjust", string(four([]byte("adjustable"))))
	cmp(t, "defens", string(four([]byte("defensible"))))
	cmp(t, "irrit", string(four([]byte("irritant"))))
	cmp(t, "replac", string(four([]byte("replacement"))))
	cmp(t, "adjust", string(four([]byte("adjustment"))))
	cmp(t, "depend", string(four([]byte("dependent"))))
	cmp(t, "adopt", string(four([]byte("adoption"))))
	cmp(t, "homolog", string(four([]byte("homologou"))))
	cmp(t, "commun", string(four([]byte("communism"))))
	cmp(t, "activ", string(four([]byte("activate"))))
	cmp(t, "angular", string(four([]byte("angulariti"))))
	cmp(t, "homolog", string(four([]byte("homologous"))))
	cmp(t, "effect", string(four([]byte("effective"))))
	cmp(t, "bowdler", string(four([]byte("bowdlerize"))))
}

func Test5A(t *testing.T) {
	cmp(t, "probat", string(five_a([]byte("probate"))))
	cmp(t, "rate", string(five_a([]byte("rate"))))
	cmp(t, "ceas", string(five_a([]byte("cease"))))
}

func Test5B(t *testing.T) {
	cmp(t, "control", string(five_b([]byte("controll"))))
	cmp(t, "roll", string(five_b([]byte("roll"))))
}

func TestWordlist(t *testing.T) {
	f, err := os.Open("testdata/wordlist.txt")
	if err != nil {
		t.Fatal(err)
	}
	in := bufio.NewReader(f)
	f, err = os.Open("testdata/wordlist_expect.txt")
	if err != nil {
		t.Fatal(err)
	}
	out := bufio.NewReader(f)
	for word, err := in.ReadSlice('\n'); err == nil; word, err = in.ReadSlice('\n') {
		stem, err := out.ReadSlice('\n')
		if err != nil {
			panic(err)
		}
		cmp(t, strings.TrimSpace(string(stem)), string(Stem(word)), string(word))
	}
}

func BenchmarkStem(b *testing.B) {
	words, err := readLines("testdata/wordlist.txt")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	var w []byte
	for b.Loop() {
		for i := range words {
			w = Stem(words[i])
		}
	}
	_ = w
}

func cmp[T comparable](t *testing.T, want, got T, msg ...string) {
	t.Helper()
	if got != want {
		if len(msg) > 0 {
			t.Errorf("%s - want: %v, got: %v", strings.Join(msg, " "), want, got)
			return
		}
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func readLines(filename string) ([][]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	return lines, scanner.Err()
}
