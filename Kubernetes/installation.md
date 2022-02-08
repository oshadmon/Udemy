# Install 

Sample installation was done on _Ubuntu 20.04_. The shared links provide information for other operating systems. 

## Docker 
Install [docker](../Docker Mastery with Kubernetes & Swarm)
```bash
sudo apt-get -y update 
sudo apt-get -y install docker.io docker-compose 

# Permissions 
USER=`whoami`
sudo groupadd docker
sudo usermod -aG docker ${USER}
newgrp docker

sudo apt-get -y update

# validate
docker --version
<< COMMENT
Docker version 20.10.7, build 20.10.7-0ubuntu5~20.04.2
COMMENT  
```

## kompose
Based on [kompose](https://kompose.io/installation/)
```bash
sudo apt-get -y update 
sudo apt-get -y install curl 
curl -L https://github.com/kubernetes/kompose/releases/download/v1.26.1/kompose-linux-amd64 -o kompose
chmod +x kompose
sudo install ./kompose /usr/local/bin/kompose
sudo apt-get update

# Validate 
kompose version 
<<COMMENT
# Output  
1.18.0 (06a2e56)
COMMENT

```

## kubectl
Based on [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

```bash
# prerequisites
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl

# Download the Google Cloud public signing key
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

# Add the Kubernetes apt repository
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

# Update apt package index with the new repository and install kubectl
sudo apt-get update
sudo apt-get install -y kubectl
sudo apt-get update

# Validate 
kubectl version 

<< COMMENT 
# Output
Client Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.3", GitCommit:"816c97ab8cff8a1c72eccca1026f7820e93e0d25", GitTreeState:"clean", BuildDate:"2022-01-25T21:25:17Z", GoVersion:"go1.17.6", Compiler:"gc", Platform:"linux/amd64"}
COMMENT
```

## minikube 
Based on [minikube](https://minikube.sigs.k8s.io/docs/start/)
```commandline
sudo apt-get update
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# Validate 
minikube version
<< COMMENT 
# Output
minikube version: v1.25.1
commit: 3e64b11ed75e56e4898ea85f96b2e4af0301f43d
COMMENT
```

