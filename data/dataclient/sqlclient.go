package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarUsuario test
func InsertarUsuario(objeto *model.Usuario) bool {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	resp := false

	comando := "SELECT ID FROM Usuario WHERE (Usuario = '" + objeto.Usuario + "' OR Email = '" + objeto.Email + "') LIMIT 1"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Usuario WHERE (Usuario = ? OR Email = ?) LIMIT 1", objeto.Usuario, objeto.Email)

	var resultado string

	for query.Next() {
		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}

	if resultado == "" {
		fmt.Println("nombre: ", objeto.Nombre)
		defer db.Close()
		insert, err := db.Query("INSERT INTO Usuario (Nombre, Apellidos, Usuario, Email, Contrasena) VALUES (?, ?, ?, ?, ?)", objeto.Nombre, objeto.Apellidos, objeto.Usuario, objeto.Email, objeto.Contrasena)

		if err != nil {
			panic(err.Error())
		}
		insert.Close()
		resp = true
	}

	return resp
}

//Login funcion
func Login(objeto *model.Login) string {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT Contrasena FROM Usuario WHERE (Usuario = '" + objeto.Usuario + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT Contrasena FROM Usuario WHERE (Usuario = '" + objeto.Usuario + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var resultado string
	for query.Next() {

		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//ConsultaID test
func ConsultaID(usuario string) int {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT ID FROM Usuario WHERE (Usuario = '" + usuario + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Usuario WHERE (Usuario = '" + usuario + "')")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var resultado int
	for query.Next() {
		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//InsertarImagen test
func InsertarImagen(url string, texto string, id int) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO Imagen(URL, Texto, Usuario_ID) VALUES (?, ?, ?)", url, texto, id)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//Listar test
func Listar() []model.RImagen {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT ID, URL, Texto FROM Imagen"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID, URL, Texto FROM Imagen")
	if err != nil {
		panic(err.Error())
	}
	resultado := make([]model.RImagen, 0)
	for query.Next() {
		var imagen = model.RImagen{}
		err = query.Scan(&imagen.ID, &imagen.URL, &imagen.Texto)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, imagen)
	}
	return resultado
}
