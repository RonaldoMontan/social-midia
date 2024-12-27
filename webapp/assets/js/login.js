$('#login').on('submit', login);

//Obter valores do front e devolver resposta interativa.
function login(event) {

    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(function() {
        alert("Welcome !");
        window.location= "/home";
    }).fail(function(erro){
        console.log(erro)
        alert("Verify your credentials !")
    });
}