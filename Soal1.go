// input: "omama"   hasil: "aaomm" ('a', dan 'o' sebagai huruf hidup dan 'm' sebagai huruf mati)
// input: "osama"   hasil: "aaoms"

package main
import "fmt"
import "strings"
import "sort"

func main() {
	sortWord("omama")
	sortWord("osama")
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

func sortWord(str string) {
	cases := str
	var v string
	var c string
	for i:= 0; i < len(cases); i++ {
		if isVowel(string(cases[i])) {
			v = v + string(cases[i])
		} else {
			c = c + string(cases[i])
		}
	}

	// lakukan split biar bisa diurutkan
	v_arr := strings.SplitAfter(v, "")
	c_arr := strings.SplitAfter(c, "")
	sort.Strings(v_arr)
	sort.Strings(c_arr)

	fmt.Println(strings.Join(v_arr,"") + strings.Join(c_arr,""))
}