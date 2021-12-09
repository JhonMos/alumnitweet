package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JhonMos/alumnitweet/middlew"
	"github.com/JhonMos/alumnitweet/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al Server*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoDb(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDb(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoDb(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoDb(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoDb(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoDb(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoDb(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoDb(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoDb(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoDb(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoDb(routers.ObtenerBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
