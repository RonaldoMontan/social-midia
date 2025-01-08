$('#new-publication').on('submit', createPublication);

function createPublication(event) {
    
    event.preventDefault();

    $.ajax({
        url: "/publication",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(erro){
        console.log(erro);
        alert("Error");
    });
}