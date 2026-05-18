package handlers

import (
	"servidorHTTP/app/utils"
	"fmt"
	"net/http"
)

func CriarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	dataNasc := r.FormValue("data_nascimento")
	diagnostico := r.FormValue("diagnostico")

	db := utils.ConectarDB()
	defer db.Close() // fecha a conexão quando a função terminar

	_, err := db.Exec(
		"INSERT INTO pacientes (nome, cpf, data_nascimento, diagnostico) VALUES ($1, $2, $3, $4)",
		nome, cpf, dataNasc, diagnostico,
	)
	if err != nil {
		http.Error(w, "Erro ao cadastrar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Paciente %s cadastrado com sucesso!", nome)
}