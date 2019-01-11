package bcryptor

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print("input > ")
		bs, err := bcrypt.GenerateFromPassword([]byte(<-ch), bcrypt.DefaultCost)
		if err != nil {
			fmt.Fprintln(os.Stderr, "err:", err)
		}
		fmt.Fprintln(os.Stdout, string(bs))
	}
}
