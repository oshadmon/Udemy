# Kuiper
[EMQ X Kuiper](https://github.com/emqx/kuiper) - SQL based rule engine, integrated with EdgeX. 
Kuiper is an edge lightweight IoT data analytics / streaming software implemented by Golang, and it can be run at all kinds of resource constrained edge devices. 
* **Source**: The data source of streaming data, such as data from MQTT broker. In EdgeX scenario, the data source is EdgeX message bus, which could be ZeroMQ or MQTT broker.
* **SQL**: SQL is where you specify the business logic of streaming data processing. Kuiper provides SQL-like statements to allow you to extract, filter & transform data.
* **Sink**: Sink is used for sending analysis result to a specified target. For example, send analysis result to another MQTT broker, or an HTTP rest address.


## Links
* Project is by [EMQX](https://www.emqx.io/) 
* [EdgeX Rules Engine](https://github.com/emqx/kuiper/blob/master/docs/en_US/edgex/edgex_rule_engine_tutorial.md)  
* [SQL](https://github.com/emqx/kuiper/blob/master/docs/en_US/edgex/edgex_rule_engine_tutorial.md) 
   * [Streaming Specific](https://github.com/emqx/kuiper/blob/master/docs/en_US/sqls/streams.md) 
   * [Language Elements](https://github.com/emqx/kuiper/blob/master/docs/en_US/sqls/query_language_elements.md) 
   * [Windowing](https://github.com/emqx/kuiper/blob/master/docs/en_US/sqls/windows.md) 
   * [Built-in Functions](https://github.com/emqx/kuiper/blob/master/docs/en_US/sqls/built-in_functions.md) 

## Setup 
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
docker-compose
```

3. Validate connection
```
curl http://127.0.0.1:48075 
```

