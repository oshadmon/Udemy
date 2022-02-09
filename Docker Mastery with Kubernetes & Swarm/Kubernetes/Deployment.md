# Deploy Kubernetes Pod
## Commandline
* Check if `kubectl` is running: `kubectl version`
* Deploy via CLI: 
  * kubectl v1.18 or newer: `kubectl create deployment my-nginx --image nginx:1.16.0-alpine`
  * kubectl v1.17 or older: `kubectl run my-nginx --image nginx:1.16.0-alpine`
* List pods: `kubectl get pods`
* List all `kubectl` objects: `kubectl get all` 
* Remove resources: `kuebctl delete deployment my-nginx`

## Scaling ReplicaSets
### Commands 
1. Start a deployment of `apache` pod
```bash
kubectl create deployment my-appache --image httpd
```

2. To replicate
```bash
kubectl scale deploy/my-apache --replicas 2

# alternative cmd 
kubectl scale deployment my-apache --replicas 2
```

3. Validate
```bash
kubectl get all 
NAME                            READY   STATUS    RESTARTS   AGE
pod/my-apache-94f68777f-4lq5n   1/1     Running   0          45s
pod/my-apache-94f68777f-qn2kl   1/1     Running   0          2m35s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   3h50m

NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/my-apache   2/2     2            2           2m35s

NAME                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/my-apache-94f68777f   2         2         2       2m35s
```

### Deployment Plan
* Deployment update to 2 replicas
* ReplicaSet Controller sets pod count to 2 
* Control Plan assigns node to pod 
* Kubelet sees pod is needed & starts a container

## Inspecting Process(s)
* logs - ongoing information on a process  
```bash
# deployment
kubectl logs deploy/my-apache

# pod
kubectl logs pod/my-apache-94f68777f-qn2kl
```

* describe - summary of a process 
```bash
# deployment
kubectl describe deploy/my-apache

# pod
kubectl describe pod/my-apache-94f68777f-qn2kl
```

## Notes
1. User should change deployment(s)  rather than pods
2. In replication when removing a pod, the deployment process automtically recreates said pod