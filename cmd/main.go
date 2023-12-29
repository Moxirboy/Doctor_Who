package main

import (
	"net/http"
)

func main() {
	
	//getProductsHandler := http.HandlerFunc(getProducts)
	//http.Handle("/products", getProductsHandler)
	//http.ListenAndServe(":5005", nil)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world")
}
