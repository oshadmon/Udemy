# EdgeX GUI Options

## EdgexUI 
   * **Link**: http://127.0.0.1:4000
   * **About**: EdgeX UI written in Golang, added as a way to view device services and other information via a web browser rather than the command line.

### Steps
1. In "Services" seciton of the docker-compose.yml add the following lines for a EdgeX UI to be enabled  If asked, the default username/password is admin 
```
ui:
  container_name: edgex-ui-go
  hostname: edgex-ui-go
  image: nexus3.edgexfoundry.org:10004/docker-edgex-ui-go:master
  networks:
    edgex-network: null
  ports:
    - "0.0.0.0:4000:4000/tcp"
  read_only: true
```

2. Rerun docker start process 
```
docker-compose pull 
docker-compose up -d 
```

## Portainer 
   * **Link**: http://127.0.0.1:9000
   * **About**: Portainer is way to view and interact with the containers in a graphical manner

### Steps
1. In _Services_ section of the docker-compose.yml add the following lines for Portainer to be enabled 
```
portainer:
  image: portainer/portainer
  ports:
    - "0.0.0.0:9000:9000"
  container_name: portainer
  command: -H unix:///var/run/docker.sock
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock:z
    - portainer_data:/data
```

2. For Portainer UI user needs to also add a volume 
```
volumes:
  db-data:
  log-data:
  consul-config:
  consul-data:
  portainer_data: # <-- add 
```

3. Rerun docker start process 
```
docker-compose pull 
docker-compose up -d 
```
**Note**:  If you'd like, you could add both sets of UIs in one process

