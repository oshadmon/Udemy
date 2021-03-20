# Export Data 

0. Validate EdgeX is running with data coming 

1. Create a stream to consume data from EdgeX message bus.
```
curl -X POST \
  http://127.0.0.1:48075/streams \
  -H 'Content-Type: application/json' \
  -d '{
  "sql": "create stream demo() WITH (FORMAT=\"JSON\", TYPE=\"edgex\")"
}'
```

2. Validate stream was created 
```
curl -X GET  http://127.0.0.1:48075/streams 
```

## MQTT Process 
3. Create a rule to send result data to an [MQTT broker](https://github.com/emqx/kuiper/blob/master/docs/en_US/rules/sinks/mqtt.md)  
```
curl -X POST \
  http://127.0.0.1:48075/rules \
  -H 'Content-Type: application/json' \
  -d '{
  "id": "mqtt-rule",
  "sql": "SELECT * FROM demo",
  "actions": [
    {
      "mqtt": {
        "server": "driver.cloudmqtt.com:18785",
        "username": "ibglowct",
        "password": "MSY4e009J7ts", 
        "topic": "result"
      }
    },
    {
      "log":{}
    }
  ]
}'

```

4. Validae rule was created 
```
curl http://127.0.0.1:48075/rules
```

## REST process 
5. Create a rule to send via [REST](https://github.com/emqx/kuiper/blob/master/docs/en_US/rules/sinks/rest.md) 
```
curl -X POST \
  http://127.0.0.1:48075/rules \
  -H 'Content-Type: application/json' \
  -d '{
  "id": "rest-rule1",
  "sql": "SELECT * FROM demo",
  "actions": [
    {
      "rest": {
        "method": "put",
        "url": "24.23.250.144:7849",
        "headers": {"type": "json", "dbms": "edgex", "table": "random_integer_generator01", "mode": "streaming", "Content-Type": "text/plain"}
      }
    },
    {
      "log":{}
    }
  ]
}'
```
