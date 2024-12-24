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
        URL: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            nikc: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        }
    })
}