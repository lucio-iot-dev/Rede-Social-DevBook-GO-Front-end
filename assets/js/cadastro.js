$(`#formulario-cadastro`).on(`submit`, criarUsuario);


function criarUsuario(evento) {
  evento.preventDefault(); //comando para no console log  ao apertar f12 para ver a mensagem no console ,ela não sumir rapido e não atualizar a pagina
   console.log("Dentro da função usuário!");

   if ($(`#senha`).val() != $(`#confirmar-senha`).val()) {
       
     alert("As senhas não coicidem!");
     return;
   }

   $.ajax({
     url: "/usuarios",
     method: "POST",
     data: {
       nome: $(`#nome`).val(),
       email: $(`#email`).val(),
       nick: $(`#nick`).val(),
       senha: $(`#senha`).val()

     }
   }).done(function() { // Ok 201  20  204
       alert("Usuário cadastrado com sucesso!");
   }).fail(function(erro) {  // Erros 400  404  401  403  500
      console.log(erro);
      alert("Erro ao cadastrar o usuário!");
   });

}