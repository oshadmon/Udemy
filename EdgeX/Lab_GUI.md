# EdgX GUI Options
## EdgexUI 
--> **Link**: http://127.0.0.1:4000
--> **About**: EdgeX UI written in Golang, added as a way to view device services and other information via a web browser rather than the command line.

## Portainer 
--> **Link**: http://127.0.0.1:9000


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

