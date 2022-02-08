# Networking

* Each container is its own private network 
* If all dockers are on the smae network there is no need to open special ports. ie just like there wouldn't be a need to open special ports between different machines on the same network. 
* Docker drivers extends docker to use more components within the machine 

## Commands
* `-p` or `--port` - specify port to open for a docker instance
* `docker container inspect --format '{{ .NetworkSettings.IPAddress }}' <${container_name}>` - Get IP for docker instance 
* `docker container port <${container_name}>  ` - get port for docker instance 
* `--network` to specify network type
    * **brdige** or `Docker0` - the default network that bridges through the NAT to the physical network. This assigns the “default” IP address for the container  
    * **host** - a special network that skips docker networking, and connects directly to the host. This prevents the security boundaries of containerization, but can improve network performance  
    * **none** - removes ehto0 and only leaves you with local host
* Create Network `docker network create <${network_name}>  --driver <${network_type}>`
* Inspect network `docker network inspect <${network_name}>`
* Attach network to container `docker network connect <${network_name}> <${container_name}>`
*  Disconnect `docker network disconnect <${network_name}> <${container_name}>`

## DNS & Containers 
* Docker has a built-in DNS  server that containers use by default 
* host/container name isn’t necessarily consistent whereas IP is
* `--link` command enables DNS 
* Alternativly users can create internal networks to connect between dockers on separate network configurations	

## Install Ping 
[iputils-ping](https://packages.debian.org/buster/iputils-ping) sends ICMP ECHO_RREQUEST packets to host in order to test if host is reachable 
```
# Start Ubuntu docker instance in exectuable 
docker container exec -it ubuntu bash

# Within Ubuntu
# update env 
apt-get -y update ; apt-get -y upgrade ; apt-get updpate 

# Install iputils
apt-get install -y iputils-ping  
```

## Round-Robin
A round-robin using elastic search v.2 docker instances
```
# Create Network 
docker network create ${network_name}

# Start 2 elastic search instances 
docker container run -d --net round_robin --net-alias search elasticsearch:2 

# find DNS 
docker container run --rm --net new-network alpine nslookup search.

# Execute round-robin using CentOS against elastic containers
docker container run --rm --net new-network centos curl -s search:9200
```
