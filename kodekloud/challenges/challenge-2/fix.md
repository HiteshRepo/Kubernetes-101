- look at logs in : 
`cat /var/log/pods/kube-system_kube-apiserver-controlplane_<some-hash>/kube-apiserver/<current-number>.log`
- logs says:
```text
2023-08-09T23:43:52.146673209-04:00 stderr F I0810 03:43:52.146378       1 server.go:551] external host was not specified, using 192.5.36.6
2023-08-09T23:43:52.147960557-04:00 stderr F I0810 03:43:52.147811       1 server.go:165] Version: v1.27.0
2023-08-09T23:43:52.147982123-04:00 stderr F I0810 03:43:52.147851       1 server.go:167] "Golang settings" GOGC="" GOMAXPROCS="" GOTRACEBACK=""
2023-08-09T23:43:52.510294367-04:00 stderr F E0810 03:43:52.510084       1 run.go:74] "command failed" err="open /etc/kubernetes/pki/ca-authority.crt: no such file or directory"
```
- check `/etc/kubernetes/pki/` directory
```sh
ls /etc/kubernetes/pki/
apiserver.crt              
apiserver-etcd-client.key  
apiserver-kubelet-client.crt  
ca.crt  
etcd                
front-proxy-ca.key      
front-proxy-client.key  
sa.pub
apiserver-etcd-client.crt  
apiserver.key              
apiserver-kubelet-client.key  
ca.key  
front-proxy-ca.crt  
front-proxy-client.crt  
sa.key
```
there is no `ca-authority.crt` but there is `ca.crt` instead

- lets edit `kube-apiserver.yaml` at `/etc/kubernetes/manifests` and update flag `--client-ca-file` to `/etc/kubernetes/pki/ca.crt`

- view kube config file at `/root/.kube/config`
    the server address is: `https://controlplane:6433`
    replace `controlplane` this with ip mentioned in `--advertise-address` flag in `/etc/kubernetes/manifests/kube-apiserver.yaml`