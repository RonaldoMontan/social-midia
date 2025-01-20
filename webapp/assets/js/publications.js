$('#new-publication').on('submit', createPublication); //# é para referenciar id

$(document).on('click', '.like-publication', likePublication);// . é para referenciar uma classe
$(document).on('click', '.unlike-publication', unlikePublication);
console.log($('.unlike-publication').length); // Verifica se os elementos foram inseridos

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
    console.log("Evento: likePublication disparado");

    event.preventDefault();
    const elementClick = $(event.target);
    const publicationId = elementClick.closest('div').data('publication-id');

    elementClick.prop('disabled', true);

    $.ajax({
        url: `/publication/${publicationId}/like`,
        method: "POST"
    }).done(function() {
        const contLike = elementClick.next('span');
        const quantityLike = parseInt(contLike.text());

        contLike.text(quantityLike + 1);

        // Modificar as classes do elemento para descurtir
        elementClick.addClass('unlike-publication');
        elementClick.css('color', '#915cc2');
        elementClick.removeClass('like-publication');

    }).fail(function() {
        alert("não deu certo");

    }).always(function() {
        elementClick.prop('disabled', false);
    });
}

function unlikePublication(event) {
    console.log("Evento: unlikePublication disparado");

    event.preventDefault();
    const elementClick = $(event.target);
    const publicationId = elementClick.closest('div').data('publication-id');

    elementClick.prop('disabled', true);

    $.ajax({
        url: `/publication/${publicationId}/unlike`,
        method: "POST"

    }).done(function() {

        console.log("dentro do done")
        const contLike = elementClick.next('span');
        const quantityLike = parseInt(contLike.text());

        contLike.text(quantityLike - 1);

        // Modificar as classes do elemento para curtir
        elementClick.removeClass('unlike-publication');
        elementClick.css('color', 'black');
        elementClick.addClass('like-publication');

    }).fail(function() {
        console.log('caiu no fail - unlike')

    }).always(function() {
        elementClick.prop('disabled', false);
    });
}
