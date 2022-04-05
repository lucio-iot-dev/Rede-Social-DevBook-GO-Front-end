$('#nova-publicacao').on('submit', criarPublicacao) // # é para id
$('.curtir-publicacao').on('click', curtirPublicacao) // . é para classe


function criarPublicacao(evento) {
  evento.preventDefault()

  $.ajax({
    url: '/publicacoes',
    method: 'POST',
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    }
  })
    .done(function () {
      window.location = '/home'
    })
    .fail(function () {
      alert('Erro ao criar a publicação')
    })
}

function curtirPublicacao(evento) {
  evento.preventDefault();

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  console.log(publicacaoId);
}
