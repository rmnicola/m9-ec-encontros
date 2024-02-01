---
title: Sprint 3
sidebar_position: 3
sidebar_class_name: artefato
slug: /sprint3
---

# Artefatos Sprint 3 (Computação)

## 1. Evolução do banco de dados e dashboard

### 1.1. Enunciado

O grupo deverá modificar o projeto, integrando o sistema a um sistema de
armazenamento em núvem com segurança (e.g. MongoDB Atlas com whitelist ou
blacklist de IP) assim como evoluir o frontend desenvolvido previamente para
uma solução feita sob medida (e.g., aplicação web) e com autenticação (e.g.,
JWT, OAuth2).

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

1. O sistema de armazenamento em nuvem deve demonstrar os seguintes critérios
   de segurança: confidencialidade, integridade e autenticidade. Ao menos um
   teste de ataque deve alvejar cada um desses requisitos (até 2,5 pontos);
2. Os testes devem demonstrar que um dispositivo IoT (real ou simulado) é capaz
de enviar dados periódicamente, validando o armazenamento no banco de dados e a
exibição no dashboard (até 2,5 pontos);
3. Ao menos um teste de ataque deve alvejar o sistema de autenticação do
   dashboard (até 2,5 pontos);
4. Os testes de segurança (ataques) devem apresentar pré-condição, etapas do
   teste, e pós-condição (até 2,5 pontos).

## 2. Documentação 

### 2.1. Enunciado

O grupo deverá registrar a evolução do sistema de armazenamento, com a adição
do armazenamento em nuvem com segurança. Também deve descrever o
desenvolvimento do dashboard feito sob medida e seu sistema de autenticação de
usuários.

### 2.2. Padrão de entrega

1. O texto desenvolvido pelo grupo deve estar disponível no repositório do
   Github (branch 'main'), dentro da pasta 'docs' e em formato markdown;
2. As figuras utilizadas no documento devem sempre ser referenciadas no texto,
   com descrições textuais que estimulem a coesão entre o que é apresentado
   visualmente e o restante do texto;
3. Todas as figuras utilizadas na documentação devem estar salvas dentro da
   pasta 'docs', em uma subpasta chamada 'assets'.

### 2.3. Padrão de qualidade

1. O relatório deve apresentar uma versão revisada da arquitetura com uma
   descrição textual e usando o diagrama de implantação UML (até 2,5 pontos);
2. O relatório deve referenciar os testes de cada novo componente implementado,
   e os requisitos associados a cada novo componente (até 2,5 pontos);
3. O relatório deve apresentar pelo menos um diagrama de sequência revisado
   mostrando a comunicação segura entre dashboard e o restante do sistema (até
   2,5 pontos);
4. O relatório deve apresentar pelo menos um diagrama de sequência revisado
   mostrando a comunicação segura entre banco de dados e o restante do sistema
   (até 2,5 pontos).
