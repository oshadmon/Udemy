# Prerequisite
*  [Install Docker & Docker-compose](../Docker/install_docker.sh)

# Download EdgeX Docker Image 
1. Create directory EdgeX (`mkdir $HOME/edgex ; cd $HOME/edgex`) - All commands will be within this directory unless stated otherwise 

2. Download EdgeX file 
   * [x86](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-hanoi-no-secty.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-hanoi-no-secty.yml -o docker-compose.yml
```
   * [ARM](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-hanoi-no-secty-arm64.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-hanoi-no-secty-arm64.yml -o docker-compose.yml
```

3. Update docker-compose.yml to contain machine IP rather than 127.0.0.1 
```
sed -i 's/127.0.0.1/${IP}/g' docker-comppse.yaml
```

# Start / Stop Docker Image 
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

