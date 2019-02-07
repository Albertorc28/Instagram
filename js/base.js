$(document).ready(function() {
    ActualizarHistorial();
    console.log("funciona js");
    var formularioRegistro = $("#formularioRegistro input:last-child");
    console.log(document.cookie);
    console.log(formularioRegistro);

    $(formularioRegistro).click(function() {
        var palabra = $("#nombre").val();
        var palabras = $("#apellidos").val();
        var user = $("#usuario").val();
        var correo = $("#email").val();
        var password = $("#contrasena").val();
        console.log(palabra, palabras, user, correo, password);
        
        var envio = {
            nombre: palabra,
            apellidos: palabras,
            usuario: user,
            email: correo,
            contrasena: password
        };
        console.log(envio);
        
        $.post({
            url:"/insert",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
            },      
            dataType: "json"

        }).done(function(data) {
            console.log(data);
           if (data == true){
               window.location.href = "/login";
           } else {
               alert("El usuario ya existe.");
           }

        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });
    $("#login").click(function() {
        var user = $("#usuariologin").val();
        var password = $("#contrasenalogin").val();
        console.log(user, password);
        
        var envio = {
            usuario: user,
            contrasena: password
        };
        console.log(envio);
        
        $.post({
            url:"/loginusuario",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
            },      
            dataType: "json"

        }).done(function(data) {
            if (data == true){
                window.location.href = "/";
            }
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });

    if (document.cookie == ""){
        $("#logeado").hide();
        
    }else{
        $("#no_logeado").hide();
    }
  
    $("#fotosubir").click(function() {   
        location.href="imgfile";  
    });
    
    if (document.cookie == ""){
        $("#fotosubir").hide();      
    }
});
function ActualizarHistorial() {   
    $.ajax({
        url: "/listado",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            if(data != null)
                console.log(data.length + " objetos obtenidos");
            Historial_Imagen(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}
function Historial_Imagen(array) {
    var div = $("#imagenes .col #listado");
    div.children().remove();
    if(array != null && array.length > 0) {
        for(var x = 0; x < array.length; x++){
            $("#imagenes .col #listado").append(
                "<img src='/files/"+array[x].URL+"' width='200px'>"+
                "<p>"+array[x].Texto+"</p>"
            );
        }
    } else {
        div.append('<tr><td colspan="3">No hay imagenes</td></tr>');     
    }    
}


