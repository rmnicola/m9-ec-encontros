---
title: Sprint 1
sidebar_position: 1
slug: /sprint1
---

# Artefatos Sprint 1 (Computação)

## 1. Proposta inicial e simulador de hardware

### 1.1. Enunciado

O grupo deverá implementar um cliente MQTT que se comunica com um broker
público (e.g., utilizando a bilbioteca paho-mqtt do Python). O intuito é a
capacidade de simular a comunicação gerada pelos dispositivos do parceiro,
demonstrando a compatibilidade com o protocolo de comunicação por ele utilizado
assim como as especificações dos componentes em si. Testes deverão ser
especificados, e seus resultados devem ser registrados para mostrar o correto
funcionamento deste primeiro protótipo.

### 1.2. Padrão de entrega

Esses são os critérios mínimos para que eu considere a atividade como entregue.
Fique atento, pois o não cumprimento de qualquer um desses critérios pode, no
melhor dos casos, gerar um desconto de nota e, no pior deles, invalidar a
atividade.

1. O código fonte da solução deve estar disponível no repositório do grupo no
   Github (na branch 'main'), em uma pasta denominada 'src'.
2. Os testes desenvolvidos devem estar disponíveis na pasta 'tests' e, quando
   possível, aplicados de forma automatizada (e.g., utilizando o Github
   Actions). Quando não for possível essa integração, deve-se descrever os
   testes e os resultados em um arquivo markdown dentro da pasta 'tests'. 
3. As instruções para que o parceiro possa executar o projeto devem estar
   claramente discriminadas na documentação, com um link no README para que a
   seção possa ser prontamente acessada.
4. O projeto em seu estado atual deve estar disponível em um release do Github
   cujo nome deve incluir a numeração da sprint.
5. As questões centrais de desenvolvimento da sprint e os testes devem ser
   apresentadas de forma clara durante o review com o parceiro.

### 1.3. Padrão de qualidade

1. O código e respectivos comentários apresentam informações de conexão (e.g.,
   identificador do cliente, endereço do broker MQTT) (até 2,0 pontos);
2. Os resultados dos testes mostram que o protótipo publica em pelo menos 1
   tópico com uma frequência compatível com os dispositivos especificados pelo
   parceiro (até 2,0 pontos);
3. Os resultados dos testes mostram que o protótipo é capaz de receber dados de
   um tópico ao qual se subscreve e, em resposta, publicar dados em outro
   tópico (até 2,0 pontos);
4. Os resultados dos testes mostram que o protótipo é capaz de interagir com
   o(s) protocolo(s) de comunicação especificados pelo cliente, utilizando um
   middleware para convertê-lo(s) para MQTT quando necessário (até 2,0 pontos);
5. Os testes devem apresentar pré-condição, etapas do teste, e pós-condição
   (até 2,0 pontos).

## 2. Documentação 

### 2.1. Enunciado

O grupo deverá especificar requisitos funcionais, não funcionais e restrições,
testes relacionados e propor uma arquitetura inicial para o projeto. Pelo menos
um requisito deve estar relacionado à segurança e pelo menos um requisito deve
estar relacionado à escalabilidade.
### 2.2. Padrão de entrega

1. O texto desenvolvido pelo grupo deve estar disponível no repositório do
   Github (branch 'main'), dentro da pasta 'docs' e em formato markdown;
2. As figuras utilizadas no documento devem sempre ser referenciadas no texto,
   com descrições textuais que estimulem a coesão entre o que é apresentado
   visualmente e o restante do texto;
3. Todas as figuras utilizadas na documentação devem estar salvas dentro da
   pasta 'docs', em uma subpasta chamada 'assets'.

### 2.3. Padrão de qualidade

1. O grupo deve especificar requisitos funcionais e não funcionais conforme a
   norma ISO 25010 (até 1,5 pontos);
2. Os requisitos não funcionais devem apresentar uma estratégia de validação
   associada (e.g., testes com aspectos quantitativos) (até 2,0 pontos);
3. Na descrição textual, o grupo deve registrar como os requisitos
   especificados estão contemplados na arquitetura (até 2,0 pontos);
4. Pelo menos uma figura da arquitetura deve ser um diagrama UML (e.g.,
   sequência, implantação, classes, entre outros) (até 1,5 pontos).
5. Pelo menos um requisito deve estar relacionado à segurança. Este requisito
   deve contemplar vulnerabilidades e sua localização na arquitetura proposta
   (até 1,0 pontos);
6. Pelo menos três requisitos devem estar relacionados aos seguintes aspectos:
   integridade dos dados, disponibilidade de serviço e tempo de resposta (até
   1,0 ponto);
7. Pelo menos um requisito deve estar relacionado à escalabilidade,
   contemplando o volume de requisições máximo que sistema suporta (até 1,0
   ponto).
