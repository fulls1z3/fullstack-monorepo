package hello

import "fmt"

func GetHello(s string) string {
	return fmt.Sprintf("Hello World from %s", s)
}
