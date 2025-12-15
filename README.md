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
tenant-repo/
├── base/
│   ├── beetle/
│   │   └── config.yaml
│   ├── sonar/
│   │   └── config.yaml
│   └── tiger/
│       └── config.yaml
├── dev/
│   ├── beetle/
│   │   └── config.yaml
│   ├── sonar/
│   │   └── config.yaml
│   └── tiger/
│       └── config.yaml
└── production/
    ├── beetle/
    │   └── config.yaml
    ├── sonar/
    │   └── config.yaml
    └── tiger/
        └── config.yaml
```

`base/<service>/config.yaml`

- default values for all environments

`dev/<service>/config.yaml` and `prd/<service>/config.yaml`

- environment-specific overrides

## Example files

`base/beetle/config.yaml`

```yaml
jwt:
  hmacSecret: "PUT_YOUR_HMAC_SECRET_HERE"
  publicKey: "PUBLIC_KEY"
  scheme: "Bearer"
  useHmac: false
redis:
  authMode: "AWS_IAM"
  awsRegion: "us-east-1"
  awsReplicationGroupId: "live-beetle"
  awsUserId: "live-beetle"
  endpoint: "rediss://clustercfg.live-beetle.7cdemb.use1.cache.amazonaws.com:6379"
  integrationCheck: true
  password: ""
  serverMode: "CLUSTER"
  user: ""
```

`dev/beetle/config.yaml`

```yaml
jwt:
  useHmac: true
redis:
  awsRegion: "us-west-2"
```

## Example kustomizations

In this example repoisotry setup, a differnet, corresponding repository defines
all the kustomizations and kubernetes manifests required for the microservice
deployment.

`dev/kustomization.yaml`

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: apps
resources:
  - github.com/your-org/flux_developer//base?ref=main
configMapGenerator:
  - name: dev-beetle-config
    files:
      - github.com/your-org/tenant-repo//base/beetle/config.yaml
      - github.com/your-org/tenant-repo//dev/beetle/config.yaml
```

`prd/kustomization.yaml`

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: apps
resources:
  - github.com/your-org/flux_developer//base?ref=main
configMapGenerator:
  - name: prd-beetle-config
    files:
      - github.com/your-org/tenant-repo//base/beetle/config.yaml
      - github.com/your-org/tenant-repo//prd/beetle/config.yaml
```
