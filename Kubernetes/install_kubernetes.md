# Installation

## Kubernetes 
Based on [LinOxide](https://linoxide.com/install-kubernetes-on-ubuntu/)
0. Configure Hostname & disable host memory
```bash
sudo hostnamectl set-hostname
sudo swapoff -a
```
1. Add repository (as root user)
```bash 
sudo -i
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
```
2. Add repo & update
```bash 
echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" >> ~/kubernetes.list
sudo mv ~/kubernetes.list /etc/apt/sources.list.d
sudo apt update
```

3. Install Kubernetes
* Install
```bash
sudo apt install -y kubeadm=1.18.0-00 --allow-unauthenticated
curl -s https://packages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages | grep Version | awk '{print $2}'
```
* Validate 
```shell
kubectl version 
# Client Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.3", GitCommit:"816c97ab8cff8a1c72eccca1026f7820e93e0d25", GitTreeState:"clean", BuildDate:"2022-01-25T21:25:17Z", GoVersion:"go1.17.6", Compiler:"gc", Platform:"linux/amd64"}
# The connection to the server localhost:8080 was refused - did you specify the right host or port?
```


## Kompose
Kompose is a tool to migrate from [docker-compose](https://docs.docker.com/compose/) to Kubernetes  

1. Download Kompose from GitHub
```bash
curl -L https://github.com/kubernetes/kompose/releases/download/v1.18.0/kompose-linux-amd64 -o kompose
```

2. Convert _kompose_ to executable & move to `PATH`
```commandline
chmod +x kompose
sudo mv ./kompose /usr/local/bin/kompose

# validate version
kompose version
```
