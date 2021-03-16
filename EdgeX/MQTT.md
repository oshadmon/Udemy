# MQTT 

## HiveMQ  process install 

1. In docker-compose.yml add `app-service-mqtt` bellow `rulesengine` (roughly line 255)  
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

2. Start MQTT 
```
docker-compose up -d
```

3. Visit [HiveMQTT](http://hivemq.com/demos/websocket-client/) & "connect" 

4. Set _Add a new Topic Subscription_ to `EdgeXEvents` & data should appear in _Messages_ section. 

## CloudMQTT 

Process Link: https://fuji-docs.edgexfoundry.org/Ch-APIExportClientRegistration-MQTT.html

1. Create a new [CloudMQTT](https://cloudmqtt.com/) Topic

2. Within the new topic access _Users & ACL_

3. Add new user
   --> **Username**: exportpublisher
   --> **Password**: edgex4AnyLog!

4. Add a topic & access rights (_ACL_ section)
   --> topic 
   --> **Pattern**: edgexdatatopic

5. Add `export-client` to docker-compose.yml 
Note, I was unable to run the process only with [EdgeX v1.1 Fuji](https://github.com/edgexfoundry/developer-scripts/blob/master/releases/fuji/compose-files/docker-compose-fuji-no-secty.yml) 
```
  export-client:
    image: edgexfoundry/docker-export-client-go:1.1.0
    ports:
      - "48071:48071"
    container_name: edgex-export-client
    hostname: edgex-export-client
    networks:
      - edgex-network
    environment:
      <<: *common-variables
    volumes:
      - db-data:/data/db
      - log-data:/edgex/logs
      - consul-config:/consul/config
      - consul-data:/consul/data
    depends_on:
      - data
```

6. Start `export-client` process
```
docker-compose up -d export-client
```

7. 
