# Deploy Kubernetes 

## Convert docker-compose
1. Create [docker-compose.yml](../Docker/docker-compose.yml) for containers to convert into Kubernetes

2. Using kompose convert docker-compose.yml into separate yaml files.
```shell
kompose convert

# Output 
INFO Kubernetes file "database-deployment.yaml" created 
INFO Kubernetes file "pgdata-persistentvolumeclaim.yaml" created
```

## Updating Kubernetes files 
1. Convert Database Username & Password with base64
```shell
echo -n '${VARIABLE_TO_CONVERT}' | base64
```

2. Create `secret.yaml` file - based on [docker-compose.yml](../Docker/docker-compose.yml) 
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: postgres-secret
data:
  PSQL_USERNAME: YW55bG9n # ori 
  PSQL_PASSWORD: ZGVtbw== # demo 
```

3. Update database-deployment.yaml to use secret.yaml for username and password (update env params)
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    kompose.cmd: kompose convert
    kompose.version: 1.18.0 (06a2e56)
  creationTimestamp: null
  labels:
    io.kompose.service: database
  name: database
spec:
  replicas: 1
  strategy:
    type: Recreate
  template:
    metadata:
      creationTimestamp: null
      labels:
        io.kompose.service: database
    spec:
      containers:
      - env:
        - name: POSTGRES_PASSWORD
          valueForm:
            name: postgres-secret
            value: PSQL_PASSWORD
        - name: POSTGRES_USER
          valueForm:
            name: postgres-secret
            value: PSQL_USERNAME
        image: postgres:14.0-alpine
        name: postgres
        resources: {}
        stdin: true
        tty: true
        volumeMounts:
        - mountPath: /var/lib/postgresql/data
          name: pgdata
      restartPolicy: Always
      volumes:
      - name: pgdata
        persistentVolumeClaim:
          claimName: pgdata
status: {}
```



