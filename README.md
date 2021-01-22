# Kubectl Groupscale

This app is meant as an example plugin for **[kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)**. It can be used to scale all [Kubernetes](https://kubernetes.io/) `deployments` with specified label key/value pair.

>**Note:** The app is definitely not perfect and is meant just for demonstration purposes only!

## Usage

```bash
kubectl groupscale -label "scale=yes" -count 2
```

The above command will scale all the deployments across all the namespaces to specified count.
