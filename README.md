# OPG EWBI Custom Resource Definitions (opg-ewbi-crd)

This repository contains the Go type definitions and generated YAML manifests for the Custom Resource Definitions (CRDs) defined in the GSMA Operator Platform Group (OPG) East/WestBound Interface (EWBI) specifications.

This repository serves as the single source of truth for all Kubernetes resources used across the EWBI federation ecosystem. It enables downstream components such as `opg-ewbi-operator`, `opg-ewbi-api`, and third-party orchestrators to import the Go API types without depending on the controller runtime implementations.

---

## Repository Structure

```
opg-ewbi-crd/
├── api/
│   └── operator/
│       └── v1beta1/                       # Go type definitions and DeepCopy implementations
│           ├── applicationdeployment_types.go
│           ├── applicationonboarding_types.go
│           ├── artefact_types.go
│           ├── availabilityzone_types.go
│           ├── consumptionmonitoring_types.go
│           ├── federation_types.go
│           ├── groupversion_info.go
│           ├── image_types.go
│           └── zz_generated.deepcopy.go
├── config/
│   └── crd/
│       ├── bases/                         # Generated YAML manifests for individual CRDs
│       │   ├── opg.ewbi.katalis.com_applicationdeployments.yaml
│       │   ├── opg.ewbi.katalis.com_applicationonboardings.yaml
│       │   ├── opg.ewbi.katalis.com_artefacts.yaml
│       │   ├── opg.ewbi.katalis.com_availabilityzones.yaml
│       │   ├── opg.ewbi.katalis.com_consumptionmonitorings.yaml
│       │   ├── opg.ewbi.katalis.com_federations.yaml
│       │   └── opg.ewbi.katalis.com_images.yaml
│       ├── kustomization.yaml             # Kustomize entrypoint
│       └── kustomizeconfig.yaml
├── crds/                                  # Flat directory with individual CRD YAMLs
├── install.yaml                           # Single concatenated manifest for direct installation
├── Makefile                               # Automation targets for code and manifest generation
├── PROJECT                                # Kubebuilder project metadata
├── go.mod                                 # Go module definition
└── README.md
```

---

## Testing and Installing the CRDs

### 1. Prerequisites

Before installing the CRDs, ensure you have:
- A Kubernetes cluster version 1.24 or later (such as a local Kind or Minikube cluster).
- `kubectl` configured with access to the target cluster.

To start a clean local testing cluster with Kind:
```bash
kind create cluster --name ewbi-test
kubectl config use-context kind-ewbi-test
```

### 2. Installation Options

You can install the CRDs into your cluster using any of the following approaches:

#### Option A: Single Manifest File (Recommended for Quick Tests)
Apply all CRDs in a single command using the concatenated file:
```bash
kubectl apply -f install.yaml
```

#### Option B: Directory of Manifests
Apply all manifests from the `crds/` directory:
```bash
kubectl apply -f crds/
```

#### Option C: Kustomize
Apply the CRDs using Kustomize:
```bash
kubectl apply -k config/crd/
```

#### Option D: Makefile
Use the provided Makefile target:
```bash
make install
```

### 3. Verifying CRD Installation

Confirm that all seven EWBI CRDs have been registered and are established in the cluster:
```bash
kubectl get crd | grep "opg.ewbi.katalis.com"
```

Expected output:
```
applicationdeployments.opg.ewbi.katalis.com     <date>
applicationonboardings.opg.ewbi.katalis.com     <date>
artefacts.opg.ewbi.katalis.com                  <date>
availabilityzones.opg.ewbi.katalis.com          <date>
consumptionmonitorings.opg.ewbi.katalis.com     <date>
federations.opg.ewbi.katalis.com                <date>
images.opg.ewbi.katalis.com                     <date>
```

Inspect the schema validation rules of a specific CRD:
```bash
kubectl describe crd federations.opg.ewbi.katalis.com
```

### 4. Testing Resource Creation

Verify that the Kubernetes API server accepts and validates custom resources defined by these CRDs:

Create a test namespace:
```bash
kubectl create namespace test-federation
```

Apply a test `Federation` manifest:
```bash
cat <<EOF | kubectl apply -f -
apiVersion: opg.ewbi.katalis.com/v1beta1
kind: Federation
metadata:
  name: test-federation-host
  namespace: test-federation
spec:
  federationData:
    relationType: HOST
    technologyType: REST
    origOPFederationId: op-test-01
    origOPCountryCode: IT
    initialDate: "2026-01-01T10:00:00Z"
    restOptions:
      partnerStatusLink: "http://partner.example.com/callback"
EOF
```

Verify that the resource was stored correctly:
```bash
kubectl -n test-federation get federations
kubectl -n test-federation get federation test-federation-host -o yaml
```

Clean up the test resource:
```bash
kubectl -n test-federation delete federation test-federation-host
kubectl delete namespace test-federation
```

### 5. Uninstalling the CRDs

To remove all EWBI CRDs from the cluster:
```bash
kubectl delete -f install.yaml
# or:
make uninstall
```

---

## Importing the Go Module

Downstream Go applications can import the CRD types directly:

```bash
go get github.com/neonephos-katalis/opg-ewbi-crd@latest
```

Example usage in Go code:

```go
package main

import (
    "k8s.io/apimachinery/pkg/runtime"
    utilruntime "k8s.io/apimachinery/pkg/util/runtime"
    
    opgcrds "github.com/neonephos-katalis/opg-ewbi-crd/api/operator/v1beta1"
)

func main() {
    scheme := runtime.NewScheme()
    utilruntime.Must(opgcrds.AddToScheme(scheme))
    
    // Custom resource instantiation
    fed := &opgcrds.Federation{
        Spec: opgcrds.FederationSpec{
            // ...
        },
    }
}
```

---

## Code and Manifest Generation

When Go types in `api/operator/v1beta1/*_types.go` are modified:

1. Regenerate `zz_generated.deepcopy.go`:
   ```bash
   make generate
   ```

2. Regenerate YAML manifests in `config/crd/bases/`, `crds/`, and `install.yaml`:
   ```bash
   make manifests
   ```

3. Validate Go formatting and vetting:
   ```bash
   make fmt
   make vet
   ```

---

## Automated Upstream Synchronization

This repository includes a GitHub Actions workflow (`.github/workflows/sync-upstream.yml`) configured to synchronize CRD definitions from the `main` branch of `opg-ewbi-operator`.

- Triggers:
  - Scheduled run every 6 hours.
  - Manual trigger via `workflow_dispatch`.
  - Repository webhook dispatch via `repository_dispatch`.
- Behavior:
  - Clones the latest `main` branch of `opg-ewbi-operator`.
  - Copies Go type definitions and generated YAML manifests.
  - Regenerates `install.yaml` and executes `go mod tidy`.
  - If any differences are detected, it automatically creates a new branch and submits a Pull Request against `main`.

---

## License

This project is licensed under the Apache License 2.0. See the LICENSE and NOTICE files for details.
