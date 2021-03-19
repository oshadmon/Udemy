# Prerequisites

1. Update & Upgrade 
```
sudo apt-get -y update ; sudo apt-get -y upgrade ; sudo apt-get -y update
```

2. Install [Golang](https://golang.org/) 
```
wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
sudo export PATH=$PATH:/usr/local/go/bin
```

3. validate version 
```
go version
```
