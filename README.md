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
├── beetle/
│   ├── base/
│   │   └── config.yaml
│   ├── dev/
│   │   └── config.yaml
│   └── prd/
│       └── config.yaml
├── sonar/
│   ├── base/
│   │   └── config.yaml
│   ├── dev/
│   │   └── config.yaml
│   └── prd/
│       └── config.yaml
└── tiger/
    ├── base/
    │   └── config.yaml
    ├── dev/
    │   └── config.yaml
    └── prd/
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
