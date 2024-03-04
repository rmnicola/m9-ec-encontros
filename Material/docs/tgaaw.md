---
title: Prova 1
sidebar_position: 15
slug: /the-gods-are-always-watching
unlisted: true
---

# P1 Módulo 9 EC 2024

## 1. Enunciado

Considere um sistema de monitoramento de temperatura para uma
cadeia de supermercados. Cada loja possui vários sensores de temperatura
espalhados em áreas críticas, como o setor de congelados e o setor de produtos
frescos. Esses sensores publicam regularmente as leituras de temperatura em
tópicos específicos do MQTT. O sistema central deve se inscrever nesses tópicos
para coletar os dados e gerar alertas em caso de temperaturas fora dos limites
aceitáveis.

### Ferramentas
- **Mosquitto**: Como broker MQTT.
- **Paho**: Para a implementação do cliente MQTT em Python.
- **SQLite**: Para armazenamento local dos dados de temperatura e alertas.

### Requisitos
1. Os sensores devem publicar leituras de temperatura em um tópico MQTT no formato `loja/sensor/temperatura`.
2. O sistema central deve se inscrever nos tópicos de temperatura de todas as lojas e registrar as leituras em um banco de dados local.
3. O sistema central deve verificar se a temperatura está dentro do intervalo aceitável para cada setor. Caso contrário, deve gerar um alerta e armazená-lo no banco de dados.

### Tarefa
Escreva testes automatizados para validar os requisitos acima antes de implementar o código da aplicação.

### Passos Sugeridos
1. **Configuração do Ambiente**:
   - Instale o Mosquitto Broker e configure-o para aceitar conexões.
   - Instale as bibliotecas Paho MQTT e SQLite para Python.

2. **Escrita dos Testes**:
   - Utilize um framework de teste, como o `unittest` ou `pytest`, para escrever os testes.
   - Escreva testes para verificar se os sensores podem publicar leituras de temperatura no tópico correto.
   - Escreva testes para verificar se o sistema central pode se inscrever nos tópicos de temperatura e registrar os dados no banco de dados.
   - Escreva testes para verificar se o sistema central pode detectar temperaturas fora do intervalo aceitável e gerar os alertas correspondentes.

3. **Implementação do Código**:
   - Após escrever os testes, implemente o código da aplicação que passará nos testes. Isso inclui a lógica para publicação de dados pelos sensores, a lógica de inscrição e armazenamento de dados pelo sistema central, e a lógica de verificação de temperaturas e geração de alertas.

4. **Execução dos Testes**:
   - Execute os testes automatizados para garantir que o código implementado atende aos requisitos.

### Considerações Adicionais
- Você pode expandir este projeto incluindo funcionalidades como uma interface web para visualização dos dados de temperatura e alertas, integração com sistemas de notificação para enviar alertas em tempo real, ou até mesmo a implementação de algoritmos de predição para antecipar possíveis problemas com base nas tendências de temperatura.

Esta atividade proporciona um desafio maior e envolve a integração de várias tecnologias e conceitos, tornando-a mais interessante e relevante para um contexto real de desenvolvimento de software.

## 2. Padrão de entrega

:::warning

Esses são os critérios mínimos para que eu considere a atividade como entregue.
Fique atento, pois o não cumprimento de qualquer um desses critérios pode, no
melhor dos casos, gerar um desconto de nota e, no pior deles, invalidar a
atividade.

:::

1. A atividade deve ser feita em um repositório aberto no github. Seu link deve
   ser fornecido no card da adalove;
2. No README do repositório deve ter instruções claras de como instalar e rodar
   o sistema criado, comandos em blocos de código e uma expliação sucinta do
   que fazem;
3. Ainda no README, deve haver um vídeo/imagens demonstrando plenamente o
   funcionamento do sistema criado;

## 3. Padrão de qualidade
