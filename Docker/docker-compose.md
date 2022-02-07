# Docker Compose 
* CLI tool (`docker-compose`)

## YAML File 
* YAML file - describe the package of containers
  * [versions](https://docs.docker.com/compose/compose-file/compose-versioning/) - the version of `docker-compose` to use  
  * containers 
  * networks
  * volumes
  
**Command**: `docker-compose up -d -f ${FILE_NAME}`

**Sample YAML file**:
```yaml
version: '3.1'
services: # containers (same as docker run)
  service-name: # friendly name
    image: # image to use 
    command: # Alternative command to the default cmd specific by the image
    environment: # -e params
    volumes: # -v params
volumes: # docker volume create 
networks: # docker network create 
```

## CLI
  * Start docker-compose: `docker-compse up`
  * Stop & clean: `docker-compose down`
  * list containers: `docker-compose ps`
  * list services: `docker-compose top`

## Volumes
[docker-compose.yml](docker-compose.yml) provides an example of working volumes  

## Build from dockerfile
```yaml
service: 
  service-name: 
    build:
      context: # path
      dockerfile: # dockerfile 
    image:
...
```