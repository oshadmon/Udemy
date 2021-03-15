# MQTT 

The original documentation suggest using [HiveMQTT](https://www.hivemq.com/) rather than [CloudMQTT](https://www.cloudmqtt.com/). 

## Install 

1. Uncomment MQTT section 
```
  app-service-mqtt:
    image: edgexfoundry/docker-app-service-configurable:1.1.0
    ports:
      - "127.0.0.1:48101:48101"
    container_name: edgex-app-service-configurable-mqtt
    hostname: edgex-app-service-configurable-mqtt
    networks:
      - edgex-network
    environment:
      <<: *common-variables
      edgex_profile: mqtt-export
      Service_Host: edgex-app-service-configurable-mqtt
      Service_Port: 48101
      MessageBus_SubscribeHost_Host: edgex-core-data
      Binding_PublishTopic: events
      Writable_Pipeline_Functions_MQTTSend_Addressable_Address: broker.mqttdashboard.com
      Writable_Pipeline_Functions_MQTTSend_Addressable_Port: 1883
      Writable_Pipeline_Functions_MQTTSend_Addressable_Protocol: tcp
      Writable_Pipeline_Functions_MQTTSend_Addressable_Publisher: edgex
      Writable_Pipeline_Functions_MQTTSend_Addressable_Topic: EdgeXEvents
    depends_on:
      - consul
      - data
```

2. Update MQTT information
   * **MQTT Address**: `broker.mqttdashboard.com` --> `driver.cloudmqtt.com`
   * **MQTT Port**: 1883 --> 18896
   * **MQTT Publisher**: `edgex` --> `kqypbmie`

3. Start MQTT 
```
docker-compose up -d
```

