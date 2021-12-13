package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JhonMos/alumnitweet/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagtemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro paǵina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagtemp)

	results, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los archivos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
