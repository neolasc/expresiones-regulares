package main
import (
	"fmt"
	"regexp"
)
func prettyMatches(m []string) string {
	s := "["
	for i, e := range m {
		s += e
		if i < len(m)-1 {
			s += "|"
		}
	}
	s += "]"
	return s
}

func prettySubmatches(m [][]string) string {
	s := "[\n"
	for _, e := range m {
		s += "    " + prettyMatches(e) + "\n"
	}
	s += "]"
	return s
}

var (exps = []string{"if", "while", ">","=","else","try","catch","[a-z]","[A-Z]"}

	text = `while ( x > 0) else catch try if D`
)
func main() {



	for _, e := range exps {
		re := regexp.MustCompile(e)
		fmt.Println(e + ":")
		fmt.Println("1. encontrar String: ", re.FindString(text))
		fmt.Println("2. Buscar índice de String: ", re.FindStringIndex(text))
		fmt.Printf("3. buscar todos los Strings: %v\n", prettyMatches(re.FindAllString(text, -1)))
		fmt.Printf("4. buscar todos los indices de String: %v\n", re.FindAllStringIndex(text, -1))
		
	}

fmt.Println(r.MatchString("while"))
}
