# Docker Buildx 

The following provides an example of how to use docker buildx 

## Steps
0. [Install docker](docker_install.sh) 

1. Build buildx
```
docker build --platform=local -o . git://github.com/docker/buildx
mkdir -p ~/.docker/cli-plugins
mv buildx ~/.docker/cli-plugins/docker-buildx
```

2. Register ARM device 
```
docker run --rm --privileged docker/binfmt:820fdd95a9972a5308930a2bdfb8573dd4447ad3 

```

3. Validate docker buildx is avilable 
```
docker buildx ls
NAME/NODE DRIVER/ENDPOINT             STATUS  PLATFORMS
default   docker                              
  default default                     running linux/amd64, linux/386, linux/arm64, linux/ppc64le, linux/s390x, linux/arm/v7, linux/arm/v6
```

4. Create buildx env 
```
docker buildx create --name ${ENV_NAME}
```

5. Enable new env 
```
docker buildx use ${ENV_NAME}
docker buildx inspect --bootstrap

# validate (example) 
user@vm1:~$ docker buildx ls 
NAME/NODE DRIVER/ENDPOINT             STATUS  PLATFORMS
test *    docker-container                    
  test0   unix:///var/run/docker.sock running linux/amd64, linux/386, linux/arm64, linux/ppc64le, linux/s390x, linux/arm/v7, linux/arm/v6
default   docker                              
  default default                     running linux/amd64, linux/386, linux/arm64, linux/ppc64le, linux/s390x, linux/arm/v7, linux/arm/v6
```

6. Build Image
```
docker buildx build . --platform linux/amd64,linux/386,linux/arm64,linux/ppc64le,linux/s390x,linux/arm/v7,linux/arm/v6  -t ${IMAGE_NAME}:${IMAGE_TAG} --push
```

## Links 
* https://medium.com/@artur.klauser/building-multi-architecture-docker-images-with-buildx-27d80f7e2408
* https://collabnix.com/building-arm-based-docker-images-on-docker-desktop-made-possible-using-buildx/ 
* https://www.docker.com/blog/getting-started-with-docker-for-arm-on-linux/ 

