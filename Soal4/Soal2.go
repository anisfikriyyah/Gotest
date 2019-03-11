package main

import "fmt"
import "net/http"
import "strings"
import "sort"


func isVowel(s string) bool {
	var vocal = []string{"a","i","u","e","o"}
	for i := 0; i < len(vocal); i++ {
		if s == vocal[i] {
		return true
		}
	}
	return false
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cases := "osama"
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

		fmt.Fprintln(w,strings.Join(v_arr,"") + strings.Join(c_arr,""))
    })

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}