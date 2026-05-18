package handlers

import (
	"servidorHTTP/app/utils"
	"fmt"
	"net/http"
)

func AtualizarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	cpf := r.FormValue("cpf")
	novoNome := r.FormValue("nome")
	novoDiag := r.FormValue("diagnostico")

	db := utils.ConectarDB()
	defer db.Close()

	result, err := db.Exec(
		"UPDATE pacientes SET nome=$1, diagnostico=$2 WHERE cpf=$3",
		novoNome, novoDiag, cpf,
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	linhas, _ := result.RowsAffected()
	if linhas == 0 {
		fmt.Fprintf(w, "Nenhum paciente encontrado com CPF %s", cpf)
		return
	}

	fmt.Fprintf(w, "Paciente atualizado com sucesso!")
}
