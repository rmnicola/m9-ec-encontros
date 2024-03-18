---
title: Kafka
sidebar_position: 2
sidebar_class_name: autoestudo
slug: /kafka
---

import Admonition from '@theme/Admonition';

# Introdução ao Kafka

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
        src="https://www.youtube.com/embed/uvb00oaa3k8" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>

O Apache Kafka é uma plataforma de streaming distribuído que permite a
construção de aplicações em tempo real e de alta performance para o
processamento de fluxos de dados. Desenvolvido inicialmente pelo LinkedIn e
posteriormente doado à Apache Software Foundation, o Kafka é amplamente
utilizado para a construção de sistemas de mensagens, processamento de logs,
monitoramento de dados em tempo real, entre outros.

## 1. Componentes do kafka

<img 
  src="https://www.peerbits.com/static/23cabd7755efc9e60f1b9c7da278a880/31b09/components-and-concepts-of-apache-kafka-peerbits.webp"
  alt="Apache kafka" 
  style={{ 
    display: 'block',
    marginLeft: 'auto',
    maxHeight: '70vh',
    marginRight: 'auto'
  }} 
/>
<br/>

Os componentes principais do Apache Kafka são:

1. **Produtores (Producers):** São os componentes responsáveis por publicar
   (produzir) mensagens nos tópicos do Kafka. Eles podem ser aplicações ou
   sistemas que geram dados e os enviam ao Kafka para serem processados ou
   armazenados.

2. **Consumidores (Consumers):** São os componentes que se inscrevem em um ou
   mais tópicos para consumir (ler) as mensagens. Eles podem ser aplicações que
   processam os dados recebidos do Kafka de acordo com a lógica de negócios.

3. **Tópicos (Topics):** São as categorias ou canais nos quais as mensagens são
   publicadas. Os tópicos são divididos em partições, que são as unidades
   básicas de armazenamento e paralelismo no Kafka.

4. **Partições (Partitions):** São subdivisões de um tópico. Cada partição é um
   log sequencial imutável de mensagens. As partições permitem o escalonamento
   horizontal do Kafka, pois diferentes partições podem ser armazenadas e
   processadas em diferentes servidores.

5. **Brokers:** São os servidores que compõem o cluster do Kafka. Eles são
   responsáveis por armazenar as mensagens e atender às solicitações dos
   produtores e consumidores. Cada broker pode armazenar um ou mais tópicos e
   suas respectivas partições.

6. **ZooKeeper:** Embora não seja um componente exclusivo do Kafka, o ZooKeeper
   é essencial para o gerenciamento do cluster do Kafka. Ele é usado para
   coordenar os brokers, gerenciar a configuração dos tópicos e
   particionamento, e manter o estado do cluster.

7. **Controlador (Controller):** É um broker especial dentro do cluster que é
   responsável por manter a saúde do cluster e realizar operações
   administrativas, como a eleição de líderes para as partições.

8. **Líder de Partição (Partition Leader):** Dentro de cada partição, um dos
   brokers é eleito como líder. O líder é responsável por lidar com todas as
   leituras e gravações para essa partição.

## 2. Fundamentos do Kafka

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
        src="https://www.youtube.com/embed/B5j3uNBH8X4" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>

## 3. Performance do Kafka

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
        src="https://www.youtube.com/embed/UNUz1-msbOM" 
        frameborder="0" 
        allowFullScreen>
    </iframe>
</div>

</Admonition>
