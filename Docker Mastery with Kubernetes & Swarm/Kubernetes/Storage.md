# Storage
* Volumes are stateless 
* Different from _Docker_
* StatefulSets: a type of resource that make pods more _sticky_
* Use DB as a service

## Volumes
* Container Storage Interface (CSI) plugin: A standard API for all containers
  * Still relatively new 

**Volumes**:
* Tied to the life cycle of the pod 
* All containers in a single pod can share the same volume(s) 

**Persistent Volume**:
* Created at the cluster level
* Separate storage from pod 
* multiple pods can share them

## Ingress 
* Ingress Controllers - a 3<sup>rd</sup> party proxies that supports _OSI Layer 7_ (_HTTP_)
* Use something like [_Nginx_](https://www.nginx.com/), [_Traefik_](https://traefik.io/) or [_HAProxy_](https://www.haproxy.org/)

