# Features

## Overlay & Multi-Host Networking
* when deploying add `--driver overlay`
* Intra-swarm communcation
* IPSec encrytion between nodes 
* Can be either on one or multiple overlay networks

## Routing Mesh
* routes ingress packets for service to proper tasks 
* instead of load-balancer (aka stateless load balancer)  
* Spans to all nodes 
* shouldn't care which servers you're on
* Layer 3 (IP / Port layer)
  * to run Layer 4 use Nginx or HAProxy

## Multi-Service App
[Distributed Voting App](https://github.com/dockersamples/example-voting-app)
* 1 volume 
* 2 networks 
* 5 services
  * Web front & backend 
  * worker process 
  * 2 databases 
    * Redis 
    * Postgres


## Stacks & Production Grade Compose 
* YAML must be version "3" or higher 
* YAML file can contain **both** content for _docker-compose_ and _Swarm_
```yaml
services:
  service-name:
    deploy:
      replicats:
      update_config: 
        parallelism: 
        delay: 
      restart_policy:
        condition: on-failure
```

## Secrets 
