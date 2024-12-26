$('#form-register').on('submit', createUser);


function createUser(event) {

    event.preventDefault();
    console.log("ronaldo ok");

    const password = $('#password').val();
    const confirm_password = $('#confirm-password').val();

    if (password !== confirm_password) {

        alert("The passwords do not match!");
        console.log("senha não é igual");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(function() {
        alert("User successfully registred");
    }).fail(function(erro) {
        console.log(erro);
        alert("Error while register user");
    });
}