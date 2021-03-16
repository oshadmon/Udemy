# Setup & Install 

## Prerequisite
* [Install Docker & Docker-compose](../Docker/install_docker.sh)
* [Linux JQ](https://www.educba.com/linux-jq/) 
```
sudo apt-get -y install jq
```

## Download EdgeX Docker Image 
1. Create directory EdgeX (`mkdir $HOME/edgex ; cd $HOME/edgex`) - All commands will be within this directory unless stated otherwise 

2. Download EdgeX file 
  * [x86](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/hanoi/compose-files/docker-compose-hanoi-no-secty.yml)
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/hanoi/compose-files/docker-compose-hanoi-no-secty.yml -o docker-compose.yml 
```
  * [ARM](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/hanoi/compose-files/docker-compose-hanoi-no-secty-arm64.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/hanoi/compose-files/docker-compose-hanoi-no-secty-arm64.yml -o docker-compose.yml
```

3. Enable [device-random](Data_Generator_REST.md) to generate data 

## Start / Stop Docker Image 
1. Start docker-compose - redirects output to file rather than screen  
```
docker-compose pull 
docker-compose up > output.err 2>&1 & 
```

2. Validate EdgeX is running
   a. EdgeX docker containers are running 
```
docker ps -a 
```
   b. Bssic Event count to check you can access Query EdgeX 
```
curl http://localhost:48080/api/v1/event/count
``` 

3. Stop docker-compose
```
docker-compose stop
```

4. Remove (all) deplyoments from docker 
```
docker rm `docker ps -a | grep -v CONTAINER | awk -F " " '{print $1}'`
```

5. Remove volume
```
docker volume prune 
```

## Notes 
* By default EdgeX is configured to support local curl (ie `localhost` or `127.0.0.1`); however, by replacing `127.0.0.1` with the physical IP commands can be done remotely. Note, when doing so, you need to restart EdgeX & use the physical IP rather than `localhost` as the examples show. 
```
sed -i 's/127.0.0.1/${IP}/g' docker-comppse.yaml
```
* Be advised that depending on firewall configs, certain pots may need to be opened.

