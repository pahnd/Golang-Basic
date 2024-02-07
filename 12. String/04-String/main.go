package exanple

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	a := strings.Repeat("!", l)
	s := a + msg + a
	s = strings.ToUpper((s))
	fmt.Println(s)
	msgNew := `
	
	The weather looks good.
I should go and play.


	`
	fmt.Println(strings.TrimSpace(msgNew))

}
