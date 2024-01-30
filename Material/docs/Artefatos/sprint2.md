---
title: Sprint 2
sidebar_position: 2
slug: /sprint2
---

# Artefatos Sprint 2 (Computação)

## 1. Versão inicial do sistema

### 1.1. Enunciado

O grupo deverá implantar um broker MQTT próprio na nuvem com autenticação e
autorização (e.g., versão open source do HiveMQ), registrar os dados em um
banco de dados estruturado (e.g., SQLite) e construir uma versão inicial do
dashboard (e.g., usando Streamlit).

### 1.2. Padrão de entrega

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

1. Broker MQTT próprio na nuvem (e.g., versão open source do HiveMQ implantado
   na AWS) com testes comprovando seu funcionamento (até 2,0 pontos);
2. O Broker MQTT deve implementar autenticação de usuários  (e.g., versão open
   source do HiveMQ com autenticação simples com usuário e senha em ACL) com
   testes comprovando seu funcionamento (até 2,0 pontos);
3. O Broker MQTT deve implementar autorização para os seus tópicos, com pelo
   menos uma restrição de tópico para subscrição e uma restrição para
   publicação, com testes comprovando seu funcionamento (até 2,0 pontos);
4. Registro de dados em um banco de dados estruturado (e.g., SQLite), com
   testes comprovando seu funcionamento (até 2,0 pontos);
5. Primeira versão de dashboard (e.g., Streamlit) mostrando os dados do banco
   de dados estruturado, com testes comprovando seu funcionamento (até 2,0
   pontos).

## 2. Documentação 

### 2.1. Enunciado

O grupo deverá descrever a comunicação segura por meio do protocolo MQTT
utilizando diagramas de sequência UML. Também será necessário o registro de
onde os elementos da arquitetura estão implantados utilizando o diagrama de
implantação UML. Testes de segurança do broker devem ser especificados e
executados para mostrar a correta implementação dos mecanismos de autenticação
e autorização.

### 2.2. Padrão de entrega

1. O texto desenvolvido pelo grupo deve estar disponível no repositório do
   Github (branch 'main'), dentro da pasta 'docs' e em formato markdown;
2. As figuras utilizadas no documento devem sempre ser referenciadas no texto,
   com descrições textuais que estimulem a coesão entre o que é apresentado
   visualmente e o restante do texto;
3. Todas as figuras utilizadas na documentação devem estar salvas dentro da
   pasta 'docs', em uma subpasta chamada 'assets'.

### 2.3. Padrão de qualidade

1. A comunicação no diagrama UML de sequência deve contemplar a autenticação
   (até 2,0 pontos);
2. A comunicação no diagrama UML de sequência deve contemplar a autorização
   (até 2,0 pontos);
3. O sistema completo deve ser descrito em um diagrama UML de implantação (até
2,0 pontos);
4. O relatório deve apresentar os resultados dos ataques no broker próprio
   antes da implementação dos mecanismos de segurança (até 2,0 pontos);
5. O relatório deve apresentar os resultados dos ataques no broker próprio após
   a implementação dos mecanismos de segurança (até 2,0 pontos).
