$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
  evento.preventDefault(); //Para prevenir o comportamento padrão do formulário

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      senha: $('#senha').val(),
    }
  }).done(function() { // Ok 201  20  204
    window.location = "/home";
}).fail(function(erro) {  // Erros 400  404  401  403  500
   console.log(erro);
   alert("Usuário ou senha inválidos!");
});
    
}