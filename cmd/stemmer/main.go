package main

import (
	"bufio"
	"fmt"
	"os"

	"andy.dev/porter"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(porter.Stem(s.Bytes()))
	}
}
