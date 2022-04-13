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
  }).done(function() {
    window.location = "/home";
}).fail(function() {
  Swal.fire("Ops...", "Usuários ou senhas incorretos!", "error");
});
    
}