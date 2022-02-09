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
* Do not mix the three approaches 
* Learn imperative CLI 
* Use YAML 

<table>
  <tr>
    <th></th>
    <th>Imperative Commands</th>
    <th>Imperative Objects</th>
    <th>Declarative Objects</th>
  </tr>
  <tr>
    <td></td>
    <td>Used for dev/learning/personal projects</td>
    <td>utilizing files instead for production</td>
    <td>only using YAML</td>
  </tr>
  <tr>
    <td></td>
    <td></td>
    <td>Save Changes</td>
    <td>Best for production</td>
  </tr>
  <tr>
    <td></td>
    <td>hard to automate</td>
    <td>hard to automate</td>
    <td>easy to automate</td>
  </tr>
  <tr>
    <td></td>
    <td></td>
    <td></td>
    <td>harder to understand and predict changes</td>
  </tr>
  <tr>
    <th>Utilized Commands</th>
    <td valign="top">
      <ul valign="top">
        <li>run</li>
        <li>expose</li>
        <li>scale</li>
        <li>edit</li>
        <li>create</li>
        <li>deploy</li>
      </ul>
    </td>
    <td valign="top">
      <ul valign="top">
        <li>create -f</li>
        <li>replace -f</li>
        <li>delete -f</li>
      </ul>
    </td>
    <td valign="top">
      <ul valign="top">
        <li>apply -f</li>
      </ul>
    </td>
  </tr>
</table>


## Future of kubectl
* Starting at v1.12 `kubectl run` changes
* deprecate all generators but starting a pod - to make it as similar to _Docker_ as possible
* reason to use _Pod_ is when you want similar experience to `docker run` - one shot container 

