---
title: MongoDB Atlas com Python
sidebar_position: 2
sidebar_class_name: autoestudo
slug: /altas-python
---

import Admonition from '@theme/Admonition';

# Utilizando o MongoDB Atlas c/ Python

Para interagir com um banco de dados MongoDB Atlas usando Python, você pode
usar a biblioteca `pymongo`. Abaixo está um exemplo simples que mostra como se
conectar a um cluster do MongoDB Atlas, inserir um documento em uma coleção e
recuperar documentos da coleção.

Primeiro, certifique-se de ter o `pymongo` instalado. Se não tiver, você pode
instalá-lo usando pip:

```bash
pip install pymongo
```

Em seguida, use o seguinte código Python como exemplo:

```python
from pymongo import MongoClient

# Substitua 'your_connection_string' pela sua string de conexão do MongoDB
# Atlas
connection_string = 'your_connection_string'

# Conectar ao cluster do MongoDB Atlas
client = MongoClient(connection_string)

# Selecionar o banco de dados
db = client['nome_do_banco_de_dados']

# Selecionar a coleção
collection = db['nome_da_coleção']

# Inserir um documento na coleção
document = {'nome': 'João', 'idade': 30}
collection.insert_one(document)

# Recuperar documentos da coleção
for doc in collection.find():
    print(doc)
```

Certifique-se de substituir `'your_connection_string'` pela sua string de
conexão do MongoDB Atlas, que você pode obter do painel do Atlas. Além disso,
substitua `'nome_do_banco_de_dados'` e `'nome_da_coleção'` pelos nomes
apropriados do seu banco de dados e coleção, respectivamente.
