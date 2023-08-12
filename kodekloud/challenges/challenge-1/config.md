- nano ~/.kube/config

- new user 'martin'
```yaml
- name: martin
  user: 
    client-certificate: /root/martin.crt
    client-key: /root/martin.key
```

- new context
```yaml
- context:
    cluster: kubernetes
    user: martin
  name: developer
  namespace: development
```

--------
steps:
- k config use-context kubernetes-admin@kubernetes
- k apply -f role.yaml
- k apply -f binding.yaml
- k  config use-context developer
- k apply -f pvc.yaml
- k apply -f pod.yaml
- k apply -f svc,yaml
