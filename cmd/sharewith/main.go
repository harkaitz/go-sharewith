package main
import (
	"fmt"
	"github.com/harkaitz/go-sharewith"
)

func main() {
	fmt.Println(string(sharewith.A("https://google.com", "Texto1", "hash")))
}
