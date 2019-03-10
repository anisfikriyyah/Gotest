// input: "omama"   hasil: "huruf mati: 2, huruf hidup: 2"  (huruf hidup yaitu o dan a. a tidak perlu muncul 2x)

package main
import "fmt"

func main() {
	fmt.Println(isVowel("a"))
}

func isVowel(s string) bool {
	var vocal = []string{"a","i","u","e","o"}
	for i := 0; i < len(vocal); i++ {
	  if s == vocal[i] {
		return true
	  }
	}
	return false
}