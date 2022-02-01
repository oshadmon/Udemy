# Kubernetes
The following discuses migrating docker-compose into Kubernetes -- Based on [DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-migrate-a-docker-compose-workflow-to-kubernetes#:~:text=1%20Prerequisites.%20...%202%20Step%201%20%E2%80%94%20Installing,Database%20Service%20and%20an%20Application%20Init%20Container.%20)

**Machine Setup**: VirtualBox with Ubuntu 20.04 
* CPU: 4 
* Disk Memory: 4096MB 
* VM Size: 15GB


## Prerequisites
* [Docker](../Docker/docker_io_install.sh)


## Table of Content
* [Installation](installation.md) 
  * [kompose](installation.md#kompose) - conversion tool for Docker Compose to container orchestrators such as Kubernetes
  * [kubectl](installation.md#kubectl) - command line tool lets you control Kubernetes clusters
  * [minikube](installation.md#minikube) - tool that enables you to run Kubernetes on your laptop or other local machine

## Useful Links 
* [Kubernetes Docs](https://kubernetes.io/docs/reference/)
