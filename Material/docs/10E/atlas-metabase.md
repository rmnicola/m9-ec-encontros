---
title: Integração MongoDB Atlas
sidebar_position: 1
sidebar_class_name: autoestudo
slug: /atlas-metabase
---

:::tip

Essa é uma versão inicial dessa seção. Vou revisar em breve com imagens para os
guias.

:::

# Integrando o mongoDB Atlas com o Metabase

:::warning

Para conseguir fazer a integração descrita nessa seção, você deve:

* Ter uma conta no MongoDB Atlas
* Ter uma instância do Metabase rodando (local ou em nuvem)

:::

## 1. Configurando o MongoDB Atlas

1. **Crie um Cluster**: Faça login na sua conta do MongoDB Atlas e crie um novo
   cluster se ainda não tiver um.
2. **Coloque seu IP na Whitelist**: Na seção "Network Access" do Atlas,
   adicione o endereço IP do servidor onde o Metabase está hospedado à lista de
   IPs permitidos.
3. **Crie um usuário do banco de dados**: Na seção "Database Access", crie um
   novo usuário com permissões de leitura e escrita no banco de dados que você
   deseja conectar.

## 2. Configurando o Metabase

### 2.1. Configuração convencional

1. Faça login no Metabase e vá para "Admin" > "Databases" > "Add database".
2. Escolha "MongoDB" como o tipo de banco de dados.
3. Preencha os campos com as seguintes informações:
   - **Name**: Dê um nome para a sua conexão.
   - **Host**: Cole a URI para conexão com seu banco de dados. Ela vai ter um
     formato similar a: **nome-do-cluster.seu-id.mongodb.net**. Para
     encontrar seu URI vá para a seção de clusters do Atlas, clique em
     **Connect**, escolha **Drivers** e veja a URI utilizada na string de
     conexão fornecida.
   - **Database name**: O nome do banco de dados que você deseja conectar.
   - **Port**: Tipicamente a porta usada é a 27017
   - **Usuário**: O seu usuário cadastrado no Atlas
   - **Senha**: A senha do seu usuário cadastrado no Atlas
   - **Configure o metabase para usar DNS SRV**: Em **opções avançadas**
     assinale para o metabase usar o DNS SRV.
4. Clique em "Save" para salvar a configuração.

### 2.2. Configuração usando string de conexão do MongoDB Atlas

:::warning

Esse modo de configuração ainda não funcionou para mim. Vou continuar tentando,
mas se não der certo vou só deletar essa subseção.

:::

1. No MongoDB Atlas, vá para a seção "Clusters" e clique no botão "Connect" do
   seu cluster.
2. Escolha "Connect your application" como método de conexão.
3. Selecione o driver correto (geralmente "Java") e copie a string de conexão
   fornecida. Ela deve ter o formato:
   ```
   mongodb+srv://<username>:<password>@<cluster-url>/<dbname>?retryWrites=true&w=majority
   ```
4. Faça login no Metabase e vá para "Admin" > "Databases" > "Add database".
5. Escolha "MongoDB" como o tipo de banco de dados.
6. Escolha a opção **colar uma string de conexão** e cole a sua string que
   extraiu do Atlas.
7. Clique em "Save" para salvar a configuração.
