# Kubernetes CLI

## Deployment Process 
1. convert docker-compose for kubernetes 
```commandline
mkdir $HOME/kube
cd $HOME/kube 
kompose convert -f ~/docker-compose/docker-compose.yml
```

2. Start minikube & Evaluate
```commandline
minikube start --insecure-registry="${LOCAL_IP}:5000"
eval $(minikube docker-env)
# for RPI 
minikube -p minikube docker-env
```

3. Start processes
```commandline
kubectl apply -f $HOME/kube
```

4Stop processes 
```commandline
kubectl delete -f $HOME/kube
```

## Updating Service Information
In cases where standard `port-forwarding` doesn't work (ex. Minikube), user should attempt to change `ClusterIP` to `NodePort` in the 
1. Edit the service you're unable to `port-forward` against
```commandline
kubectl edit service ${SERVICE_NAME}
```

2. replace `ClusterIP` to `NodePort`(using sed)
```yaml
# before 
  ports:
  - name: "13480"
    nodePort: 31266
    port: 13480
    protocol: TCP
    targetPort: 13480
  - name: "13481"
    nodePort: 31956
    port: 13481
    protocol: TCP
    targetPort: 13481
  - name: "13482"
    nodePort: 31296
    port: 13482
    protocol: TCP
    targetPort: 13482
  selector:
    io.kompose.service: ${SERVICE_NAME}
  sessionAffinity: None
  type: ClusterIP

# after 
  ports:
  - name: "13480"
    nodePort: 31266
    port: 13480
    protocol: TCP
    targetPort: 13480
  - name: "13481"
    nodePort: 31956
    port: 13481
    protocol: TCP
    targetPort: 13481
  - name: "13482"
    nodePort: 31296
    port: 13482
    protocol: TCP
    targetPort: 13482
  selector:
    io.kompose.service: ${SERVICE_NAME}
  sessionAffinity: None
  type: NodePort
```

3. Generate `IP:PORT` to execute against
```commandline
minikube service --url ${SERVICE_NAME}
```
## Commands
* list (running) pods
```commandline
kubectl get pods
```
* Access shell node correlated to a specific pod 
```commandline
kubectl exec -it ${POD_NAME} -- /bin/bash
```
* Get services 
```commandline
kubectl get services 
```
* Describe pod
```commandline
kubectl describe pod ${POD_NAME}
```
* Pod logs - shows program output 
```commandline
kubectl logs ${POD_NAME} 
```
* Port-Forwarding 
```commandline
kubectl port-forward --address=${LOCAL_IP} service/${SERVICE_NAME} LOCAL_PORT:REMOTE_PORT
```