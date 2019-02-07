package model

import "time"

//RUsuario struct
type RUsuario struct {
	ID         int
	Nombre     string
	Apellidos  string
	Email      string
	Usuario    string
	Contrasena string
}

//RImagen struct
type RImagen struct {
	ID    int
	URL   string
	Texto string
}

//RComentario struct
type RComentario struct {
	ID    int
	Texto string
	Fecha time.Time
}

//RLogin struct
type RLogin struct {
	Usuario    string
	Contrasena string
}
