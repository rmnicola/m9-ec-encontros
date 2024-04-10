---
title: Prova 1
sidebar_position: 15
slug: /the-gods-are-always-watching
unlisted: true
---

# PSUB Módulo 9 EC 2024

## 1. Enunciado
   
Você está desenvolvendo um aplicativo de monitoramento remoto para acompanhar a
umidade em uma rede de estufas agrícolas. Cada estufa possui diversos sensores
de umidade distribuídos em pontos estratégicos, como nas áreas de cultivo de
hortaliças e flores. Esses sensores transmitem periodicamente as leituras de
umidade em tópicos específicos do MQTT. Seu aplicativo deve se conectar a esses
tópicos para coletar os dados e exibir as leituras de umidade em tempo real em
uma interface gráfica. Além disso, o aplicativo deve emitir um alerta na
interface se a umidade estiver fora dos limites aceitáveis.

Os limites de umidade são:

* Área de Hortaliças: Umidades abaixo de 30% ou acima de 70% devem acionar um
  alerta.
* Área de Flores: Umidades abaixo de 40% ou acima de 80% devem acionar um
  alerta.

Para simular a operação do sistema, você deve criar um pequeno simulador
(publisher) capaz de enviar mensagens simulando os sensores de umidade. Para
tal, considere que o **payload** dos sensores inclui as seguintes informações:

* O id do dispositivo (por esse id deve ser possível identificar a que estufa
  pertence o sensor)
* O tipo de dispositivo (hortaliças ou flores)
* O valor de umidade medido
* O timestamp de quando ocorreu a medição

Um exemplo de payload pode ser visto abaixo:

```json
{
 "id": "gh01h01",
 "tipo": "hortaliças",
 "umidade": 45,
 "timestamp": "02/03/2024 15:45"
}
```

Exemplo de saída esperada na interface (use como referência das informações que
precisam estar lá, não precisa ser necessariamente com a mesma formatação):

```bash
Estufa 2: Flores 3 | 85% Umidade [ALERTA: Umidade ALTA]
Estufa 1: Hortaliças 1 | 45% Umidade
Estufa 1: Hortaliças 2 | 25% Umidade [ALERTA: Umidade BAIXA]
Estufa 3: Flores 1 | 50% Umidade
```

O seu projeto deve obrigatóriamente apresentar **testes** para cada uma das
partes envolvidas no sistema. Os testes específicos que devem ser implementados
podem ser vistos no **padrão de qualidade** desta atividade.

Para esta atividade você pode escolher entre usar a linguagem **Go** ou
**Python**. Para simular a comunicação MQTT, sugere-se o uso do mosquitto. Os
testes desenvolvidos **não precisam** botar o broker para rodar.

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
4. O prazo para a entrega desta atividade é até o dia 08/03/2024 às 10h00min

## 3. Padrão de qualidade

:::warning

Note que metade da nota da atividade é composta por testes. Como vimos nos
conceitos de TDD, é possível e muitas vezes desejável criar testes antes de
desenvolver a aplicação. Sendo assim, o padrão de qualidade da nossa prova
prática permite que você apresente testes corretos sem que a aplicação em si
passe nesses testes. Use essa informação com sabedoria.

:::

### 3.1. Simulador (até 4.0 pontos)

#### 3.1.1. Funcionalidade

**[0,0 - 1,0]**
O simulador desenvolvido não se adequa aos sensores descritos

**[1,0 - 2,0]**
O simulador desenvolvido aparentemente atende os requisitos do payload descrito
e apresenta valores verossímeis.

#### 3.1.2. Testes

**[0,0 - 2,0]**
Foram implementados testes para validar: 

1. O correto envio das mensagens com QoS = 1; 
2. O payload recebido, garantindo que todos os campos necessários estejam
   presentes e preenchidos

### 3.2. Sistema de monitoramento (até 6.0 pontos)

#### 3.2.1. Funcionalidade

**[0,0 - 1,0]**
O sistema de monitoramento não apresenta uma ou mais informações necessárias
descritas no enunciado da atividade.

**[1,0 - 2,0]**
O sistema de monitoramento apresenta todas as informações necessárias descritas
no enunciado, mas os alarmes ainda não foram implementados ou foram
implementados de forma incompleta.

**[2,0 - 3,0]**
Sistema de monitoramento correto e com alarmes aparentemente funcionais.

#### 3.1.2. Testes

**[0,0 - 3,0]**
Foram implementados testes para validar: 

1. O correto recebimento das mensagens com QoS = 1;
2. A integridade das mensagens recebidas, garantindo que o subscriber não
   distorce as informações recebidas; e
3. O correto funcionamento de cada um dos tipos de alarmes necessários para o
   correto funcionamento do sistema (temperatura alta e baixa para os dois
   tipos de sensores).
