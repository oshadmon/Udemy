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
```

3. Start processes
```commandline
kubectl apply -f $HOME/kube
```

4. Stop processes 
```commandline
kubectl delete -f $HOME/kube
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
kubectl port-forward --address=${LOCAL_IP} service/${SERVICE_NAME} ${PORT}:${PORT}
```