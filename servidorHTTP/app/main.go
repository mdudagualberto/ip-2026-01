package main

import (
	"servidorHTTP/app/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/paciente/criar", handlers.CriarPacienteHandler)
	http.HandleFunc("/paciente/listar", handlers.ListarPacientesHandler)
	http.HandleFunc("/paciente/atualizar", handlers.AtualizarPacienteHandler)
	http.HandleFunc("/paciente/deletar", handlers.DeletarPacienteHandler)

	porta := "8080"

	fmt.Println("✅ Servidor rodando em http://localhost:" + porta)

	log.Fatal(http.ListenAndServe(":"+porta, nil))
}