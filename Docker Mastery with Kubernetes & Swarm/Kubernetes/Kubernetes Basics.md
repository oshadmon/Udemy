# Kubernetes Basics

## Definitions:
* **Kubernetes** (K8s): The whole orchestration
* **kubectl**: CLI to configure K8s
* **Node**: single server in the cluster
* **kubelet**: agent running on node
* **Control Plane**: Set of containers that manage the cluster (aka _master_)
* **Pod**: one or more container(s) running together on a node
    * the most basic unit of deployment - containers are within pods
* **Controller**: Creating / Updating pods and other objects, such as:
    * deployment
    * ReplicaSet
    * StatefulSet
    * DaemonSet
    * Job & CronJob
    * etc.
* **Service**: network endpoint to connect to a pod
* **Namespace**: Filtered group of objects in a cluster

## Install
* series of containers, CLI and configurations
* Platforms to execute kubernetes
    * [play-withk8s](play-withk8s.com)
    * [kadakoda](https://www.katacoda.com/) - has an array _sandboxes_ to tru different tools
* [Install Docs](../../Kubernetes/installation.md)

### Kubernetes Dashboard
* [Kubernetes Dashboard](https://github.com/kubernetes/dashboard)
* Not deployed by default
* Companies were hacked using the dashboard (due to no authentication)


## Kubectl commands
* `kubectl run` - pod creation
* `kubectl create` - create resource via _CLI_ or _YAML_
* `kubectl apply` - deploy / update via _YAML_


