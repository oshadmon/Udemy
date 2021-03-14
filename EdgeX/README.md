The following is a watered down version of the offical [EdgeX documentation](https://docs.edgexfoundry.org/1.2/getting-started/quick-start/). It covers
1. Install process 
2. Start/Stop/clean  
3. Add a device (Random) & control it 
4. Send data via MQTT
Please note, the links provided direct to version 1.2 of EdgeX. However, the actual documentation do not different (subsentially) between the version. 

# Prerequisite
*  [Install Docker & Docker-compose](https://github.com/AnyLog-co/AnyLog-Network/blob/develop/scripts/install/install_docker.sh) 

# Download & Install 
## Download EdgeX Docker Image 
1. Create directory EdgeX (`mkdir $HOME/edgex ; cd $HOME/edgex`) - All commands will be within this directory unless stated otherwise 

2. Download EdgeX file 
   * [x86](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty.yml -o docker-compose.yml
```
   * [ARM](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty-arm64.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty-arm64.yml -o docker-compose.yml
```

3. Update docker-compose.yml to contain machine IP rather than 127.0.0.1 
```
sed -i 's/127.0.0.1/139.162.205.95/g' docker-comppse.yaml
```

## Start / Stop Docker Image 
1. Start docker-compose - redirects output to file rather than screen  
```
docker-compose up > output.err 2>&1 & 
```

2. Stop docker-compose
```
docker-compose stop
```

3. Remove (all) deplyoments from docker 
```
docker rm `docker ps -a | grep -v CONTAINER | awk -F " " '{print $1}'`
```

4. Remove volume
```
docker volume prune 
```

# Basic Processes 
## Random Data 
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

3. View data 
```
# By replacing N-rows the command shows the first N rows 
curl http://localhost:48080/api/v1/event/device/Random-Integer-Generator01/100

# Example
ubuntu@edgex:~/edgex$ curl http://139.162.205.95:48080/api/v1/event/device/Random-Integer-Generator01/1 | jq 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   348  100   348    0     0  31636      0 --:--:-- --:--:-- --:--:-- 34800
[
  {
    "id": "bb5ef69a-4eb4-483c-a493-cf399864db89",
    "device": "Random-Integer-Generator01",
    "created": 1615683743224,
    "origin": 1615683743223265300,
    "readings": [
      {
        "id": "f530c474-2b8c-4936-8031-c4e43b7fa41f",
        "created": 1615683743224,
        "origin": 1615683743223221500,
        "device": "Random-Integer-Generator01",
        "name": "RandomValue_Int8",
        "value": "122",
        "valueType": "Int8"
      }
    ]
  }
]
```
 
4. Sending data 
