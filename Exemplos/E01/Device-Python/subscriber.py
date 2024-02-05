import paho.mqtt.client as mqtt
import json

# Callback quando uma mensagem é recebida do servidor.
def on_message(client, userdata, message):
    sensor_data = json.loads(message.payload)
    print(f"Dados do Sensor Recebidos: {sensor_data}")

# Callback para quando o cliente recebe uma resposta CONNACK do servidor.
def on_connect(client, userdata, flags, rc):
    print("Conectado com código de resultado "+str(rc))
    client.subscribe("sensor/data")

# Configuração do cliente
client = mqtt.Client("python_subscriber")
client.on_connect = on_connect
client.on_message = on_message

# Conecte ao broker
client.connect("localhost", 1891, 60)

# Loop para manter o cliente executando e escutando por mensagens
client.loop_forever()
