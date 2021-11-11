# EdgeX Setup process 
0. create a directory from which work is done unless stated otherwise 
```
mkdir edgex 
cd edgex 
``` 
 
1. Download EdgeX 
```
# download 
wget https://raw.githubusercontent.com/jonas-werner/EdgeX_Tutorial/master/docker-compose_files/docker-compose_step1.yml

# create a copy which will be using 
cp docker-compose_step1.yml docker-compose.yml 
```

2. Start EdgeX as is 
```
# download packages 
docker-compose pull 

# start EdgeX 
docker-compose up -d 
```
3. Validate EdgeX is running 
   * Validate services through browser - http://127.0.0.1:8500/ui/dc1/services 
   * Validate services through docker 
```
# check docker images 
docker image ls 

# check docker containers 
docker ps -a 
```
   * valide services through cURL 
```
# list of connected device(s) 
curl http://localhost:48082/api/v1/device
```

4. Stop EdgeX 
```
# Stop & remove docker containers
docker-compose down

# Remove volumes whe removing containers 
docker-compose down -v 
```
