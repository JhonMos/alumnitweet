package models

/* RespuestaLogin tiene el token que se devuelve de login */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
