package model

//Usuario struct
type Usuario struct {
	Nombre     string
	Apellidos  string
	Usuario    string
	Email      string
	Contrasena string
}

//Imagen struct
type Imagen struct {
	URL   string
	Texto string
}

//Comentario struct
type Comentario struct {
	Texto string
}

//Login struct
type Login struct {
	Usuario    string
	Contrasena string
}
