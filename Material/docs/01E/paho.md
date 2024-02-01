---
title: Introdução ao Eclipse Paho
sidebar_position: 3
sidebar_class_name: autoestudo
slug: /paho
---

# Introdução ao Eclipse Paho

O Eclipse Paho é um projeto de código aberto que fornece implementações de
clientes para as mensagens e protocolos de comunicação de Machine-to-Machine
(M2M) e Internet of Things (IoT). Uma das principais características do Paho é
a implementação de clientes para o protocolo MQTT (Message Queuing Telemetry
Transport), que é um protocolo de mensagens leve e eficiente, ideal para
ambientes de IoT devido ao seu baixo uso de largura de banda e eficiência
energética.

O MQTT é um protocolo de publicação/assinatura que permite a comunicação entre
dispositivos com baixa largura de banda e em redes que podem experimentar
níveis variáveis de latência devido à natureza instável das conexões de rede,
como é comum em ambientes de IoT. O Eclipse Paho fornece bibliotecas que
permitem aos desenvolvedores integrar facilmente a funcionalidade MQTT em
aplicações de software em uma variedade de linguagens de programação, incluindo
Java, C, Python, JavaScript, Go, entre outras.

Além disso, o Eclipse Paho também oferece suporte a outros protocolos
relacionados à IoT, como o MQTT-SN, que é uma variação do MQTT para redes
sensíveis (como ZigBee). A ênfase do Paho está na criação de clientes que são
leves, eficientes e fáceis de usar, tornando-o uma escolha popular entre
desenvolvedores que trabalham em projetos de IoT.

O projeto Paho é parte da Eclipse IoT Working Group, uma comunidade
colaborativa de empresas e indivíduos que trabalham juntos para desenvolver
plataformas e frameworks de código aberto para IoT. O envolvimento da Eclipse
Foundation assegura que o Paho continue a ser desenvolvido e mantido
ativamente, seguindo padrões abertos e sendo acessível a todos os
desenvolvedores interessados em soluções de IoT.

## Configurando um Broker MQTT Local com Mosquitto

Antes de utilizar as bibliotecas do Eclipse Paho em Python ou Go, é necessário configurar um broker MQTT. O Mosquitto é uma opção popular devido à sua leveza e facilidade de uso.

### Instalação do Mosquitto

1. **Windows**: Baixe o instalador do [site oficial do Mosquitto](https://mosquitto.org/download/) e siga as instruções de instalação.
2. **Linux**: Use o gerenciador de pacotes da sua distribuição. Por exemplo, no Ubuntu, execute `sudo apt-get install mosquitto mosquitto-clients`.
3. **macOS**: Use o Homebrew, executando `brew install mosquitto`.

### Configuração Básica

Após a instalação, o Mosquitto pode ser iniciado com a configuração padrão. No entanto, é recomendável criar um arquivo de configuração básico.

1. Crie um arquivo chamado `mosquitto.conf` e adicione as seguintes linhas para uma configuração básica:

   ```
   listener 1883
   allow_anonymous true
   ```

2. Inicie o Mosquitto com este arquivo de configuração: `mosquitto -c mosquitto.conf`.

## Utilizando o Eclipse Paho em Python

Para usar o Paho em Python, primeiro instale a biblioteca e, em seguida, escreva um script simples para publicar e assinar mensagens.

### Instalação

Instale a biblioteca Paho MQTT para Python:

```bash
pip install paho-mqtt
```

### Exemplo de Código

```python
import paho.mqtt.client as mqtt

# Callback quando uma mensagem é recebida do servidor.
def on_message(client, userdata, message):
    print(f"Recebido: {message.payload.decode()} no tópico {message.topic}")

# Configuração do cliente
client = mqtt.Client("python_client")
client.on_message = on_message

# Conecte ao broker
client.connect("localhost", 1883, 60)

# Inscreva no tópico
client.subscribe("test/topic")

# Publicando uma mensagem
client.publish("test/topic", "Hello MQTT")

# Loop para manter o cliente executando
client.loop_forever()
```

## Utilizando o Eclipse Paho em Go

Para usar o Paho em Go, instale a biblioteca e escreva um exemplo de código para interagir com o MQTT.

### Instalação

```bash
go get github.com/eclipse/paho.mqtt.golang
```

### Exemplo de Código

```go
package main

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "time"
)

// Função de callback para quando receber uma mensagem
var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("Recebido: %s do tópico: %s\n", msg.Payload(), msg.Topic())
}

func main() {
    opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
    opts.SetClientID("go_client")
    opts.SetDefaultPublishHandler(messagePubHandler)

    // Cria e inicia um cliente MQTT
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    // Inscreve no tópico
    if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        return
    }

    // Publica uma mensagem
    token := client.Publish("test/topic", 0, false, "Hello MQTT")
    token.Wait()

    time.Sleep(3 * time.Second)

    // Desconecta o cliente
    client.Disconnect(250)
}
```

Esses exemplos são básicos e destinam-se a demonstrar a publicação e assinatura
de mensagens usando o MQTT com o Eclipse Paho em Python e Go. Para aplicações
mais complexas, você deve explorar funcionalidades adicionais como qualidade de
serviço (QoS), retenção de mensagens e autenticação segura.
