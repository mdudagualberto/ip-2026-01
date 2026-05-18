package handlers

import (
	"servidorHTTP/app/utils"
	"fmt"
	"net/http"
)

func DeletarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	cpf := r.FormValue("cpf")

	db := utils.ConectarDB()
	defer db.Close()

	result, err := db.Exec("DELETE FROM pacientes WHERE cpf=$1", cpf)
	if err != nil {
		http.Error(w, "Erro ao deletar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	linhas, _ := result.RowsAffected()
	if linhas == 0 {
		fmt.Fprintf(w, "Nenhum paciente encontrado com CPF %s", cpf)
		return
	}

	fmt.Fprintf(w, "Paciente deletado com sucesso!")
}
