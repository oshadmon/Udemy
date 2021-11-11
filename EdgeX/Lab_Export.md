# Export Data 
## HiveMQTT 
1. In [docker-compose.yml](docker-compose_step2.yml) uncomment _app-service-mqtt_ 
```
  app-service-mqtt:
      image: edgexfoundry/docker-app-service-configurable:1.1.0
      ports:
        - "0.0.0.0:48101:48101"
      container_name: edgex-app-service-configurable-mqtt
      hostname: edgex-app-service-configurable-mqtt
      networks:
        edgex-network:
          aliases:
            - edgex-app-service-configurable-mqtt
      environment:
        <<: *common-variables
        edgex_profile: mqtt-export
        Service_Host: edgex-app-service-configurable-mqtt
        Service_Port: 48101
        MessageBus_SubscribeHost_Host: edgex-core-data
        Binding_PublishTopic: events
        # Added for MQTT export using app service
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_ADDRESS: broker.hivemq.com
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PORT: 1883
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PROTOCOL: tcp
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_TOPIC: "YOUR-UNIQUE-TOPIC"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_AUTORECONNECT: "true"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_RETAIN: "true"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_PERSISTONERROR: "false"
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PUBLISHER: 
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_USER: 
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PASSWORD:
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_QOS: ["your quality or service"]
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_KEY: [your Key]  
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_CERT: [your Certificate]

      depends_on:
        - consul
  #      - logging  # uncomment if re-enabled remote logging
        - data
```
**Note**: Update _"YOUR-UNIQUE-TOPIC"_ with unique topic 

2. In [docker-compose.yml](docker-compose_step2.yml) uncomment _rulesengine_
```
  rulesengine:
    image: emqx/kuiper:0.4.2-alpine
    ports:
      - "0.0.0.0:48075:48075"
      - "0.0.0.0:20498:20498"
    container_name: edgex-kuiper
    hostname: edgex-kuiper
    networks:
      - edgex-network
    environment:
      # KUIPER_DEBUG: "true"
      KUIPER_CONSOLE_LOG: "true"
      KUIPER_REST_PORT: 48075
      EDGEX_SERVER: edgex-app-service-configurable-rules
      EDGEX_SERVICE_SERVER: http://edgex-core-data:48080
      EDGEX_TOPIC: events
      EDGEX_PROTOCOL: tcp
      EDGEX_PORT: 5566
    depends_on:
      - app-service-rules
```

3. Start Docker process 
```
docker-compose pull 
docker-compse up -d 

# If docker is already running, then run this instead: 
docker-compose pull 
docker-compose up -d app-service-rules 
docker-compose up -d app-service-mqtt
```

4. Go to [HiveMQTT](http://www.hivemq.com/demos/websocket-client/) 

5. Press "Connect" with default values

7. "Add New Topic Subscription" with _"YOUR-UNIQUE-TOPIC"_

## CloudMQTT
1. In [CloudMQTT](https://www.cloudmqtt.com/docs/index.html#:~:text=Create%20a%20CloudMQTT%20instance%20Create%20an%20account%20and,plan%20you%20require%20depends%20on%20your%20use%20case) create a new instance

2. In [docker-compose.yml](docker-compose_step2.yml) uncomment _app-service-mqtt_ 
```
  app-service-mqtt:
      image: edgexfoundry/docker-app-service-configurable:1.1.0
      ports:
        - "0.0.0.0:48101:48101"
      container_name: edgex-app-service-configurable-mqtt
      hostname: edgex-app-service-configurable-mqtt
      networks:
        edgex-network:
          aliases:
            - edgex-app-service-configurable-mqtt
      environment:
        <<: *common-variables
        edgex_profile: mqtt-export
        Service_Host: edgex-app-service-configurable-mqtt
        Service_Port: 48101
        MessageBus_SubscribeHost_Host: edgex-core-data
        Binding_PublishTopic: events
        # Added for MQTT export using app service
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_ADDRESS: "YOUR-UNIQUE-BROKER-URL" 
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PORT: "YOUR-UNIQUE-BROKER-PORT" 
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PROTOCOL: tcp
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_TOPIC: "YOUR-UNIQUE-TOPIC"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_AUTORECONNECT: "true"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_RETAIN: "true"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_PERSISTONERROR: "false"
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PUBLISHER: 
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_USER: "YOUR-UNIQUE-USER"
        WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_ADDRESSABLE_PASSWORD: "YOUR-UNIQUE-PASSWORD"
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_QOS: ["your quality or service"]
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_KEY: [your Key]  
        # WRITABLE_PIPELINE_FUNCTIONS_MQTTSEND_PARAMETERS_CERT: [your Certificate]

      depends_on:
        - consul
  #      - logging  # uncomment if re-enabled remote logging
        - data
```
**Note**: Update _"YOUR-UNIQUE-?"_ with correct information 

3. In [docker-compose.yml](docker-compose_step2.yml) uncomment _rulesengine_
```
  rulesengine:
    image: emqx/kuiper:0.4.2-alpine
    ports:
      - "0.0.0.0:48075:48075"
      - "0.0.0.0:20498:20498"
    container_name: edgex-kuiper
    hostname: edgex-kuiper
    networks:
      - edgex-network
    environment:
      # KUIPER_DEBUG: "true"
      KUIPER_CONSOLE_LOG: "true"
      KUIPER_REST_PORT: 48075
      EDGEX_SERVER: edgex-app-service-configurable-rules
      EDGEX_SERVICE_SERVER: http://edgex-core-data:48080
      EDGEX_TOPIC: events
      EDGEX_PROTOCOL: tcp
      EDGEX_PORT: 5566
    depends_on:
      - app-service-rules
```

4. Start Docker process 
```
docker-compose pull 
docker-compse up -d 

# If docker is already running, then run this instead: 
docker-compose pull 
docker-compose up -d app-service-rules 
docker-compose up -d app-service-mqtt
```
5. In the _Websocket UI_ of the [CloudMQTT Instance](https://www.cloudmqtt.com/) you should begin seeing data coming in

