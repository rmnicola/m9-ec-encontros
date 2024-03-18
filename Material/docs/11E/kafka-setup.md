---
title: Setup Kafka local
sidebar_position: 3
sidebar_class_name: autoestudo
slug: /kafka-setup
---

import Admonition from '@theme/Admonition';

# Setup de um cluster Kafka local

Para fazer o setup de um cluster Kafka local, basta utilizar os containers do
`zookeeper` e do `kafka` distribuídos pela **confluent**.

## 1. Criando um nó Kafka

Para criar o nosso nó de Kafka, vamos precisar criar dois containers: um para o
Zookeeper e outro para o Kafka em si. O jeito mais fácil de configurar esses
serviços é utilizando o `docker compose`. Para isso, crie um arquivo
`docker-compose.yml` com o conteúdo abaixo:

```yaml title="docker-compose.yml" showLineNumbers
version: '2'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.4.4
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - 22181:2181
  
  kafka:
    image: confluentinc/cp-kafka:7.4.4
    depends_on:
      - zookeeper
    ports:
      - 29092:29092
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:29092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
```

Para rodar o docker compose, use:

```bash
docker-compose up -d
```

Para verificar se o seu cluster está rodando, use:

```bash
docker-compose logs kafka | grep -i started
```

## 2. Criando um cluster

A diferença dessa configuração para a de um nó apenas é que vamos configurar
mais de um `zookeeper` e `kafka`. Quantos? Quantos forem necessários. Em um
cluster em nuvem, essa tarefa tipicamente é gerenciada por um orquestrador como
o `Kubernetes`.

```yaml showLineNumbers title="docker-compose.yml"
version: '2'
services:
  zookeeper-1:
    image: confluentinc/cp-zookeeper:7.4.4
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - 22181:2181

  zookeeper-2:
    image: confluentinc/cp-zookeeper:7.4.4
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - 32181:2181
  
  kafka-1:
    image: confluentinc/cp-kafka:7.4.4
    depends_on:
      - zookeeper-1
      - zookeeper-2

    ports:
      - 29092:29092
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper-1:2181,zookeeper-2:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-1:9092,PLAINTEXT_HOST://localhost:29092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
  kafka-2:
    image: confluentinc/cp-kafka:7.4.4
    depends_on:
      - zookeeper-1
      - zookeeper-2
    ports:
      - 39092:39092
    environment:
      KAFKA_BROKER_ID: 2
      KAFKA_ZOOKEEPER_CONNECT: zookeeper-1:2181,zookeeper-2:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-2:9092,PLAINTEXT_HOST://localhost:39092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
```

## 3. Interagindo com o cluster usando o Conduktor

### 3.1. Instalando o Conduktor

Para instalar a ferramenta GUI do **Conduktor**, a melhor forma é utilizar o
`flatpak`:

```bash
flatpak install flathub io.conduktor.Conduktor
```

:::tip

Note que, para o comando acima funcionar, você precisa instalar o `flatpak` com
seu gerenciador de pacotes e adicionar o `flathub` aos repositórios do
`flatpak`. Isso pode ser feito utilizando o comando abaixo:

```bash
flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
```

:::

A partir daqui você deve conseguir utilizar o launcher do seu sistema para
abrir o Conduktor. Caso isso não seja possível, você pode rodar o programa a
partir da linha de comando:

```bash
flatpak run io.conduktor.Conduktor
```

### 3.2. Utilizando o Conduktor

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
        src="https://www.youtube.com/embed/cBc2SHwN6Ro" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>

