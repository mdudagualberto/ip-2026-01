package handlers

import (
	"servidorHTTP/app/utils"
	"fmt"
	"net/http"
)

func ListarPacientesHandler(w http.ResponseWriter, r *http.Request) {
	db := utils.ConectarDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, nome, cpf, data_nascimento, diagnostico FROM pacientes")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "<h1>Pacientes Cadastrados</h1><ul>")

	for rows.Next() {
		var p utils.Paciente
		rows.Scan(&p.ID, &p.Nome, &p.CPF, &p.DataNascimento, &p.Diagnostico)

		fmt.Fprintf(w, "<li><b>%s</b> | CPF: %s | Diagnóstico: %s</li>",
			p.Nome, p.CPF, p.Diagnostico)
	}

	fmt.Fprintf(w, "</ul>")
}