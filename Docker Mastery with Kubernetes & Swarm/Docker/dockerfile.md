# Dockerfile
Dockerfile is the recipe for a docker container 

## Sample Dockerfile
```
# Specify the "core" of the system (ex. OS, language, database)
FROM ${distribution}:${release}  

#Set env variables 
ENV env_variable_name=${value}

# Expose ports 
EXPOSE ${PORRT_NUM}

# Directory to work within 
WORKDIR /the/workdir/path

# Copy from local into WORKDIR
COPY ${source} ${local}
# Run shell processes
RUN ${process}

# Script to run once Docker is up 
CMD ${process}
```

## Building Image 
* When rebuilding, only changed values rebuilt 
* Components that change more freqently should be lower on the Dockerfile 
