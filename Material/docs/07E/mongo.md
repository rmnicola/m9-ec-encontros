---
title: MongoDB
sidebar_position: 2
sidebar_class_name: autoestudo
slug: /mongo
---

import Admonition from '@theme/Admonition';

# Introdução ao MongoDB

MongoDB é um sistema de banco de dados NoSQL orientado a documentos, que
armazena dados em um formato semelhante ao JSON, chamado BSON. É conhecido por
sua flexibilidade e escalabilidade, tornando-o uma escolha popular para
aplicativos modernos e em tempo real.

## 1. Elementos do MongoDB
- **Database**: É o contêiner de alto nível que armazena coleções de
  documentos. Um único servidor MongoDB pode conter vários bancos de dados.
- **Collections**: São agrupamentos de documentos, semelhantes às tabelas em
  bancos de dados relacionais. Uma coleção pode conter documentos com esquemas
  diferentes.
- **Documentos**: São registros individuais armazenados em uma coleção,
  representados em formato BSON. Cada documento possui um campo "_id" único que
  atua como chave primária.

## 2. Operações Básicas
- **find**: Usada para buscar e recuperar documentos de uma coleção.
- **insert**: Utilizada para adicionar novos documentos a uma coleção.
- **update**: Permite modificar documentos existentes em uma coleção.
- **delete**: Remove documentos de uma coleção.

## 3. Desempenho de Busca: Indexes
Indexes são estruturas de dados especiais que melhoram a velocidade das
operações de busca. Sem índices, o MongoDB precisa realizar uma varredura
completa da coleção para encontrar os documentos correspondentes. Com índices,
o MongoDB pode limitar a busca a uma parte menor dos dados.

## 4. Index TTL (Time To Live)
O Index TTL é um recurso que permite definir um tempo de vida para os
documentos. Após o período especificado, os documentos são automaticamente
excluídos do banco de dados. Isso é útil para dados que têm relevância
temporária, como sessões de usuário ou logs.

## 5. Utilizando o pymongo

Para instalar a biblioteca de python para uso **síncrono** do mongoDB,
instale-a com:

```bash
pip install pymongo
```

A seguir, crie um arquivo `.py` e preencha-o com o código abaixo:

```python showLineNumbers tilte="exemplo_mongo.py"
from pymongo import MongoClient

# Conexão com o MongoDB
client = MongoClient('mongodb://localhost:27017/')
db = client['meu_banco']

# Inserção de documentos
colecao = db['minha_colecao']
colecao.insert_one({'nome': 'João', 'idade': 30})
colecao.insert_many([{'nome': 'Maria', 'idade': 25}, {'nome': 'Pedro', 'idade': 35}])

# Busca de documentos
for doc in colecao.find({'idade': {'$gt': 30}}):
    print(doc)

# Atualização de documentos
colecao.update_one({'nome': 'João'}, {'$set': {'idade': 31}})

# Exclusão de documentos
colecao.delete_one({'nome': 'Pedro'})

# Criação de um índice TTL
colecao.create_index([('data_criacao', 1)], expireAfterSeconds=3600)
```

Neste exemplo, utilizamos a biblioteca PyMongo para interagir com o MongoDB.
Primeiro, conectamos ao banco de dados e criamos uma coleção. Em seguida,
realizamos operações de inserção, busca, atualização e exclusão de documentos.
Por fim, criamos um índice TTL para que os documentos sejam automaticamente
excluídos após uma hora.

## 6. Criando schemas MongoDB

<Admonition 
    type="info" 
    title="Autoestudo">

<div style={{ textAlign: 'center' }}>
    <iframe 
        style={{
            display: 'block',
            margin: 'auto',
            width: '100%',
            height: '50vh',
        }}
        src="https://www.youtube.com/embed/QAqK-R9HUhc" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>

## 7. Modelagem de dados com MongoDB

<Admonition 
    type="info" 
    title="Autoestudo">

<div style={{ textAlign: 'center' }}>
    <iframe 
        style={{
            display: 'block',
            margin: 'auto',
            width: '100%',
            height: '50vh',
        }}
        src="https://www.youtube.com/embed/3GHZd0zv170" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>

