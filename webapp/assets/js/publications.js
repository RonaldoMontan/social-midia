$('#new-publication').on('submit', createPublication); //# é para referenciar id
$('.like-publication').on('click', likePublication);// . é para referenciar uma classe

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

function likePublication(event) {
    console.log("curtiu a publicação ai ?")

    const elementClick = $(event.target);
    const publicationId = elementClick.closest('div').data('publication-id');//pega a classe a cima do evento clicado

    elementClick.prop('disabled', true);//desabilita o elemento até finalizar a requisição

    $.ajax({
        url: `/publication/${publicationId}/like`,
        method: "POST"
    }).done(function() {
        const contLike = elementClick.next('span');
        const quantityLike = parseInt(contLike.text());
        
        contLike.text(quantityLike +1 );

    }).fail(function() {
        alert("não deu certo");
    
    }).always(function() {
        elementClick.prop('disabled', false);// habilita o botão novamente após concluir a requisição
    })
}