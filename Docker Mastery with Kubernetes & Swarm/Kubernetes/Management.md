# Managements
## Basic Functions
* Generators (ex. _run_, _create_, _expose_) - processes with some kind of automation behind them. Samples Include:
  * services 
  * number of replicas (default 1)
  * ports to open
  
## Imperative vs Declarative 
* **Imperative**: Focus on _how_ a program operates (ie. order of steps) 
  * Way to learn any tool 
  * Know the current state (at each point)
  * not easy to automate
* **Declarative**: Focus on _what_ a program should do 
  * know the end result
  * don't know the current state 
  * same process can be used for multiple things 
  * Requires understanding YAML 
  * Easy to automate
* [GitOps](https://about.gitlab.com/topics/gitops/) - Utilziing _Git_ for DevOps
 
### Sample YAML
The following YAML was generated using: `kubectl create deployment sample --image nginx --dry-run -o yaml`
* `--dry-run`: pretend to run 
* `-o yaml`: show generated YAML file

```yaml
apiVersion: apps/v1 # API version  
kind: Deployment # kind of deployment
metadata: # resource information 
  creationTimestamp: null
  labels: # attributes related to deployment 
    app: sample
  name: sample
spec:  
  replicas: 1 # number of replicas
  selector: # match deployment to pod 
    matchLabels:
      app: sample
  strategy: {}
  template: # pod replicate information 
    metadata:
      creationTimestamp: null
      labels:
        app: sample
    spec: # "docker information" 
      containers: 
      - image: nginx
        name: nginx
        resources: {}
status: {}
```

## Management Approaches

| Imperative Commands | Imperative Objects | Declarative Objects | 
| ------------------- | ------------------- | ------------------- |
| Used for dev/learning/personal projects | utilizing files instead for production | only using YAML |
| | save changes | best for production | 
| hard to automate | hard to automate | easy to automate | 
| | | harder to understand and predict changes |
| Commands Used <br><ul><li>`run`<li>`expose`<li>`scale`<li>`edit`<li>`create`<li>`deploy`</li></ul> | Commands Used: <b><ul><li>`create -f`<li>`replace -f`<li>`delete -f`</li></ul>  |Commands Used: <b><ul><li>`apply -f`</li></ul>  |

* Do not mix the three approaches 
* Learn imperative CLI 
* Use YAML 


## Future of kubectl
* Starting at v1.12 `kubectl run` changes
* deprecate all generators but starting a pod - to make it as similar to _Docker_ as possible
* reason to use _Pod_ is when you want similar experience to `docker run` - one shot container 

