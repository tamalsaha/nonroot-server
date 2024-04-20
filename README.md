# nonroot-server

```bash
KO_DOCKER_REPO=ghcr.io/appscodeci ko build --base-import-paths
```

```bash
kubectl create ns demo
kubectl label ns demo pod-security.kubernetes.io/enforce=restricted
```

```bash
KO_DOCKER_REPO=ghcr.io/appscodeci ko resolve --base-import-paths -f deployment.yaml | kubectl apply -f -
```


Background Reading:

- https://medium.com/@olivier.gaumond/why-am-i-able-to-bind-a-privileged-port-in-my-container-without-the-net-bind-service-capability-60972a4d5496

- https://github.com/kubernetes/enhancements/issues/2763

- https://github.com/kubernetes/kubernetes/issues/56374
