package main

import (
	"net/http"
	"aprendizadoalura/ApWeb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}