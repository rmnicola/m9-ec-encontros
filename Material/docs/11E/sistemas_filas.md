---
title: Sistemas de filas
sidebar_position: 1
sidebar_class_name: autoestudo
slug: /filas
---

import Admonition from '@theme/Admonition';

# Sistemas de fila

O processamento de eventos complexos é uma abordagem importante na análise e
gerenciamento de grandes volumes de dados em tempo real. As filas de mensagens
desempenham um papel crucial nesse processo, permitindo a comunicação
assíncrona entre diferentes partes de um sistema. Ao lidar com o processamento
de eventos, uma consideração importante é a escolha entre sistemas baseados em
disco e sistemas baseados em memória, cada um com suas próprias vantagens e
desvantagens.

## 1. O papel das filas no processamento de eventos

As filas desempenham um papel crucial no processamento de alto volume de dados,
especialmente em sistemas distribuídos e arquiteturas baseadas em
microserviços. Elas atuam como amortecedores entre os produtores de dados e os
consumidores, permitindo que o processamento ocorra de forma assíncrona e
ajudando a gerenciar picos de carga. Isso é essencial para manter a
estabilidade e a escalabilidade do sistema.

### 1.1. Função da Fila como Amortecedor

Uma das principais funções das filas no processamento de alto volume de dados é
servir como um amortecedor entre os componentes que produzem dados e os que os
consomem. Isso significa que as filas podem armazenar temporariamente os dados
até que o consumidor esteja pronto para processá-los, evitando assim o
congestionamento e a perda de dados durante picos de tráfego.

**Exemplo Prático: Sistema de E-commerce**

Imagine um sistema de e-commerce que recebe um alto volume de pedidos durante
uma promoção. Cada pedido precisa ser processado, o que inclui verificar o
estoque, atualizar o banco de dados de pedidos, processar o pagamento e iniciar
o processo de envio.

Sem uma fila, o sistema poderia ficar sobrecarregado, levando a tempos de
resposta lentos ou até mesmo falhas do sistema. Com uma fila, os pedidos podem
ser enfileirados assim que chegam e processados de forma assíncrona. Isso
permite que o sistema absorva os picos de tráfego, mantendo um processamento
estável e evitando a sobrecarga dos componentes de processamento.

A fila atua como um amortecedor, absorvendo os picos de demanda e permitindo
que os componentes do sistema trabalhem em um ritmo sustentável. Isso ajuda a
garantir a disponibilidade do sistema, a consistência dos dados e a experiência
positiva do usuário, mesmo em momentos de alta demanda.

## 2. Tipos de sistema de processamento de eventos

Sistemas de processamento de eventos são fundamentais para lidar com fluxos
contínuos de dados em tempo real. Esses sistemas podem ser categorizados com
base em como armazenam e gerenciam os dados: em disco ou em memória. Cada tipo
tem suas próprias características, vantagens e desvantagens.

### 2.1. Sistemas de Processamento de Eventos em Disco

Sistemas em disco armazenam os eventos em um meio de armazenamento persistente,
como um disco rígido ou SSD. Eles são projetados para garantir a durabilidade e
a retenção dos dados, mesmo em caso de falhas do sistema. 

**Exemplos:**
- **Apache Kafka:** Um sistema de mensagens distribuído que armazena streams de
  registros em tópicos. Kafka é projetado para lidar com alto volume de dados e
  oferece retenção configurável, replicação e particionamento.
- **Apache Pulsar:** Semelhante ao Kafka, o Pulsar é uma plataforma de
  mensagens e streaming distribuída que fornece armazenamento persistente e
  recursos de processamento de eventos.

**Vantagens:**
- **Durabilidade:** Os dados são armazenados de forma segura e podem sobreviver
  a falhas do sistema.
- **Retenção de Dados:** Os eventos podem ser retidos por períodos prolongados,
  permitindo reprocessamento e análise histórica.

**Desvantagens:**
- **Latência:** O acesso a disco é geralmente mais lento do que o acesso à
  memória, o que pode resultar em maior latência no processamento de eventos.
- **Custo:** A necessidade de armazenamento em disco pode aumentar os custos de
  infraestrutura.

### 2.3. Sistemas de Processamento de Eventos em Memória

Sistemas em memória armazenam os eventos na memória RAM, proporcionando acesso
rápido e baixa latência. Eles são ideais para cenários que exigem processamento
em tempo real e onde a persistência de longo prazo dos dados não é uma
prioridade.

**Exemplos:**
- **Redis Streams:** Uma extensão do banco de dados Redis que oferece um
  mecanismo de fila de mensagens em memória com suporte a múltiplos
  consumidores.
- **Apache Flink:** Uma plataforma de processamento de streams que oferece
  processamento de eventos em memória, com capacidade de checkpoint para
  garantir a consistência dos dados.

**Vantagens:**
- **Baixa Latência:** O acesso rápido à memória permite o processamento de
  eventos em tempo real.
- **Simplicidade:** Sistemas em memória tendem a ser mais fáceis de configurar
  e gerenciar em comparação com sistemas baseados em disco.

**Desvantagens:**
- **Volatilidade:** Os dados armazenados na memória são perdidos em caso de
  falha do sistema, a menos que haja mecanismos de persistência ou replicação.
- **Limitação de Capacidade:** A quantidade de dados que podem ser armazenados
  é limitada pela capacidade da RAM disponível.

## 3. Ordem de processamento de eventos

A ordem de processamento de eventos é um aspecto crítico em sistemas de
mensageria e processamento de eventos, especialmente em aplicações onde a
sequência dos eventos afeta o resultado final. Vamos explorar a diferença entre
o Apache Kafka e o RabbitMQ em relação à ordem de processamento dos eventos.

### 3.1. Apache Kafka

O Apache Kafka é projetado para manter a ordem dos eventos dentro de um tópico.
Cada tópico é dividido em partições, e as mensagens dentro de uma partição são
ordenadas e identificadas por um índice sequencial chamado offset. O Kafka
garante que os consumidores recebam as mensagens em uma partição na mesma ordem
em que foram produzidas. Isso é possível porque o Kafka armazena as mensagens
em disco de forma sequencial e os consumidores mantêm o controle de qual offset
eles já processaram.

A ordenação estrita é garantida apenas dentro de uma partição. Se um tópico
tiver várias partições, a ordem global dos eventos pode não ser preservada, a
menos que todas as mensagens relacionadas sejam enviadas para a mesma partição.
Isso geralmente é alcançado usando uma chave de partição consistente.

### 3.2. RabbitMQ

O RabbitMQ, por outro lado, é um broker de mensagens que implementa o protocolo
AMQP (Advanced Message Queuing Protocol). Ele oferece diversas funcionalidades,
incluindo roteamento de mensagens, trocas de mensagens (exchanges) e filas.
Embora o RabbitMQ garanta a ordem das mensagens dentro de uma única fila, há
cenários em que a ordem de processamento pode ser afetada:

1. **Concorrência e Acknowledgements:** Se vários consumidores estiverem lendo
   de uma fila e processando mensagens de forma assíncrona, a ordem de
   processamento pode variar dependendo do tempo de processamento de cada
   mensagem e do momento em que o acknowledgment (confirmação de recebimento) é
   enviado de volta ao RabbitMQ.

2. **Roteamento e Exchanges:** O RabbitMQ permite roteamento complexo de
   mensagens através de exchanges, que podem direcionar mensagens para
   múltiplas filas com base em regras de roteamento. Isso pode levar a uma
   ordem de processamento variável, dependendo de como as mensagens são
   distribuídas entre as filas.

3. **Prioridades de Mensagens:** O RabbitMQ suporta filas com prioridades de
   mensagens, o que pode alterar a ordem de entrega das mensagens se elas
   tiverem prioridades diferentes.

## 4. RabbitMQ vs Kafka

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
        src="https://www.youtube.com/embed/_5mu7lZz5X4" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>
