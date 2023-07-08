package main

import "net/http"

func main() {
	http.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
