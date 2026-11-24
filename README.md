# OPG EWBI CRD (Custom Resource Definitions)

Questo repository contiene le definizioni Go e i manifesti YAML delle **Custom Resource Definitions (CRDs)** dell'interfaccia **GSMA OPG East/WestBound Interface (EWBI)**, parte del progetto IPCEI-CIS.

Il modulo è pensato per essere la **singola fonte di verità (Single Source of Truth)** per le risorse Kubernetes EWBI, consentendo ad altri repository (come `opg-ewbi-operator`, `opg-ewbi-api` o orchestratori esterni) di importare i tipi Go senza dipendere dall'intera implementazione dei controller.

---

## 📦 Struttura del Repository

```
.
├── api/
│   └── operator/
│       └── v1beta1/                       # Definizioni dei tipi Go e DeepCopy
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
│       ├── bases/                         # Manifesti YAML generati per ciascuna CRD
│       │   ├── opg.ewbi.katalis.com_applicationdeployments.yaml
│       │   ├── opg.ewbi.katalis.com_applicationonboardings.yaml
│       │   ├── opg.ewbi.katalis.com_artefacts.yaml
│       │   ├── opg.ewbi.katalis.com_availabilityzones.yaml
│       │   ├── opg.ewbi.katalis.com_consumptionmonitorings.yaml
│       │   ├── opg.ewbi.katalis.com_federations.yaml
│       │   └── opg.ewbi.katalis.com_images.yaml
│       ├── kustomization.yaml             # Configurazione Kustomize
│       └── kustomizeconfig.yaml
├── crds/                                  # Copia dei manifesti YAML per installazione diretta
├── install.yaml                           # File manifest unico con tutte le CRDs
├── Makefile                               # Target per generazione codice e manifesti
├── go.mod                                 # Modulo Go
└── PROJECT                                # Metadata Kubebuilder
```

---

## 🚀 Utilizzo

### 1. Installazione delle CRD nel cluster Kubernetes

#### Tramite `kubectl` diretto:
```bash
# Installazione da file cumulativo
kubectl apply -f install.yaml

# Oppure installazione dalla cartella crds/
kubectl apply -f crds/
```

#### Tramite `kustomize`:
```bash
kubectl apply -k config/crd/
```

#### Tramite `Makefile`:
```bash
make install
```

Per disinstallare:
```bash
make uninstall
```

---

### 2. Importazione del Modulo in Progetti Go

Per utilizzare i tipi delle CRD in un operatore o servizio Go:

```bash
go get github.com/neonephos-katalis/opg-ewbi-crd@latest
```

Nel codice Go:

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
    
    // Creazione o interazione con le risorse...
    fed := &opgcrds.Federation{
        // ...
    }
}
```

---

## 🛠️ Generazione Codice e Manifesti

Se vengono modificati i file dei tipi in `api/operator/v1beta1/*_types.go`:

1. Rigenerare il codice DeepCopy:
   ```bash
   make generate
   ```

2. Rigenerare i manifesti YAML e il file `install.yaml`:
   ```bash
   make manifests
   ```

---

## 📄 Licenza

Distribuito sotto licenza Apache 2.0. Vedere [LICENSE](LICENSE) e [NOTICE](NOTICE) per maggiori dettagli.
