# Intro to Docker

## General Info 
   * Released in 2013  by dotcloud (original name  of docker inc.)  
   * Containers are serverless and functions as a service

## Why need docker? 
   * Enhances the speed of development
   * Allows you to deploy and use a software in an identical way 
   * Used by most (if not all) major companies 
   * No need to rewrite the application, but rather the way it is packaged 

## Versons 
   * Over 12 versions 
      * Direct 
      * Tool Suite (Win 10  / Mac) -- includes a “VM” 
      * Cloud have vendor specific features 
   * More than a container run time 
   * Docker CE (Community) vs EE (Enterprise) 
      * EE  when needing management GUI and support services (24/7) 
      * Edge (Beta) comes about every month 
      * CE is supported for 1 quarter (~3 months) at which point a new “stable” version is available 

## Install 
   * Docker is native for linux 
   * RHEL CE (yam based) is under Centos 
   * [Docker Machine](https://docs.docker.com/machine/overview/) - Install and managee virtual docker hosts 
   * [Docker Compse](https://docs.docker.com/compose/overview/) - Allows for multiple containers (docker instances) to run on a single node using Dockerfile


```
# Install Docker (Edge)
curl -sSL https://get.docker.com/ | sh

# Configure Permissions to local machine
sudo usermod -aG docker ${user}

# Install Docker Machine
base=https://github.com/docker/machine/releases/download/v0.16.0 &&
curl -L $base/docker-machine-$(uname -s)-$(uname -m) >/tmp/docker-machine &&
sudo mv /tmp/docker-machine /usr/local/bin/docker-machine && 
chmod +x /usr/local/bin/docker-machine

# Install Docker Compose 
curl -L "https://github.com/docker/compose/releases/download/1.26.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# Configure Docker 
sudo groupadd docker
sudo usermod -aG docker ${user}
newgrp docker
```



