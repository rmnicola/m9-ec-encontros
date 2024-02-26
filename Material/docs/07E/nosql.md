---
title: NoSQL
sidebar_position: 1
sidebar_class_name: autoestudo
slug: /nosql
---

# Bases de dados NoSQL

## 1. SQL vs NoSQL

Bancos de dados SQL (Structured Query Language) e NoSQL (Not Only SQL) são dois
tipos principais de sistemas de gerenciamento de banco de dados que são usados
para armazenar e gerenciar dados. Eles têm diferenças fundamentais em termos de
estrutura, modelo de dados, escalabilidade, desempenho e casos de uso. Vamos
explorar essas diferenças:

### 1.1. Modelo de dados
   - **SQL**: Utiliza um modelo de dados relacional, onde os dados são
     organizados em tabelas com linhas e colunas. As tabelas são relacionadas
     entre si através de chaves estrangeiras. Este modelo é ideal para dados
     estruturados e relações complexas.
   - **NoSQL**: Adota uma variedade de modelos de dados, incluindo documentos
     (MongoDB), chave-valor (Redis), colunas largas (Cassandra) e grafos
     (Neo4j). Esses modelos são mais flexíveis e adequados para armazenar dados
     não estruturados ou semi-estruturados.

### 1.2. Esquema
   - **SQL**: Possui um esquema rígido, o que significa que a estrutura das
     tabelas e os tipos de dados das colunas devem ser definidos antes de
     inserir os dados. Alterações no esquema podem ser trabalhosas.
   - **NoSQL**: Geralmente tem um esquema flexível ou sem esquema, permitindo a
     adição de novos campos sem a necessidade de modificar o esquema existente.
     Isso oferece maior flexibilidade para lidar com dados variáveis.

### 1.3. Escalabilidade
   - **SQL**: Tradicionalmente, os bancos de dados SQL são escalados
     verticalmente, o que significa que você aumenta a capacidade do servidor
     (CPU, RAM, armazenamento) para lidar com mais carga. Isso pode se tornar
     caro e ter limitações físicas.
   - **NoSQL**: Projetados para escalabilidade horizontal, os bancos de dados
     NoSQL podem ser distribuídos em vários servidores ou nós, facilitando o
     aumento da capacidade de maneira mais econômica e flexível.

### 1.4. Transações
   - **SQL**: Oferece suporte a transações ACID (Atomicidade, Consistência,
     Isolamento, Durabilidade), que são importantes para garantir a integridade
     dos dados em operações que envolvem múltiplas etapas.
   - **NoSQL**: Muitos bancos de dados NoSQL priorizam a disponibilidade e a
     particionabilidade (teorema CAP) em detrimento da consistência estrita,
     seguindo o modelo BASE (Basicamente Disponível, Estado Suave, Eventual
     Consistência). No entanto, alguns sistemas NoSQL começaram a oferecer
     suporte a transações ACID.

### 1.5. Consultas
   - **SQL**: Utiliza uma linguagem de consulta estruturada (SQL) para realizar
     consultas complexas, incluindo junções, agrupamentos e operações de
     agregação.
   - **NoSQL**: As consultas variam de acordo com o tipo de banco de dados
     NoSQL e podem ser menos complexas que as do SQL. Por exemplo, os bancos de
     dados de documentos usam consultas baseadas em JSON, enquanto os bancos de
     dados de grafos usam linguagens específicas para consultas de grafos.

### 1.6. Casos de Uso
   - **SQL**: Ideal para aplicações que requerem transações complexas, como
     sistemas de gerenciamento de relacionamento com o cliente (CRM), sistemas
     de gerenciamento de recursos empresariais (ERP) e sistemas bancários.
   - **NoSQL**: Adequado para aplicações que lidam com grandes volumes de dados
     não estruturados ou semi-estruturados, como big data, análise de dados em
     tempo real, armazenamento de dados de redes sociais e sistemas de
     recomendação.

## 2. ACID vs BASE

ACID e BASE são dois conjuntos de propriedades que garantem a confiabilidade
das transações em sistemas de banco de dados. Eles representam abordagens
diferentes para lidar com a consistência e a disponibilidade dos dados.

### 2.1. ACID 

ACID (Atomicidade, Consistência, Isolamento, Durabilidade):

1. **Atomicidade**: Garante que todas as operações em uma transação sejam
   concluídas com sucesso ou nenhuma seja aplicada. É tudo ou nada.
2. **Consistência**: Assegura que uma transação leve o banco de dados de um
   estado consistente para outro estado consistente, mantendo a integridade dos
   dados.
3. **Isolamento**: Garante que as transações sejam executadas de forma isolada
   umas das outras, evitando interferências e garantindo que os resultados
   sejam os mesmos como se as transações fossem executadas sequencialmente.
4. **Durabilidade**: Uma vez que uma transação foi confirmada, ela permanece
   gravada no banco de dados, mesmo em caso de falha do sistema ou queda de
   energia.

Os sistemas de banco de dados que seguem o modelo ACID, como os bancos de dados
relacionais SQL, são ideais para aplicações que exigem alta consistência e
integridade dos dados, como sistemas financeiros e de contabilidade.

### 2.2. BASE 

BASE (Basicamente Disponível, Estado Suave, Eventual Consistência):

1. **Basicamente Disponível**: O sistema garante a disponibilidade dos dados,
   mesmo na presença de falhas, mas pode não garantir a consistência imediata
   dos dados em todos os nós.
2. **Estado Suave**: O sistema pode estar em um estado de transição por um
   período, permitindo a eventual consistência dos dados.
3. **Eventual Consistência**: O sistema garante que, eventualmente, todos os
   nós terão os mesmos dados, mas não garante quando isso ocorrerá. A
   consistência é alcançada ao longo do tempo.

Os sistemas de banco de dados que seguem o modelo BASE, como muitos bancos de
dados NoSQL, são projetados para escalabilidade e flexibilidade, sacrificando a
consistência imediata em favor da disponibilidade e da tolerância a partições.
Eles são adequados para aplicações que podem tolerar alguma inconsistência
temporária, como redes sociais e sistemas de análise de big data.

## 3. Tipos de bancos NoSQL

Os bancos de dados NoSQL são projetados para armazenar e gerenciar grandes
volumes de dados distribuídos de maneira não relacional. Eles são categorizados
em quatro tipos principais: bancos de dados de documentos, chave-valor, colunas
e grafos. Cada tipo tem suas próprias características e casos de uso ideais:

### 3.1. Bancos de Dados de Documentos:
   - **Exemplo**: MongoDB
   - **Características**: Armazenam dados em documentos, geralmente no formato
     JSON, BSON ou XML. Cada documento pode ter uma estrutura diferente,
     permitindo flexibilidade na representação de dados.
   - **Casos de Uso**: Aplicações que lidam com dados semiestruturados ou
     variáveis, como sistemas de gerenciamento de conteúdo, plataformas de
     e-commerce e aplicações móveis.

### 3.2. Bancos de Dados Chave-Valor:
   - **Exemplo**: Redis, DynamoDB
   - **Características**: Armazenam dados como pares de chave-valor. Eles são
     otimizados para operações de leitura e gravação rápidas e são eficientes
     para armazenar grandes quantidades de dados.
   - **Casos de Uso**: Sessões de usuário, caches, filas de mensagens e
     configurações de aplicativos.

### 3.3. Bancos de Dados de Colunas:
   - **Exemplo**: Cassandra, HBase
   - **Características**: Organizam dados em colunas em vez de linhas, o que é
     eficiente para consultas em grandes conjuntos de dados e para operações de
     agregação. Eles são projetados para escalabilidade horizontal e desempenho
     em consultas analíticas.
   - **Casos de Uso**: Sistemas de análise de big data, armazenamento de séries
     temporais e sistemas de recomendação.

### 3.4. Bancos de Dados de Grafos:
   - **Exemplo**: Neo4j, ArangoDB
   - **Características**: Armazenam dados como nós e arestas em um grafo,
     representando entidades e seus relacionamentos. Eles são otimizados para
     consultas complexas que envolvem relacionamentos entre entidades.
   - **Casos de Uso**: Redes sociais, sistemas de recomendação, análise de
     fraudes e aplicações de mapeamento de rotas.

Cada tipo de banco de dados NoSQL tem suas próprias vantagens e desvantagens, e
a escolha do tipo adequado depende das necessidades específicas da aplicação,
da natureza dos dados e dos padrões de acesso.
