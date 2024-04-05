---
title: Prova 2
sidebar_position: 16
slug: /would-you-kindly
unlisted: true
---

# P2 Módulo 9 EC 2024

## 1. Enunciado

As cidades inteligentes utilizam tecnologia de ponta para melhorar a qualidade
de vida dos seus habitantes e a eficiência dos serviços urbanos. Um aspecto
crucial é a capacidade de coletar, transmitir e processar dados em tempo real
para diversos fins, como monitoramento de tráfego, gestão de resíduos, controle
de iluminação pública e resposta a emergências. Apache Kafka, uma plataforma de
streaming de eventos distribuída, desempenha um papel vital na implementação
dessas funcionalidades, permitindo a comunicação eficiente e confiável entre
diferentes sistemas e serviços.

Uma cidade está implementando um novo sistema de monitoramento de qualidade do
ar. O sistema consiste em diversos sensores espalhados pela cidade que captam
níveis de poluentes em tempo real. Os dados coletados pelos sensores precisam
ser transmitidos para um sistema central de análise de dados para monitoramento
e resposta rápida em caso de detecção de níveis perigosos de poluição. Ao
implementar o sistema, os especialistas perceberam que o sistema central é
incapaz de processar muitas requisições para armazenamento de dados de sensores
ao mesmo tempo e, quando o sistema alcança o seu limite, começa a perder dados
enviados. Você foi contratado para resolver esse problema. Para isso, você deve
utilizar o Apache Kafka para implementar uma fila de eventos e, assim, criar um
amortecimento com persistência para que o sistema consiga absorver os momentos
de pico sem perdas de dados. Para fazer isso, você deve:

1. Implementar um Producer (Produtor): Deve coletar dados simulados de sensores
   de qualidade do ar e publicá-los em um tópico do Kafka chamado qualidadeAr.
   Os dados devem incluir: 
    * Id do sensor, timestamp, tipo de poluente e nivel da medida.
2. Implementar um Consumer (Consumidor): Deve assinar o tópico qualidadeAr e
   processar os dados recebidos, exibindo-os em um formato legível, além de
   armazená-los para análise posterior (escolha a forma de armazenamento que
   achar mais adequada).
3. Implementar testes de Integridade e Persistência: Criar testes automatizados
   que validem a integridade dos dados transmitidos (verificando se os dados
   recebidos são iguais aos enviados) e a persistência dos dados (assegurando
   que os dados continuem acessíveis para consulta futura, mesmo após terem
   sido consumidos).

:::tip

Não há restrições quanto a usar uma versão local ou em nuvem do Kafka. Tendo
dito isso, sejam cautelosos. Se por algum motivo o seu cluster Kafka em nuvem
não estiver responsivo, use uma versão local ;)

:::

Abaixo, pode-se ver um exemplo de payload enviado pelo sensor:

```json
{
    "idSensor": "sensor_001",
    "timestamp": "2024-04-04T12:34:56Z",
    "tipoPoluente": "PM2.5",
    "nivel": 35.2
}
```

## 2. Padrão de entrega

:::warning

Esses são os critérios mínimos para que eu considere a atividade como entregue.
Fique atento, pois o não cumprimento de qualquer um desses critérios pode, no
melhor dos casos, gerar um desconto de nota e, no pior deles, invalidar a
atividade.

:::

1. A atividade deve ser feita em um repositório **aberto** no github. Seu link
   deve ser fornecido na resposta do forms da prova;
2. No README do repositório deve ter instruções claras de como instalar e rodar
   o sistema criado, comandos em blocos de código e uma expliação sucinta do
   que fazem;
3. Ainda no README, deve haver um vídeo/imagens demonstrando plenamente o
   funcionamento do sistema criado;
4. O prazo para a entrega desta atividade é até o dia 04/04/2024 às 11h00min

## 3. Padrão de qualidade

:::warning

Note que metade da nota da atividade é composta por testes. Como vimos nos
conceitos de TDD, é possível e muitas vezes desejável criar testes antes de
desenvolver a aplicação. Sendo assim, o padrão de qualidade da nossa prova
prática permite que você apresente testes corretos sem que a aplicação em si
passe nesses testes. Use essa informação com sabedoria.

:::

### 3.1. Producer (até 4.0 pontos)

**[0,0 - 1,0]**
O producer não é capaz de enviar as mensagens conforme o formato especificado
para o broker Kafka.

**[1,0 - 3,0]**
O producer é capaz de enviar as mensagens conforme o formato especificado de
forma consistente.

**[3,0 - 4,0]**
Além de enviar as mensagens de acordo com o formato especificado, o código
desenvolvido permite a modificação do formato de mensagem sem precisar de um
refatoramento substancial (e.g. se adapta ao formato especificado em um arquivo
JSON/YAML)


### 3.2. Consumer (até 3.0 pontos)

**[0,0 - 1,0]**
O consumer não é capaz de consumir as mensagens do broker Kafka conforme o
formato especificado.

**[1,0 - 2,0]**
O consumer é capaz de consumir as mensagens do broker conforme o formato
espeficiado, mas a exibição para o usuário não foi implementada ou foi
implementada de uma forma pouco amigável (e.g. só "joga" o payload na tela sem
formatação alguma)

**[2,0 - 3,0]**
O consumer é capaz de consumir as mensagens do broker Kafka e exibe-as de forma
clara e legível em uma interface de terminal.

### 3.3. Testes (até 3.0 pontos)

Os testes necessários para esse projeto são:

* Teste de integridade - garante que os dados enviados pelo producer chegam ao
  tópico do Kafka **sem alterações**. (até 2.0 pontos)
* Teste de persistência - garante que os dados enviados se mantem disponíveis
  para acesso em disco mesmo após o seu consumo inicial. (até 1.0 ponto)
