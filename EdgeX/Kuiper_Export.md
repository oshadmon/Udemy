# Export Data 


0. Validate EdgeX is running with data coming 
```
curl http://127.0.0.1:48080/api/v1/reading | jq
```

1. Create a [stream](https://github.com/emqx/kuiper/blob/master/docs/en_US/streams.md) to consume data from EdgeX message bus.
```
curl -X POST \
  http://127.0.0.1:48075/streams \
  -H 'Content-Type: application/json' \
  -d '{"sql": "create stream demo() WITH (FORMAT=\"JSON\", TYPE=\"edgex\")"}'
```

2. Validate stream was created 
```
curl -X GET  http://127.0.0.1:48075/streams/${STREAM-ID} 
```
**Note**: Process needs to occur only once, per stream (ie more than one rule can exist per stream) 

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

4. Validate rule was created 
```
curl http://127.0.0.1:48075/rules/${RULE-ID}
```

## REST process 
3. Create a rule to send via [REST](https://github.com/emqx/kuiper/blob/master/docs/en_US/rules/sinks/rest.md) 
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
        "url": "http://139.162.205.95:2249",
        "headers": {"type": "json", "dbms": "edgex", "table": "random_integer_generator01", "mode": "streaming", "Content-Type": "text/plain"}
      }
    },
    {
      "log":{}
    }
  ]
}'
```

4. Validate rule was created 
```
curl http://127.0.0.1:48075/rules/${RULE-ID}
```


## Remove 
* To Remove stream 
```
curl -X DELETE http://127.0.0.1:48075/streams/${STREAM-ID}
```

* To Remove a rule
```
curl -X DELETE http://127.0.0.1:48075/rules/{RULE-ID}
```

