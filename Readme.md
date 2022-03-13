## Exercise - 1
1. Install kubectl - brew install kubectl 
2. Install minikube - brew install minikube
3. minikube status
4. minikube start --memory=16000m --cpus=4 --driver=docker
5. minikube status
    * minikube status
    * minikube
    * type: Control Plane
    * host: Running
    * kubelet: Running
    * apiserver: Running
    * kubeconfig: Configured
6. kubectl config get-contexts
CURRENT   NAME       CLUSTER    AUTHINFO   NAMESPACE
*minikube   minikube   minikube   default
7. kubectl get nodes -o wide
NAME       STATUS   ROLES                  AGE     VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
minikube   Ready    control-plane,master   7m32s   v1.22.3   192.168.49.2   <none>        Ubuntu 20.04.2 LTS   5.10.47-linuxkit   docker://20.10.8
8. cat ~/.kube/config
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Users/hitesh.pattanayak/.minikube/ca.crt
    extensions:
    - extension:
        last-update: Sat, 11 Dec 2021 10:19:55 IST
        provider: minikube.sigs.k8s.io
        version: v1.24.0
      name: cluster_info
    server: https://127.0.0.1:65441
  name: minikube
contexts:
- context:
    cluster: minikube
    extensions:
    - extension:
        last-update: Sat, 11 Dec 2021 10:19:55 IST
        provider: minikube.sigs.k8s.io
        version: v1.24.0
      name: context_info
    namespace: default
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: /Users/hitesh.pattanayak/.minikube/profiles/minikube/client.crt
    client-key: /Users/hitesh.pattanayak/.minikube/profiles/minikube/client.key
```
9. kubectl config view
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Users/hitesh.pattanayak/.minikube/ca.crt
    extensions:
    - extension:
        last-update: Sat, 11 Dec 2021 10:19:55 IST
        provider: minikube.sigs.k8s.io
        version: v1.24.0
      name: cluster_info
    server: https://127.0.0.1:65441
  name: minikube
contexts:
- context:
    cluster: minikube
    extensions:
    - extension:
        last-update: Sat, 11 Dec 2021 10:19:55 IST
        provider: minikube.sigs.k8s.io
        version: v1.24.0
      name: context_info
    namespace: default
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: /Users/hitesh.pattanayak/.minikube/profiles/minikube/client.crt
    client-key: /Users/hitesh.pattanayak/.minikube/profiles/minikube/client.key
```
10. minikube ip
11. minikube ssh
12. eval $(minikube docker-env)
13. docker images - gets images within minikube context

## Excercise - 2

1. kubectl get nodes -o wide
NAME       STATUS   ROLES                  AGE     VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
minikube   Ready    control-plane,master   3h50m   v1.22.3   192.168.49.2   <none>        Ubuntu 20.04.2 LTS   5.10.47-linuxkit   docker://20.10.8

2. kubectl get pods
No resources found in default namespace.

3. kubectl apply -f Ex-2-Pods/metadata-pod.yaml -> pod/metadata created
4. kubectl describe pod metadata:
    Events:
      Type    Reason     Age   From               Message
      ----    ------     ----  ----               -------
      Normal  Scheduled  33s   default-scheduler  Successfully assigned default/metadata to minikube
      Normal  Pulling    31s   kubelet            Pulling image "sunitparekh/metadata:v1.0"
      Normal  Pulled     13s   kubelet            Successfully pulled image "sunitparekh/metadata:v1.0" in 17.638965766s
      Normal  Created    13s   kubelet            Created container metdata
      Normal  Started    13s   kubelet            Started container metdata
5. kubectl get pods -o wide                                                                                                                                                                     
    NAME       READY   STATUS    RESTARTS   AGE     IP           NODE       NOMINATED NODE   READINESS GATES
    metadata   1/1     Running   0          4m40s   172.17.0.3   minikube   <none>           <none>
6. minikube ssh
7. curl 172.17.0.3:8080/actuator/info -> {"app":{"name":"Metadata Service","description":"Metadata service also known as config service. It hold the metadata/config required across different services.","version":"1.0.0"}}
8. kubectl apply -f https://k8s.io/examples/admin/dns/dnsutils.yaml -> pod/dnsutils created
9. kubectl get pods -o wide                                                                                                                                                                     
    NAME       READY   STATUS    RESTARTS   AGE     IP           NODE       NOMINATED NODE   READINESS GATES
    dnsutils   1/1     Running   0          32s     172.17.0.4   minikube   <none>           <none>
    metadata   1/1     Running   0          9m21s   172.17.0.3   minikube   <none>           <none>
10. kubectl exec -it dnsutils -- /bin/sh
11. apt-get install wget or apt install curl
12. curl 172.17.0.3:8080/actuator/info
    1. {"app":{"name":"Metadata Service","description":"Metadata service also known as config service. It hold the metadata/config required across different services.","version":"1.0.0"}}
13. kubectl port-forward pod/metadata 9090:8080 -> PTR - local-port:container-port
14. Hit localhost:9090/actuator/info in browser
15. kubectl apply -f Ex-2-Pods/metadata-pod-with-env.yaml -> pod/metadata-env created
16. kubectl get pods -o wide                                                                                                                                                                     
    NAME           READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
    dnsutils       1/1     Running   0          21m   172.17.0.4   minikube   <none>           <none>
    metadata       1/1     Running   0          30m   172.17.0.3   minikube   <none>           <none>
    metadata-env   1/1     Running   0          25s   172.17.0.5   minikube   <none>           <none>
17. minikube ssh
18. curl 172.17.0.5:8080/actuator/info -> {"app":{"name":"Metadata Service","description":"Metadata service also known as config service. It hold the metadata/config required across different 
19. kubectl apply -f Ex-2-Pods/metadata-pod-with-limits.yaml -> pod/metadata-limits created
20. kubectl get pods -o wide                                                                                                                                                                     
    NAME              READY   STATUS    RESTARTS   AGE     IP           NODE       NOMINATED NODE   READINESS GATES
    dnsutils          1/1     Running   0          26m     172.17.0.4   minikube   <none>           <none>
    metadata          1/1     Running   0          35m     172.17.0.3   minikube   <none>           <none>
    metadata-env      1/1     Running   0          5m27s   172.17.0.5   minikube   <none>           <none>
    metadata-limits   1/1     Running   0          25s     172.17.0.6   minikube   <none>           <none>
21. kubectl describe metadata-limits
    ```
    metdata:
        Container ID:   docker://fef13048a8e2aefae575227a6bf9a84b2361fac82fa036bd6b58c93219c6ad74
        Image:          sunitparekh/metadata:v1.0
        Image ID:       docker-pullable://sunitparekh/metadata@sha256:88c8333bf43df2c9ca3ac47e7d69dd3d19095858b69f328f2eb4f083cc717191
        Port:           8080/TCP
        Host Port:      0/TCP
        State:          Running
          Started:      Sat, 11 Dec 2021 14:52:47 +0530
        Ready:          True
        Restart Count:  0
        Limits:
          cpu:     500m
          memory:  512Mi
        Requests:
          cpu:     250m
          memory:  250Mi
        Environment:
          info.app.version:  2.0.0
        Mounts:
          /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-6qc79 (ro)
    ```
22. kubectl describe metadata-env - since limits are not specified, this will consume upto node’s available resources on load increase
```
    metdata:
        Container ID:   docker://9c8f33f423611fad4bb3d35e73273b134583367313cc5e976d3ed5ae424ebe34
        Image:          sunitparekh/metadata:v1.0
        Image ID:       docker-pullable://sunitparekh/metadata@sha256:88c8333bf43df2c9ca3ac47e7d69dd3d19095858b69f328f2eb4f083cc717191
        Port:           8080/TCP
        Host Port:      0/TCP
        State:          Running
            Started:      Sat, 11 Dec 2021 14:47:45 +0530
        Ready:          True
        Restart Count:  0
        Environment:
            info.app.version:  2.0.0
        Mounts:
            /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-mhdcl (ro)
```

## Exercise - 3

1. kubectl apply -f Ex-3-ReplicaSets/metadata-rs.yaml -> replicaset.apps/metadata-rs created
2. kubectl get rs                                                                                                                                                                               
    NAME          DESIRED   CURRENT   READY   AGE
    metadata-rs   2         2         0       19s
3. Kubectl get pods -o wide
    NAME                READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
    metadata-rs-9c56h   1/1     Running   0          59s   172.17.0.4   minikube   <none>           <none>
    metadata-rs-vgs69   1/1     Running   0          59s   172.17.0.3   minikube   <none>           <none>
4. kubectl port-forward pod/metadata-rs-9c56h 9090:8080 - hit localhost:9090/actuator/health - {"status":"UP"}
5. kubectl port-forward pod/metadata-rs-vgs69 9091:8080 - hit localhost:9091/actuator/health - {"status":"UP"}
6. kubectl delete rs/metadata-rs - replicaset.apps "metadata-rs" deleted - deletes above 2 pods too
7. Changing readinessProbe’s initialDelaySeconds to 2 from 20 and periodSeconds to 1 from 5
8. Changing livenessProbe’s initialDelaySeconds to 2 from 20 and periodSeconds to 1 from 5
9. kubectl apply -f Ex-3-ReplicaSets/metadata-rs.yaml -> replicaset.apps/metadata-rs created
10. kubectl get rs                                                                                                                                                                               
    NAME          DESIRED   CURRENT   READY   AGE
    metadata-rs   2         2         0       19s
11. Kubectl get pods -o wide
    NAME                READY   STATUS           RESTARTS                AGE   IP             NODE       NOMINATED NODE   READINESS GATES
    metadata-rs-clxfh   0/1     CrashLoopBackOff    6 (73s ago)          5m1s   172.17.0.4    minikube   <none>           <none>
    metadata-rs-hrqnb   0/1     CrashLoopBackOff    6 (73s ago)          5m1s   172.17.0.3    minikube   <none>           <none>
12. kubectl get events --sort-by='.metadata.creationTimestamp' --field-selector involvedObject.kind=Pod -w
    LAST SEEN   TYPE      REASON          OBJECT                  MESSAGE
        115s        Warning   Unhealthy       pod/metadata-rs-clxfh   Liveness probe failed: Get "http://172.17.0.3:8080/actuator/info": dial tcp 172.17.0.3:8080: connect: connection refused
        115s        Warning   Unhealthy       pod/metadata-rs-hrqnb   Liveness probe failed: Get "http://172.17.0.4:8080/actuator/info": dial tcp 172.17.0.4:8080: connect: connection refused
        115s        Warning   Unhealthy       pod/metadata-rs-hrqnb   Readiness probe failed: Get "http://172.17.0.4:8080/actuator/health": dial tcp 172.17.0.4:8080: connect: connection refused
        115s        Warning   Unhealthy       pod/metadata-rs-clxfh   Readiness probe failed: Get "http://172.17.0.3:8080/actuator/health": dial tcp 172.17.0.3:8080: connect: connection refused
        115s        Normal    Killing         pod/metadata-rs-hrqnb   Container metdata failed liveness probe, will be restarted
        115s        Normal    Killing         pod/metadata-rs-clxfh   Container metdata failed liveness probe, will be restarted
13. kubectl delete rs/metadata-rs - replicaset.apps "metadata-rs" deleted - deletes above 2 pods too

## Excercise 4 & 5

1. kubectl apply -f Ex-4-Services/metadata-svc.yaml -> service/metadata-svc created
2. kubectl apply -f Ex-4-Services/metadata-rs.yaml -> replicaset.apps/metadata-rs created
3. kubectl get pods -o wide                                     
    NAME                READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
    metadata-rs-8rzkf   1/1     Running   0          90s   172.17.0.4   minikube   <none>           <none>
    metadata-rs-mh7dx   1/1     Running   0          90s   172.17.0.3   minikube   <none>           <none>
4. kubectl get svc                                              
    NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    kubernetes     ClusterIP   10.96.0.1       <none>        443/TCP          12h
    metadata-svc   NodePort    10.110.70.144   <none>        8080:32323/TCP   70s
5. kubectl get endpoints                                        
    NAME           ENDPOINTS                         AGE
    kubernetes     192.168.49.2:8443                 12h
    metadata-svc   172.17.0.3:8080,172.17.0.4:8080   104s
6. kubectl apply -f https://k8s.io/examples/admin/dns/dnsutils.yaml -> created a pod to test svc dns within cluster
7. kubectl exec -it dnsutils — bin/sh
8. apt install curl
9. curl metadata-svc:8080/actuator/info
10. exit
11. minikube ip - 192.168.49.2
12. curl 192.168.49.2:32323/metadata -> [] as there is no entry there yet
13. curl --header “Content-Type: application/json” --request POST --data ‘{“group”: “voyager”, “name”: “city”, “value”:”Pune”}’ http://192.168.49.2:32323/metadata
14. curl 192.168.49.2:32323/metadata  -> gives different no. of entries every time because of in-memory datastore
15. step# 12 and 13 not working for me, so had to do port-forwarding of service:
    1. kubectl port-forward svc/metadata-svc 9090:8080
    2. curl 127.0.0.1:9090/metadata -> []
    3. curl --header “Content-Type: application/json” --request POST --data ‘{“group”: “voyager”, “name”: “city”, “value”:”Pune”}’ http://127.0.0.1:9090/metadata
    4. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value":"Bangalore"}' http://127.0.0.1:9090/metadata
    5. curl 127.0.0.1:9090/metadata -> [{"group":"voyager","lastUpdatedTs":"2021-12-11T17:39:32.527","name":"city","id":"61b4e2504cedfd0001613899"},{"group":"voyager","lastUpdatedTs":"2021-12-11T17:39:32.527","name":"city","id":"61b4e2494cedfd0001613894"}]
16. kubectl delete -f Ex-4-Services/metadata-rs.yaml -> replicaset.apps "metadata-rs" deleted
17. kubectl delete -f Ex-4-Services/metadata-svc.yaml -> service "metadata-svc" deleted
18. kubectl delete -f https://k8s.io/examples/admin/dns/dnsutils.yaml -> pod "dnsutils" deleted
19. kubectl apply -f Ex-4-Services/mongo-svc.yaml -> service/mongo created
20. kubectl get svc                                              
    NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)     AGE
    kubernetes   ClusterIP   10.96.0.1      <none>        443/TCP     12h
    mongo        ClusterIP   10.97.28.231   <none>        27017/TCP   59s
21. kubectl apply -f Ex-4-Services/mongo-pod.yaml -> pod/mongo-pod created
22. kubectl get pods -o wide                                     
    NAME        READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
    mongo-pod   1/1     Running   0          68s   172.17.0.3   minikube   <none>           <none>
23. kubectl get endpoints                                        
    NAME         ENDPOINTS           AGE
    kubernetes   192.168.49.2:8443   12h
    mongo        172.17.0.3:27017    2m21s
24. kubectl apply -f Ex-4-Services/metadata-svc.yaml -> service/metadata-svc created
25. kubectl apply -f Ex-4-Services/metadata-rs.yaml -> replicaset.apps/metadata-rs created
26. kubectl get svc                                              
    NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)     AGE
    kubernetes   ClusterIP   10.96.0.1      <none>        443/TCP     12h
    mongo        ClusterIP   10.97.28.231   <none>        27017/TCP   59s
    metadata-svc   NodePort    10.107.64.68   <none>        8080:32323/TCP   68s
27. kubectl apply -f Ex-4-Services/mongo-pod.yaml -> pod/mongo-pod created
28. kubectl get pods -o wide                                     
    NAME        READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
    mongo-pod   1/1     Running   0          68s   172.17.0.3   minikube   <none>           <none>
    metadata-rs-55cnp   1/1     Running   0          71s   172.17.0.4   minikube   <none>           <none>
    metadata-rs-tlqpt   1/1     Running   0          71s   172.17.0.5   minikube   <none>           <none>
29. kubectl get endpoints                                        
    NAME         ENDPOINTS           AGE
    kubernetes   192.168.49.2:8443   12h
    mongo        172.17.0.3:27017    2m21s
    metadata-svc   172.17.0.4:8080,172.17.0.5:8080   82s
30. kubectl port-forward svc/metadata-svc 9090:8080
31. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value":"Bangalore"}' http://127.0.0.1:9090/metadata
32. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value”:”Pune}’ http://127.0.0.1:9090/metadata
33. curl 127.0.0.1:9090/metadata -> [{"group":"voyager","lastUpdatedTs":"2021-12-11T17:39:32.527","name":"city","id":"61b4e2504cedfd0001613899"},{"group":"voyager","lastUpdatedTs":"2021-12-11T17:39:32.527","name":"city","id":"61b4e2494cedfd0001613894"}]
34. kubectl delete -f Ex-4-Services/metadata-rs.yaml -> replicaset.apps "metadata-rs" deleted
35. kubectl delete -f Ex-4-Services/mongo-pod.yaml -> pod "mongo-pod" deleted
36. kubectl delete -f Ex-4-Services/mongo-svc.yaml -> service "mongo" deleted
37. kubectl delete -f Ex-4-Services/metadata-svc.yaml -> service "metadata-svc" deleted

## Excercise 6

1. kubectl apply -f Ex-6-Deployment/mongo-svc.yaml -> service/mongo created
2. kubectl apply -f Ex-6-Deployment/mongo.yaml -> deployment.apps/mongo created
3. kubectl apply -f Ex-6-Deployment/metadata-svc.yaml -> service/metadata-svc created
4. kubectl apply -f Ex-6-Deployment/metadata.yaml -> deployment.apps/metadata created
5. kubectl get svc                                              
    NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    kubernetes     ClusterIP   10.96.0.1       <none>        443/TCP          13h
    metadata-svc   NodePort    10.106.94.123   <none>        8080:32323/TCP   66s
    mongo          ClusterIP   10.99.205.244   <none>        27017/TCP        6m29s
6. kubectl get pods -o wide                                     
    NAME                        READY   STATUS    RESTARTS   AGE     IP           NODE       NOMINATED NODE   READINESS GATES
    metadata-6fdd79b74f-44hj5   1/1     Running   0          53s     172.17.0.5   minikube   <none>           <none>
    metadata-6fdd79b74f-xlhds   1/1     Running   0          53s     172.17.0.4   minikube   <none>           <none>
    mongo-67fb584bd8-j4jwf      1/1     Running   0          3m33s   172.17.0.3   minikube   <none>           <none>
7. kubectl get rs                                               
    NAME                  DESIRED   CURRENT   READY   AGE
    metadata-6fdd79b74f   2         2         2       93s
    mongo-67fb584bd8      1         1         1       4m13s
8. kubectl port-forward svc/metadata-svc 9090:8080
9. curl localhost:9090/actuator/info -> {“app":{"version":"10.0.0","name":"Metadata Service","description":"Metadata service also known as config service. It hold the metadata/config required across different services."}}
10. curl localhost:9090/actuator/health -> {“status":"UP"}
11. curl localhost:9090/metadata -> []
12. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value":"Bangalore"}' http://127.0.0.1:9090/metadata -> {“id":"61b4ef37cff47e0001fcc9ba","message":"Successfully saved metadata."}
13. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value”:”Pune”}’ http://127.0.0.1:9090/metadata -> {“id":"61b4ef8bcff47e0001fcc9bf","message":"Successfully saved metadata."}
14. curl localhost:9090/metadata -> [{“group":"voyager","lastUpdatedTs":"2021-12-11T18:36:19.396","name":"city","id":"61b4ef8bcff47e0001fcc9bf"},{"group":"voyager","lastUpdatedTs":"2021-12-11T18:36:19.399","name":"city","id":"61b4ef37cff47e0001fcc9ba"}]
15. Changed sunitparekh/metadata:v1.0 -> sunitparekh/metadata:v3.0 in metadata.yaml
16. kubectl apply -f Ex-6-Deployment/metadata.yaml -> deployment.apps/metadata configured
17. Change mongo-db version
18. kubectl apply -f Ex-6-Deployment/mongo.yaml -> deployment.apps/metadata configured
19. kubectl delete -f Ex-6-Deployment/metadata.yaml -> deployment.apps "metadata" deleted
20. k8-rotc git:(master) kubectl delete -f Ex-6-Deployment/metadata-svc.yaml -> service "metadata-svc" deleted
21. k8-rotc git:(master) kubectl delete -f Ex-6-Deployment/mongo.yaml -> deployment.apps "mongo" deleted
22. k8-rotc git:(master) kubectl delete -f Ex-6-Deployment/mongo-svc.yaml -> service "mongo" deleted


## Exercise 7
1. minikube ssh
2. mkdir mongodb
3. cd mongoldb
4. mkdir data
5. kubectl apply -f Ex-7-Volumes/mongo-svc.yaml -> service/mongo created
6. kubectl apply -f Ex-7-Volumes/mongo.yaml -> deployment.apps/mongo created
7. minikube ssh
8. ls -la mongodb/data/ -> files created by mongo db app
```
    drwxr-xr-x 4    999 docker  4096 Dec 12 04:01 .
    drwxr-xr-x 3 docker docker  4096 Dec 12 03:57 ..
    -rw------- 1    999 docker    46 Dec 12 03:59 WiredTiger
    -rw------- 1    999 docker    21 Dec 12 03:59 WiredTiger.lock
    -rw------- 1    999 docker  1246 Dec 12 04:01 WiredTiger.turtle
    -rw------- 1    999 docker 61440 Dec 12 04:01 WiredTiger.wt
    -rw------- 1    999 docker  4096 Dec 12 03:59 WiredTigerLAS.wt
    -rw------- 1    999 docker 20480 Dec 12 04:00 _mdb_catalog.wt
    -rw------- 1    999 docker 20480 Dec 12 04:00 collection-0-3240096960546485121.wt
    -rw------- 1    999 docker 20480 Dec 12 04:00 collection-2-3240096960546485121.wt
    -rw------- 1    999 docker  4096 Dec 12 03:59 collection-4-3240096960546485121.wt
    drwx------ 2    999 docker  4096 Dec 12 04:02 diagnostic.data
    -rw------- 1    999 docker 20480 Dec 12 04:00 index-1-3240096960546485121.wt
    -rw------- 1    999 docker 20480 Dec 12 04:00 index-3-3240096960546485121.wt
    -rw------- 1    999 docker  4096 Dec 12 03:59 index-5-3240096960546485121.wt
    -rw------- 1    999 docker 12288 Dec 12 04:01 index-6-3240096960546485121.wt
    drwx------ 2    999 docker  4096 Dec 12 03:59 journal
    -rw------- 1    999 docker     2 Dec 12 03:59 mongod.lock
    -rw------- 1    999 docker 20480 Dec 12 04:01 sizeStorer.wt
    -rw------- 1    999 docker   114 Dec 12 03:59 storage.bson
```
9. kubectl apply -f Ex-7-Volumes/metadata-svc.yaml -> service/metadata-svc created
10. kubectl apply -f Ex-7-Volumes/metadata.yaml -> deployment.apps/metadata created
11. kubectl get svc                                              
    NAME           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
    kubernetes     ClusterIP   10.96.0.1        <none>        443/TCP          23h
    metadata-svc   NodePort    10.100.192.181   <none>        8080:32323/TCP   34s
    mongo          ClusterIP   10.106.22.162    <none>        27017/TCP        4m29s
12. kubectl get pods -o wide                                     
    NAME                        READY   STATUS    RESTARTS   AGE     IP           NODE       NOMINATED NODE   READINESS GATES
    metadata-69b5c966c5-98wd6   1/1     Running   0          38s     172.17.0.4   minikube   <none>           <none>
    metadata-69b5c966c5-hhtbf   1/1     Running   0          38s     172.17.0.5   minikube   <none>           <none>
    mongo-fb89c49b7-qfg7f       1/1     Running   0          4m30s   172.17.0.3   minikube   <none>           <none>
13. kubectl get endpoints                                        
    NAME           ENDPOINTS                         AGE
    kubernetes     192.168.49.2:8443                 23h
    metadata-svc   172.17.0.4:8080,172.17.0.5:8080   88s
    mongo          172.17.0.3:27017                  5m23s
8. kubectl port-forward svc/metadata-svc 9090:8080
9. curl localhost:9090/actuator/info -> {"app":{"version":"10.0.0","name":"Metadata Service","description":"Metadata service also known as config service. It hold the metadata/config required across different services."}}
10. curl localhost:9090/actuator/health -> {"status":"UP","details":{"diskSpace":{"status":"UP","details":{"total":62725623808,"free":52071821312,"threshold":10485760}},"mongo":{"status":"UP","details":{"version":"4.2.6"}}}}
11. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value":"Bangalore"}' http://127.0.0.1:9090/metadata -> {“id":"61b5761c4cedfd0001d83975","message":"Successfully saved metadata."}
12. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value”:”Pune”}’ http://127.0.0.1:9090/metadata -> {“id":"61b5765e4cedfd0001d83976","message":"Successfully saved metadata."}
13. curl localhost:9090/metadata -> [{"lastUpdatedTs":"2021-12-12T04:11:20.97862","group":"voyager","name":"city","id":"61b5765e4cedfd0001d83976"},{"lastUpdatedTs":"2021-12-12T04:11:20.980149","group":"voyager","name":"city","id":"61b5761c4cedfd0001d83975"}]
14. Change mongo-db version from 4.2.6 to 4.2.7
15. kubectl port-forward svc/metadata-svc 9090:8080
16. curl localhost:9090/actuator/health -> {"status":"UP","details":{"diskSpace":{"status":"UP","details":{"total":62725623808,"free":51765358592,"threshold":10485760}},"mongo":{"status":"UP","details":{"version":"4.2.7"}}}}
17. curl localhost:9090/metadata -> [{"lastUpdatedTs":"2021-12-12T04:11:20.97862","group":"voyager","name":"city","id":"61b5765e4cedfd0001d83976"},{"lastUpdatedTs":"2021-12-12T04:11:20.980149","group":"voyager","name":"city","id":"61b5761c4cedfd0001d83975"}]

## Excercise 8
1. minikube ssh
2. mkdir mongo-pv
3. cd mongo-pv
4. mkdir data
5. kubectl apply -f Ex-8-PV-PVC/pv.yaml -> persistentvolume/mongo-pv created
6. kubectl apply -f Ex-8-PV-PVC/pvc.yaml -> persistentvolume/mongo-pvc created
7. kubectl apply -f Ex-8-PV-PVC/mongo.yaml -> deployment.apps/mongo created
8. kubectl apply -f Ex-8-PV-PVC/mongo-svc.yaml -> service/mongo created
9. kubectl apply -f Ex-8-PV-PVC/metadata.yaml -> deployment.apps/metadata created
10. kubectl apply -f Ex-8-PV-PVC/metadata-svc.yaml -> service/metadata-svc created
11. kubectl get pv                                               
    NAME       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM               STORAGECLASS   REASON   AGE
    mongo-pv   500Mi      RWO            Retain           Bound    default/mongo-pvc   manual                  12m
12. kubectl get pvc                                              
    NAME        STATUS   VOLUME     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
    mongo-pvc   Bound    mongo-pv   500Mi      RWO            manual         12m
13. kubectl port-forward svc/metadata-svc 9090:8080
14. curl localhost:9090/actuator/info -> {"app":{"version":"10.0.0","name":"Metadata Service","description":"Metadata service also known as config service. It hold the metadata/config required across different services."}}
15. curl localhost:9090/actuator/health -> {"status":"UP","details":{"diskSpace":{"status":"UP","details":{"total":62725623808,"free":52071821312,"threshold":10485760}},"mongo":{"status":"UP","details":{"version":"4.2.6"}}}}
16. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value":"Bangalore"}' http://127.0.0.1:9090/metadata -> {“id":"61b5827f4cedfd000175064d","message":"Successfully saved metadata."}
17. curl --header "Content-Type: application/json" --request POST --data '{"group": "voyager", "name": "city", "value”:”Pune”}’ http://127.0.0.1:9090/metadata -> {“id":"61b5827c4cedfd000175064c","message":"Successfully saved metadata."}
18. curl localhost:9090/metadata -> [{"group":"voyager","lastUpdatedTs":"2021-12-12T05:03:04.842512","name":"city","id":"61b5827f4cedfd000175064d"},{"group":"voyager","lastUpdatedTs":"2021-12-12T05:03:04.843953","name":"city","id":"61b5827c4cedfd000175064c"}]

## Exercise 9
1. kubectl apply -f Ex-9-Volumes/sc.yaml
2. kubectl apply -f Ex-9-Volumes/pvc.yaml
3. kubectl apply -f Ex-9-Volumes/mongo.yaml
4. kubectl apply -f Ex-9-Volumes/mongo-svc.yaml
5. kubectl apply -f Ex-9-Volumes/metdata.yaml
6. kubectl apply -f Ex-9-Volumes/metadata-svc.yaml
7. Insert 2 records
8. /metadata has 2 records now
9. kubectl delete deployment metadata
10. kubectl delete deployment mongo
11. kubectl delete svc metadata-svc
12. kubectl delete svc mongo
13. Kubectl delete pvc mongo-pvc
14. Volume - status detached
15. kubectl delete pv <pv-name>
16. Volume still there