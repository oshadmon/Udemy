The following is a watered down version of the offical [EdgeX documentation](https://docs.edgexfoundry.org/1.2/getting-started/quick-start/). It covers
1. Install process 
2. Start/Stop/clean  
3. Add a device (Random) & control it 
4. Send data via MQTT

# Prerequisite
*  [Install Docker & Docker-compose](https://github.com/AnyLog-co/AnyLog-Network/blob/develop/scripts/install/install_docker.sh) 

# Download EdgeX Docker Image 
https://docs.edgexfoundry.org/1.2/getting-started/quick-start/
0. Create directory EdgeX (`mkdir $HOME/edgex ; cd $HOME/edgex`) - All commands will be within this directory unless stated otherwise 
1. Download EdgeX file 
   * [x86](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty.yml -o docker-compose.yml
```
   * [ARM](https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty-arm64.yml) 
```
curl https://raw.githubusercontent.com/edgexfoundry/developer-scripts/master/releases/geneva/compose-files/docker-compose-geneva-redis-no-secty-arm64.yml -o docker-compose.yml
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
