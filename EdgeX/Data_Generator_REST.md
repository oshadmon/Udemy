# Random Data 
## Install 
1. Add [device-random](https://docs.edgexfoundry.org/1.0/examples/Ch-ExamplesRandomDeviceService/) secition to docker-compose.yaml
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

* Events (values) created by the random integer generator device
```
curl http://localhost:48080/api/v1/event 2> /dev/null | jq 

# this allows for a limit in the number of rows returned 
curl http://localhost:48080/api/v1/event/device/Random-Integer-Generator01/10 2> /dev/null | jq 
```

