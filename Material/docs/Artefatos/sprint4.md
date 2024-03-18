---
title: Sprint 4
sidebar_position: 4
sidebar_class_name: artefato
slug: /sprint4
---

# Artefatos Sprint 4 (Computação)

## 1. Gerenciamento de Escalabilidade

### 1.1. Enunciado

Este artefato contempla a melhoria do sistema de modo a visar a escalabilidade
necessária para lidar com o throughput compatível com um cenário de cidades
inteligentes. Para isso, deve-se implementar um sistema de processamento de
eventos complexos escalável horizontalmente.

* [Kafka](https://kafka.apache.org/)
* [Confluent
  Cloud](https://www.confluent.io/cloud-kafka/?utm_medium=sem&utm_source=google&utm_campaign=ch.sem_br.brand_tp.prs_tgt.confluent-brand_mt.xct_rgn.latam_lng.eng_dv.all_con.confluent-kafka-general&utm_term=confluent%20kafka&creative=&device=c&placement=&gad_source=1&gclid=EAIaIQobChMIr9ee5PP9hAMVLWFIAB2VMQ7EEAAYASAAEgI8SfD_BwE)
* [CloudKarafka](https://www.cloudkarafka.com/)

### 1.2. Padrão de entrega

:::warning

O padrão de entrega do artefato contempla uma série de requisitos para que a
atividade seja considerada concluída. Embora não haja atribuição direta de
pontuação aos itens aqui descritos, entenda-se que o **não cumprimento de
qualquer dos requisitos** pode, no melhor dos casos, acarretar em um **desconto
de nota** e, no pior caso, **invalidar a entrega por completo**. Fique atento!

:::

1. O código fonte da solução deve estar disponível no repositório do grupo no
   Github (na branch `main`), em um diretório denominado `src`. É provável que
   a sua solução não seja monolítica. Se esse for o caso, os módulos que compõe
   o seu código-fonte devem ser divididos em subdiretórios.
2. Os testes desenvolvidos devem estar disponíveis no diretório `tests` e,
   quando possível, aplicados de forma automatizada. Os testes realizados sem
   automatização devem ter seu procedimento de execução e resultados claramente
   descritos em um arquivo contido dentro do diretório de testes.
3. As instruções para que o parceiro possa executar o projeto devem estar
   claramente discriminadas na documentação, com um link no README para que a
   seção possa ser prontamente acessada.
4. O projeto em seu estado atual deve estar disponível em um release do Github
   cujo nome deve incluir a numeração da sprint (e.g. release-sprint-1).
5. As questões centrais de desenvolvimento da sprint e os testes devem ser
   apresentadas de forma clara durante o review com o parceiro.

### 1.3. Padrão de qualidade

**[0,0 - 1,0]**
Sistema de processamento de eventos complexos implementado de forma
incorreta/solução fora de contexto com o projeto.

**[2,0 - 4,0]**
Sistema de processamento de eventos complextos implementado, mas procedimentos
de teste evidenciados para validá-lo. O sistema ainda não está integrado com o
broker/banco de dados.

**[4,0 - 6,0]**
Sistema de processamento de eventos complexos implementado e com testes
validando o seu funcionamento e criação de logs. A integração com o broker foi
feita, porém não há testes para essa integração.

**[6,0 - 8,0]**
Sistema de processamento de eventos complexos implementado, integrado com o
broker e com testes para essa integração. Ainda não há integração do consumer
com o banco de dados do projeto.

**[8,0 - 9,0]**
Sistema de processamento de eventos complexos implementado, integrado com o
broker e com testes para essa integração. Também há integração do consumer com
o banco de dados, porém não há testes definidos para essa etapa de integração.

**[9,0 - 10,0]**
Sistema de processamento de eventos complexos implementado, integrado e com
testes de implementação e integração feitos.

## 2. Documentação 

### 2.1. Enunciado

O grupo deverá registrar a evolução de escalabilidade em um relatório técnico.

### 2.2. Padrão de entrega

:::warning

O padrão de entrega do artefato contempla uma série de requisitos para que a
atividade seja considerada concluída. Embora não haja atribuição direta de
pontuação aos itens aqui descritos, entenda-se que o não cumprimento de
qualquer dos requisitos pode, no melhor dos casos, acarretar em um desconto de
nota e, no pior caso, invalidar a entrega por completo. Fique atento!

:::

1. O texto desenvolvido pelo grupo deve estar disponível em uma página estática
   gerada pelo framework Docusaurus. Para tal, deve haver um diretório no
   repositório denominado `docs`, onde ficará a raíz do Docusaurus;
2. A documentação da sprint 2 deve estar inteiramente contida em uma seção
   denominada `Sprint 2`. Cada um dos artefatos deve ter sua descrição contida
   em uma página ou subseção dentro da seção `Sprint 2`.
3. As figuras utilizadas no documento devem sempre ser referenciadas no texto,
   com descrições textuais que estimulem a coesão entre o que é apresentado
   visualmente e o restante do texto; e 
4. Todas as figuras utilizadas na documentação devem estar salvas dentro do
   diretório `docs`, em um subdiretório chamado `static`.

### 2.3. Padrão de qualidade

#### 2.3.1 Qualidade textual (até 5,0 pontos)

:::warning

A qualidade textual à qual esse padrão de qualidade se refere diz respeito ao
que foi documentado no decorrer da sprint atual. No entanto, é prerrogativa do
professor orientador descontar nota por defeitos que persistiram de um sprint

:::

**[0,0 - 1,0]**
A documentação desenvolvida foge do contexto do projeto/está consideravelmente
incompleta.

**[1,0 - 2,0]**
A documentação desenvolvida aborda todos os temas obrigatórios (artefatos), mas
não apresenta-os de forma coesa, de modo que não seja possível vislumbrar o
objetivo conjunto do projeto. O documento apresenta uma linguagem
consistentemente inadequada para um documento de cunho técnico e/ou apresenta
mais de um estilo textual de forma dissonante (colcha de retalhos).

**[2,0 - 4,0]**
Documentação aborda todos os temas e apresenta-os com alguma coesão, dando um
entendimento ao menos parcial do projeto como um todo. Há poucas ou nenhuma
ocorrência de linguagem inadequada e o texto apresenta alguma consistência de
estilo.

**[4,0 - 5,0]**
Documentação completa e coesa do projeto desenvolvido. A linguagem utilizada é
condizente com o tipo de documento produzido.

#### 2.3.2 Testes (até 3,0 pontos)

**[0,0 - 0,5]**
Há pouca ou nenhuma documentação relacionada aos testes definidos e
implementados no decorrer do sprint.

**[0,5 - 2,5]**
Há documentação para os testes, mas estes estão incompletos e/ou não há testes que
abordem todos os requisitos relacionados à etapa atual do projeto.

**[2,5 - 3,0]**
Os testes estão suficientemente documentados e abordam todos os requisitos
relacionados à etapa atual do projeto.

#### 2.3.3 Instruções para execução (até 2,0 pontos)

**[0,0 - 0,5]**
Há pouca ou nenhuma instrução para execução do projeto em sua versão atual.

**[0,5 - 1,0]**
As instruções estão presentes e abordam de modo geral o que deve ser feito para
a execução do projeto, porém faltam detalhes que não podem ser facilmente
inferidos pelo leitor. As instruções apresentam uma falta de exemplos com
comandos claros para execução. Não há informações sobre compatibilidade do
software com versões de linguagens de programação, bibliotecas ou sistemas
operacionais.

**[1,0 - 1,5]**
Instruções completas e com comandos detalhados de execução, porém com poucos ou
nenhum dos exemplos mais comuns de utilização. Informações incompletas sobre
compatibilidade.

**[1,5 - 2,0]**
Instruções completas, com exemplos mais comuns, comandos detalhados de execução
e informações robustas de compatiblidade.
Requisitos contextualizados, com métricas bem definidas e com o planejamento de
testes devidamente documentado e coerente com os requisitos.
para outro.
