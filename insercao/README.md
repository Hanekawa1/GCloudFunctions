# AutomovelGoogleCloud

Projeto de uma function com datastore do google cloud.

Essa function foi desenvolvida utilizando Golang, versão 1.13, e os pacotes do Firestore para realizar a conexão e configuração de acesso e dados ao banco.

A function para listagem dos dados foi feita em Node.js, versão 10, podendo ser encontrada [neste repositório.](https://gitlab.com/edg-dev/listarautomoveisfunction)

O endpoint da function de inserção: https://us-central1-automovel-291618.cloudfunctions.net/InsereDadosAutomovel?placa=abc1234&cor=azul&preco=40000&modelo=honda&marca=civic

Imagem de chamada do endpoint de inserção:

![Imagem de chamada da function de inserção](https://gitlab.com/edg-dev/automovelgooglecloud/-/raw/master/images/function_call_insert.png)

O endpoint da function de listagem (considerando parâmetro da placa passado na url [GET]): https://us-central1-automovel-291618.cloudfunctions.net/getAutomoveis?placa=abc1234 

Imagem de chamada do endpoint de listagem: 

![Imagem de chamada da function de listagem](https://gitlab.com/edg-dev/listarautomoveisfunction/-/raw/master/images/function_call_list.png)
