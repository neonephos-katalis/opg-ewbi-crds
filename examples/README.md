# Custom Resource Examples

This directory provides representative Custom Resource (CR) manifests demonstrating valid configurations for each Custom Resource Definition (CRD) in the EWBI federation ecosystem.

The examples are grouped by component and federation technology:

## Directory Structure

```
examples/
└── operator/
    ├── rest/                              # Examples using REST technology (EWBI HTTP interface)
    │   ├── federation.yaml                # Federation configuration with REST options
    │   ├── zone.yaml                      # AvailabilityZone resource
    │   ├── image.yaml                     # Container image metadata
    │   ├── artefact.yaml                  # Application artefact package definition
    │   ├── appOnboard.yaml                # Application onboarding specification
    │   └── appDeploy.yaml                 # Application lifecycle deployment
    └── k8s/                               # Examples using direct Kubernetes technology
        ├── federation.yaml                # Federation configuration with k8sOptions
        ├── zone.yaml                      # AvailabilityZone resource
        ├── image.yaml                     # Image resource for k8s peering
        ├── artefact.yaml                  # Artefact resource for k8s peering
        ├── appOnboard.yaml                # Application onboarding resource
        ├── appDeploy.yaml                 # Application deployment resource
        └── consumptionmonitoring.yaml     # Resource consumption monitoring
```

## How to Test the Examples Locally

1. Create a local testing cluster (for example with Kind):
   ```bash
   kind create cluster --name ewbi-test
   ```

2. Install all CRDs:
   ```bash
   kubectl apply -f install.yaml
   ```

3. Ensure the target namespace exists (the sample files target the `guest` namespace by default):
   ```bash
   kubectl create namespace guest
   ```

4. Apply the REST examples:
   ```bash
   kubectl apply -f examples/operator/rest/
   ```

5. Verify that all custom resources are accepted by the Kubernetes API:
   ```bash
   kubectl -n guest get federations
   kubectl -n guest get availabilityzones
   kubectl -n guest get images
   kubectl -n guest get artefacts
   kubectl -n guest get applicationonboardings
   kubectl -n guest get applicationdeployments
   ```

6. Clean up resources:
   ```bash
   kubectl delete -f examples/operator/rest/
   kubectl delete namespace guest
   ```
