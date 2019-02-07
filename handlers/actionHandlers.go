package handlers

import (
	"encoding/json"
	"fmt"
	client "instagram/data/dataclient"
	"instagram/data/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/securecookie"
)

//Insert Función que inserta una petición en la base de datos local
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathInsert {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var usuario model.Usuario
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		if usuario.Nombre == "" || usuario.Apellidos == "" || usuario.Usuario == "" || usuario.Email == "" || usuario.Contrasena == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición está vacía")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		hashComoCadena := string(hash)
		usuario.Contrasena = hashComoCadena

		resp := client.InsertarUsuario(&usuario)

		fmt.Fprint(w, resp)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//LoginUsuario funcion
func LoginUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathLoginUsuario {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)
	resp := false
	if e == nil {
		var usuario model.Login
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		if usuario.Usuario == "" || usuario.Contrasena == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición está vacía")
			return
		}

		contrasenaBD := client.Login(&usuario)
		if err := bcrypt.CompareHashAndPassword([]byte(contrasenaBD), []byte(usuario.Contrasena)); err != nil {
			fmt.Printf("No estas logeado")
			fmt.Println(usuario.Contrasena)
		} else {
			resp = true
			setSession(usuario.Usuario, w)
			fmt.Printf("Usuario logeado")

		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, resp)
	}
	fmt.Fprintln(w, resp)
}

//Logout funcion
func Logout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

//Cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//InsertImg Función que inserta una petición en la base de datos local
func InsertImg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathInsertImg {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseMultipartForm(2000)
	//Coger el archivo y meterlo en una variable
	file, fileInto, err := r.FormFile("archivo")
	//Coger el texto del formulario y merterlo en una variable
	texto := r.FormValue("texto")
	usuario := getUserName(r)
	fmt.Println(texto, "Usuario: ", usuario)
	f, err := os.OpenFile("./files/"+fileInto.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//La linea de abajo que esta comentada me manda a la página donde está el nombre del archivo
	//fmt.Fprintf(w, fileInto.Filename)
	//Esta linea de aqui abajo me manda a la pagina principal donde están todas las fotos
	http.Redirect(w, r, "/", 301)
	//Datos de la base de datos
	id := client.ConsultaID(usuario)
	fmt.Println(id)
	//Subir foto a la base de datos
	go client.InsertarImagen(fileInto.Filename, texto, id)
}

//Listado Función que devuelve las peticiones de la base de datos dado un filtro
func Listado(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListado {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	lista := client.Listar()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	respuesta, _ := json.Marshal(&lista)
	fmt.Fprint(w, string(respuesta))
}
