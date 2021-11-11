# Random Data 
## Install 
1. In docker-compose uncomment the section for _device-random_, or copy the code snippet bellow
```
  device-random:
    container_name: edgex-device-random
    depends_on:
      - consul
      - data
      - metadata
    environment:
      CLIENTS_COMMAND_HOST: edgex-core-command
      CLIENTS_COREDATA_HOST: edgex-core-data
      CLIENTS_DATA_HOST: edgex-core-data
      CLIENTS_METADATA_HOST: edgex-core-metadata
      CLIENTS_NOTIFICATIONS_HOST: edgex-support-notifications
      CLIENTS_RULESENGINE_HOST: edgex-kuiper
      CLIENTS_SCHEDULER_HOST: edgex-support-scheduler
      CLIENTS_VIRTUALDEVICE_HOST: edgex-device-random
      DATABASES_PRIMARY_HOST: edgex-redis
      EDGEX_SECURITY_SECRET_STORE: "false"
      REGISTRY_HOST: edgex-core-consul
      Service_Host: edgex-device-random
    hostname: edgex-device-random
    image: edgexfoundry/docker-device-random-go:1.3.0
    networks:
      edgex-network: {}
    ports:
    - 127.0.0.1:49988:49988/tcp
``` 

2. Enable Random Data by rerunning `docker-compose.yml` 
```
docker-compose pull 
docker-compose up -d device-random
```

## Sample Commands

* Device Service
```
curl http://localhost:48081/api/v1/deviceservice/name/device-random 2> /dev/null | jq 
```

* Device Profile
```
curl http://localhost:48081/api/v1/deviceprofile/name/Random-Integer-Generator 2> /dev/null | jq 
```

* Device created
```
curl http://localhost:48081/api/v1/device/name/Random-Integer-Generator01 2> /dev/null | jq 	
```

* Events created
```
# Get all events 
curl http://localhost:48080/api/v1/event 2> /dev/null | jq 

# this allows for a limit in the number of rows returned 
curl http://localhost:48080/api/v1/event/device/Random-Integer-Generator01/10 2> /dev/null | jq 
```

* Get readings from device(s)
```
curl http://127.0.0.1:48080/api/v1/reading 2> /dev/null| jq 

# Sample Output
  {
    "id": "ff7d9f92-4fa0-4eb5-be8b-0a2e8e61ca3c",
    "created": 1616002344325,
    "origin": 1616002344324098300,
    "device": "Random-Integer-Generator01",
    "name": "RandomValue_Int32",
    "value": "1052517020",
    "valueType": "Int32"
  }
```

