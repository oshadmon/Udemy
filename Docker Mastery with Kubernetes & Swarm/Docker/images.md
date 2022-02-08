# Images 

## What is an Image
* App binaries and dependency; as weell as metadata about the the image eand how it's run
* Host provides the OS / kernel, thus docker is just starting a applications 
* Big images (like OS) can also be run as an applicaton lik instance
* Ultimatly an image is a union of OS and correlated programs that are required to run a specific script. These details are contained in [Dockerfile](dockerfile.md)
* An  advantage os using docker is if multiple images use the same set of packages, those packages are downloaded / installed only once. 

## Image Tags 
* Images don't have names, but rather a combination of repo and version (ie tags). 
* Can be created as either part of a build (most common) or seperate from a buld.
* When downloading the tag alllow users to select which version to download

## Pushing to Docker Hub
```
# Setting tag 
docker tag ${docker_image}:${version} ${user}/${docker_image}:${version}

# login to docker 
docker login

# Push
docker image push ${user}/${docker_image}:${version} 
```
