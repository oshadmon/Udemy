# Persistent Data 
* Containers are meant to be temporaryy and unchanged, as such data should **not** be part of the container 
* Data should be stored in volumese as they outlive the container itself. (Ex. [MySQL Dockerfile](https://github.com/docker-library/mysql/blob/bc6e37a2bed792b1c4fc6ab1ec3ce316e6a5f061/8.0/Dockerfile))
* Alternativly users can bi-directionally mount a directory to store data locally. This option skips UFS process. 
* When using a mount, changes that occur on one side (the host) will also appear on the docker. 

```
# Start a container with volume 
docker container run -d --name mysql -v my-db:/var/lib/mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=True mysql 

# Start a container with mount 
docker container run -v ${host_path}:${container_path}
```