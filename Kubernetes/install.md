# Install Kubernetes

The following directions are based-on [Computing for Geeks](https://computingforgeeks.com/deploy-kubernetes-cluster-on-ubuntu-with-kubeadm/#:~:text=Install%20Kubernetes%20Cluster%20on%20Ubuntu%2020.04%20%20,k8s-worker01.computingforgeeks.com%20%20%20Worker%20%20%20k8s-worker02.computingforgeeks.com%20)

## Requirements
Install either [docker](../Docker), [CRI-O](https://cri-o.io/) or [Containerd](https://containerd.io/)

## Install 
1. Upgrade Environment & Reset `systemctl`   
```bash 
sudo apt-get -y update
sudo apt-get -y upgrade
sudo systemctl reboot # <-- this step will restart the machine 
sudo apt-get -y update
```

2. Restart `systemctl` 
```bash
sudo systemctl reboot
```

3. Install curl & [apt-transport-https](https://manpages.ubuntu.com/manpages/bionic/man1/apt-transport-https.1.html)
 ```bash
sudo apt-get -y install curl apt-transport-https
```

4. Download & set deb package with Kubernetes
```bash 
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
```

5. Update env due to the addition to `deb` in previous step 
```bash
sudo apt-get -y update
```

6. Install kubelet, kubeadm and kubectl
```bash
sudo apt-get -y install kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

6b. Validate Install version
```bash
kubectl version --client && kubeadm version
```

## Disable (memory) SWAP 
1. Validate `/etc/fstab` exits
```bash
ls /etc/fstab
```

2. Update swap 
```bash 
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
```

3. Permanently stop SWAP 
```bash
sudo swapoff -a
```


 