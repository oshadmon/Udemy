# Export Data 

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
```



