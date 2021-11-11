# Setup EdgeX from Compile 
## Prerequisites
1. Update & Upgrade 
```
sudo apt-get -y update ; sudo apt-get -y upgrade ; sudo apt-get -y update
```

2. Git
```
sudo apt-get -y install git 
```

3. Install [Golang](https://golang.org/) 
```
# Install 
sudo apt-get -y install golang-go

# Validate 
go version 
```

4. Install [ZeroMQ](https://zeromq.org/)
```
sudo apt-get -y install libzmq3-dev
```

5. Install [Redis](https://redis.io/) - [DigitalOcean directions](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-18-04#:~:text=%20How%20To%20Install%20and%20Secure%20Redis%20on,Redis%20is%20only%20accessible%20from%20localhost.%20More%20) 
```
# Install 
sudo apt-get -y install redis-server

# Update redis.cnf 
sed -i 's/supervised no/supervised systemd/g' /etc/redis/redis.conf

# Restart Service 
sudo service redis restart 

# Install CLI
sudo apt-get -y install radis-tools 

# Validate Radis -- should say if connected 
radis-cli 
```
		
## Download & Install 
1. Clone [EdgeX](https://github.com/edgexfoundry/edgex-go) 
```
cd $HOME 
git clone https://github.com/edgexfoundry/edgex-go
cd edgex-go 
```

2. Build EdgeX
``` 
make build 
```

3. Start EdgeX - output to file 
```
make run &> output.err & 
```

4. Validate EdgeX is running 
```
```
**Note**: Running directly from git doesn't contain any extra components the docker-compose file has. 
