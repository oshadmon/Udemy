# Setup 
1. Make sure you have a docker-compose with a _rulesengine_ containing emqx/kuiper image
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

2. Start docker-compose
```
docker-compose up -d
```

3. Validate connection
```
curl http://127.0.0.1:48075 
```

