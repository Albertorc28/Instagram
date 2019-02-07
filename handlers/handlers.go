package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de scripts de CSS
const PathCSSFiles string = "/css/"

//PathRegisterFile Ruta de envío de peticiones
const PathRegisterFile string = "/register"

//PathInsert Ruta de envío de peticiones
const PathInsert string = "/insert"

//PathLoginFile Ruta de envío de peticiones
const PathLoginFile string = "/login"

//PathLoginUsuario Ruta
const PathLoginUsuario string = "/loginusuario"

//PathLogout Ruta
const PathLogout string = "/logout"

//PathInsertImg Ruta de envío de peticiones
const PathInsertImg string = "/insertImg"

//PathInsertImgFile Ruta de envío de peticiones
const PathInsertImgFile string = "/imgfile"

//PathListado Ruta de obtención de las idiomas de hoy
const PathListado string = "/listado"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Manejadores es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathCSSFiles] = CSSFile
	Manejadores[PathRegisterFile] = RegisterFile
	Manejadores[PathInsert] = Insert
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathLoginUsuario] = LoginUsuario
	Manejadores[PathLogout] = Logout
	Manejadores[PathInsertImg] = InsertImg
	Manejadores[PathInsertImgFile] = InsertImgFile
	Manejadores[PathListado] = Listado
}
