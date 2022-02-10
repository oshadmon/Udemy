# YAML  
* Command used: `kubectl apply -f`
* create & update from single file  
  * create & update a full directory
  * Apply from file stored online: `kubectl apply -f https://bret.run/pod.yaml`
* In the background kubernetes converts to JSON
* Resources can be in one file, or sparate files
* Manifest: details in YAML file regarding a deployment
  * To use multiple Manifest in a file use `---` to separate between them
```YAML
# Based on a manifest 
apiVersion: 
kind: 
metadata:
spec:
```

## YAML Basics
* **apiVersion**: version based on _API Group_  
* **kind**: Type of manifest 
  * _Kind_ - what goes in the manifest 
  * _API GROUP_ - related to the API version
    * Some resources have multiple of APIs
* **metadata**: name of manifest 
* **spec**: action(s) for different different resources
```bash
# for a full list of kube kinds: 
kubectl api-resources

# for a list of apiVersions
kubectl api-versions
 
# services 
kubectl explain service --recursive

# deployment 
kubectl explain deploy --recursive
```

### sample YAML formatting 
* Overall YAML file (no details)
```bash
kubectl explain ${kind} --recursive
```

* Details regarding spec
```bash
kubectl explain ${kind}.spec

# specific variable within spec  
kubectl explain ${kind}.spec.type
```
* [API docs](https://kubernetes.io/docs/reference/) has the same information as the examples

## Spec in YAML  
* spec can be within spec 
```commandline
kubectl explain deployment.spec.template
KIND:     Deployment
VERSION:  apps/v1

RESOURCE: template <Object>

DESCRIPTION:
     Template describes the pods that will be created.

     PodTemplateSpec describes the data a pod should have when created from a
     template

FIELDS:
   metadata     <Object>
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

   spec <Object>
     Specification of the desired behavior of the pod. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
```

## Dry-run & Diffs 
* 