$(`#formulario-cadastro`).on(`submit`, criarUsuario);


function criarUsuario(evento) {
  evento.preventDefault(); //comando para no console log  ao apertar f12 para ver a mensagem no console ,ela não sumir rapido e não atualizar a pagina
   console.log("Dentro da função usuário!");

   if (($(`#senha`).val()) != ($(`#confirmar-senha`).val())) {
       
     alert("As senhas não coicidem!");
     return;
   }

   $.ajax({
     url: "/usuarios",
     methodo: "POST",
     data: {
       nome: $(`#nome`).val(),
       email: $(`#email`).val(),
       nick: $(`#nick`).val(),
       senha: $(`#senha`).val()

     }
   });

}