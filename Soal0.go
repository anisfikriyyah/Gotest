// input: "omama"   hasil: "huruf mati: 2, huruf hidup: 2"  (huruf hidup yaitu o dan a. a tidak perlu muncul 2x)

package main
import "fmt"
import "strings"
import "sort"

func main() {
	vowelConsonan("omama")
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

func vowelConsonan(str string) {
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
	sort.Strings(v_arr)

	// looping utk ngilangin data yang sama
	var ilangSama = []string{string(v_arr[0])}
	for i:= 1; i < len(v_arr); i++ {
		if string(v_arr[i]) != string(v_arr[i-1]) {
			ilangSama = append(ilangSama, string(v_arr[i]))
		}
	}

	fmt.Println("huruf mati: ", len(c), ",huruf hidup: ", len(ilangSama) )
}