# Flux Developer

This repository demonstrates how Developer teams can manage configurations for
their microservice applications.

In this example, each micro service application configuration schema and values
are stored here.

The platform team owns and manages Kustomizations and Kubernetes manifests
required to deploy each microservice here:
https://github.com/cgradwohl/flux_platform

Platform side: you create Kustomization files with configMapGenerator to
generate final ConfigMaps

## Example repo layout

```bash
flux_developer/
│
├─ global-config.yaml
│
├─ beetle/
│  ├─ dev/
│  │  ├─ service-config.yaml
│  │  └─ kustomization.yaml
│  └─ prd/
│     ├─ service-config.yaml
│     └─ kustomization.yaml
├─ sonar/
│  ├─ dev/
│  │  ├─ service-config.yaml
│  │  └─ kustomization.yaml
│  └─ prd/
│     ├─ service-config.yaml
│     └─ kustomization.yaml
└─ tiger/
   ├─ dev/
   │  ├─ service-config.yaml
   │  └─ kustomization.yaml
   └─ prd/
      ├─ service-config.yaml
      └─ kustomization.yaml

```

`global-config.yaml`

- global values across all environments
- this configuration is present in all configmaps for all servives in all
  environments

`<service>/<environment>/service-config.yaml`

- the environment-specific configuration for a given service

## ConfigMaps

The resulting ConfigMap for a given service:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  global-config.yaml: |
    ...
  service-config.yaml: |
    ...
```

## Service Image

To build the service image, first point minikube to the correct minikube
profile:

```bash
eval $(minikube -p dev docker-env)
```

Now, verfify you are evaled into dev:

```bash
docker info | grep Name
```

Next, build the image:

```bash
cd flux_developer/services/beetle

docker build -t beetle:latest .
```

Now, verify:

```bash
docker images | grep beetle

```

### How does the deployment template reference the image?

The runtime deployment has:

```yaml
image: app:latest
imagePullPolicy: IfNotPresent
```

Each service has a kustomization:

```yaml
images:
  - name: app
    newName: beetle
    newTag: latest
```

Therefore the final rendered deployment contains:

```yaml
image: beetle:latest
imagePullPolicy: IfNotPresent
```
