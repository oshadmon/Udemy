# Random Data 
## Install 
1. Within doocker-compose.yaml remove the section within [device-random](https://docs.edgexfoundry.org/1.2/examples/Ch-ExamplesRandomDeviceService/) 
```
device-random:
    image: edgexfoundry/docker-device-random-go:1.2.1
    ports:
      - "127.0.0.1:49988:49988"
    container_name: edgex-device-random
    hostname: edgex-device-random
    networks:
      - edgex-network
    environment:
      <<: *common-variables
     Service_Host: edgex-device-random
    depends_on:
      - data
      - command
``` 

2. Enable Random Data by rerunning `docker-compose.yml` 
```
docker-compose up -d device-random
```

## Sample Commands

* Device Service
```
curl http://localhost:48081/api/v1/deviceservice/name/device-random 2> /dev/null | jq 
```

* Device Profile
```
curl http://localhost:48081/api/v1/deviceprofile/name/Random-Integer-Generator
```

* Device created
```
curl http://localhost:48081/api/v1/device/name/Random-Integer-Generator01	
```

* Events (values) created by the random integer generator device
```
curl http://localhost:48080/api/v1/event 2> /dev/null | jq 

# this allows for a limit in the number of rows returned 
curl http://localhost:48080/api/v1/event/device/Random-Integer-Generator01/10 | jq 
```

