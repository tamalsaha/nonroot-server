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
