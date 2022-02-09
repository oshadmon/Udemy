# Networking
## Service Types
* **Service**: An endpoint that's consistent to other things can access it.  
  * doesn't automatically get DNS name 
  * A service is a stab;ele address for pod(s) - to connect to a pod we need a service
* **CoreDNS**: allows users to resolves services by name

### Network Types
* ClusterIP (default) - internal virtual IP, that's reachable within the cluster
* NodePort
  * designed for something outside the cluster to communicate with the cluster 
  * (high) port gets assigned to the service
* LoadBalancer - control endpoint(s) external to the cluster 
  * requires infra support for loadbalancer (mainly in the cloud) 
  * Creates **both** _NodePort_ and _ClusterIP_ services, but tells LoadBalancer to use _NodePort_
* ExternalName (least use) - when stuff in the cluster needs to talk to the outside 
  * Adds CNAME DNS record to CoreDNS
  * Control the DNS inside the kubernetes 
* Kubernetes Ingress - Used for HTTP


## ClusterIP Service
1. create deployment for HTTP webserver 
```bash
kubectl create deployment httpenv --image bretfisher/httpenv
```

2. replicate 
```bash
kubectl scale deploy/httpenv --replicas=5
```

3. Create service
```bash
kubectl expose deploy/httpenv --port 8888
```

4. Access a parallel service for testing
  * [bretfisher/netshoot](https://github.com/bretfisher/netshoot) - A _Doker_ & _Kubernetes_ network testing tool
```bash
kubectl run tmp-shell --rm -it --image bretfisher/netshoot -- bash 
``` 

5. Basic cURL
```bash
curl [containerName]:[PORT]]

# Example
bash-5.1# curl httpenv:8888 | jq 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   437  100   437    0     0   331k      0 --:--:-- --:--:-- --:--:--  426k
{
  "HOME": "/root",
  "HOSTNAME": "httpenv-844bffcb6-tgts8",
  "KUBERNETES_PORT": "tcp://10.96.0.1:443",
  "KUBERNETES_PORT_443_TCP": "tcp://10.96.0.1:443",
  "KUBERNETES_PORT_443_TCP_ADDR": "10.96.0.1",
  "KUBERNETES_PORT_443_TCP_PORT": "443",
  "KUBERNETES_PORT_443_TCP_PROTO": "tcp",
  "KUBERNETES_SERVICE_HOST": "10.96.0.1",
  "KUBERNETES_SERVICE_PORT": "443",
  "KUBERNETES_SERVICE_PORT_HTTPS": "443",
  "PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
}
```

## NodePorts & LoadBalancer
**NodePort**
* To start service of type _NodePort_
  * Default type: is _ClusterIP_
  * When doing `kube get service` the port on the right is based on a default range between 30000-32767
```bash
kubectl expose deploy/httpenv --port 8888 --name httpenv-np --type NodePort
```

**LoadBalancer** (not standard)
* If you run LoadBalancer, but DNE it'll switch to _NodePort_

## DNS
* DNS is provided by CoreDNS
  * DNS-Based service discovery
* Out of the box you can communicate between pods within the same cluster 
* to list namespaces: `kubectl get namespaces`
  * organizational parameters 

