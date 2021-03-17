URL: https://jonamiki.com/wp-content/uploads/2020/08/EdgeX-Foundry-tutorial-ver1.1.pdf

# SetUp & Teardown process 
0. create a directory from which work is done unless stated otherwise 
```
mkdir edgex 
cd edgex 
``` 
 
1. Download EdgeX 
```
# download 
wget https://raw.githubusercontent.com/jonas-werner/EdgeX_Tutorial/master/docker-compose_files/docker-compose_step1.yml

# create a copy which will be using 
cp docker-compose_step1.yml docker-compose.yml 
```

2. Start EdgeX as is 
```
# download packages 
docker-compose pull 

# start EdgeX 
docker-compose up -d 
```
3. Validate EdgeX is running 
   * Validate services through browser - http://${MACHINE_IP}:8500/ui/dc1/services 
   * Validate services through docker 
```
# check docker images 
docker image ls 

# check docker containers 
docker ps -a 
```
   * valide services through cURL 
```
# list of connected device(s) 
curl http://localhost:48082/api/v1/device
```

4. Stop EdgeX 
```
# Stop & remove docker containers
docker-compose down

# Remove volumes whe removing containers 
docker-comopose down -v 
```

Next Section [Adding GUI](Lab_GUI.md) 

5. Update EdgeX to have portainer & EdgeX UI 
   * [EdgeX UI](http://139.162.205.95:4000/) required port 4000 to open 
   * [Portainer](http://139.162.205.95:9000/#/home) required port 9000 to open

````
# In "Services" seciton add the following lines for a EdgeX UI to be enabled 
# If asked, the default username/password is admin 
ui:
  container_name: edgex-ui-go
  hostname: edgex-ui-go
  image: nexus3.edgexfoundry.org:10004/docker-edgex-ui-go:master
  networks:
    edgex-network: null
  ports:
    - "0.0.0.0:4000:4000/tcp"
  read_only: true

# In "Services" section add the following lines for Portainer to be enabled 
portainer:
  image: portainer/portainer
  ports:
    - "0.0.0.0:9000:9000"
  container_name: portainer
  command: -H unix:///var/run/docker.sock
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock:z
    - portainer_data:/data

# For Portainer UI user needs to also add a volume 
volumes:
  db-data:
  log-data:
  consul-config:
  consul-data:
  portainer_data: # <-- add 
```
 
6. Start EdgeX

7. 
