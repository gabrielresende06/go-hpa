package main

import (
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

func Home(w http.ResponseWriter, r *http.Request)  {

	result := Raiz(0.0001, 1000000000)
	data := []byte(result)
	w.Write(data)
}

func Raiz(numero float64, interacoes int) string {
	var sum float64 = 0.0
	for i := 0; i < interacoes; i++ {
		sum += math.Sqrt(numero)
	}

	return strconv.FormatFloat(sum, 'f', 6, 64)
}