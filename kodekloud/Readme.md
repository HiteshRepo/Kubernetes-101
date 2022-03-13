## Core Concepts Revision Tasks

### Pods

1. Get number of pods running in default namespace: `kubectl get pods`
2. Create pod running nginx container: `refer kodekloud/pods/pod-nginx.yaml`
3. Get number of pods running in default namespace now (some more pods were spun up by kodekloud): `kubectl get pods`
4. Check image used to create the new pods: `kubectl describe pod <pod-name>`
5. Check node on which the new pods are running: `kubectl get pods -o wide` - check NODE column
6. Number of containers part of new pod 'webapp' (a new pod was spun up by kodekloud): `kubectl get pods` - check READY column (containers in ready state/total containers in the pod)
7. Image names of containers in pod webapp: `kubectl describe pod webapp` - check under containers[i]/Image
8. Check state of container 'agentx' in pod webapp: `kubectl describe pod webapp` - check under containers/agentx/State
9. Reason for container 'agentx' in pod webapp being in error state: `kubectl describe pod webapp` - check under events section
10. Indication of READY column: `kubectl get pods` - it denotes containers in ready state/total containers in the pod
11. Delete webapp pod: `kubectl delete pod webapp`
12. Create pod running redis container: `refer kodekloud/pods/pod-redis.yaml`

### Replicaset
1. Get number of pods in default namespace: `kubectl get pods`
2. Get number of replica set in default namespace: `kubectl get replicaset`
3. Desired number of pods in a replica set: `kubectl get replicaset` - check under DESIRED column
4. Check image used to create pods under a replica set: `kubectl describe replicaset <replica-set-name>` - PodTemplate/Containers/<conatiner-name>/Image
5. Number of pods in READY state in a replica set: `kubectl get replicaset` - check under READY column
6. Check the reason why some pods under a replica set is not ready/running: `kubectl get pods` - check pods under the replica set - describe the pod that is not ready/running and check the events section
7. Delete a pod: `kubectl delete pod <pod-name>`
8. Reason to why number of pods remain same if deleted: Replica set ensures the desired number of pods are always running/created
9. Create a replica set: `kubectl create -f <replica set file name>`
10. Reason for error `'Unable recognize "<file anme>": no matches for kind "ReplicaSet" in version "v1"'`: ReplicaSet are present under apiVersion 'apps/v1'
11. Reason for error `The ReplicaSet "replicaset-2" is invalid: spec.template.metadata.labels: Invalid value: map[string]string{"tier":"nginx"}: 'selector' does not match template 'labels'`: pod/metadata/label should match with replicaset/spec/selector/matchLabels
12. Scale a replica set to more/less pods: `kubectl scale --replicas=6 replicaset <replica-set-name>`
13. Scale a replica set to more/less pods: `kubectl edit replicaset <replica-set-name>` -> vi editor -> change replicas count

### Deployment
1. Get number of pods in default namespace: `kubectl get pods`
2. Get number of replica set in default namespace: `kubectl get replicaset`
3. Get number of deployment in default namespace: `kubectl get deployment`
4. Check image used to create pods under a deployment: `kubectl describe deployment <deployment-name>` - PodTemplate/Containers/<conatiner-name>/Image
5. Check the reason why some pods under a deployment is not ready/running: `kubectl get pods` - check pods under the deployment - describe the pod that is not ready/running and check the events section

### Namespace
1. Get number of namespace: `kubectl get namespace`
2. Get number of pods in 'research' namespace: `kubectl get pods --namespace=research`
3. Create a pod in 'finance' namespace: `refer kodekloud/pods/pod-redis-ns.yaml`

### Imperative commands
1. Deploy a pod named nginx-pod using the nginx:alpine image: `kubectl run nginx --image=nginx:alpine  --dry-run=client -o yaml` -> copied the declarative YAML result and changed the metadata/name to nginx-pod
2. Deploy a redis pod using the redis:alpine image with the labels set to tier=db: `kubectl run redis --image=redis:alpine  --dry-run=client -o yaml`
3. Create a service redis-service to expose the redis application within the cluster on port 6379: Refer kodekloud/imperative/commands.yaml
4. Create a deployment named webapp using the image kodekloud/webapp-color with 3 replicas: Refer kodekloud/imperative/commands.yaml
5. Create a new pod called custom-nginx using the nginx image and expose it on container port 8080: Refer kodekloud/imperative/commands.yaml
6. Create a new namespace called dev-ns: `kubectl create namespace dev-ns`
7. Refer kodekloud/imperative/commands.yaml last deployment
8. Refer kodekloud/imperative/commands.yaml last
