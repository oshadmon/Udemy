# Build Process
The following shows how to build a docker image that'll work on multiple platforms 
https://collabnix.com/building-arm-based-docker-images-on-docker-desktop-made-possible-using-buildx/

## Prerequisite
Install docker using hub rather than docker.io process. [Link](install_docker.sh) 

## Steps
0. Validate `docker buildx` is installed 
```
docker buildx --help
```

1. Create a builder instance 
```
docker buildx create --name ${BUILD_NAME} --platforms linux/amd64,linux/arm64,linux/arm/v7 

# Example
docker buildx create --name testbuild --platforms linux/amd64,linux/arm64,linux/arm/v7 
```

2. To validate 
```
docker buildx ls 
```

3. Switch to builder
```
docker buildx use ${BUILD_NAME} 

# Example 
docker buildx use testbuild 
```

4. To validate builder is running properlly 
```
docker buildx inspect --bootstrap
```

5. Build Repo
```
docker buildx build . --platform linux/amd64,linux/arm64,linux/arm/v7 -t ${BUILD_NAME} 
```

