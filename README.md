# Kubernetes Cloud Controller Manager for Advanced Hosting


## Install
Please note that every kubelet must run with `--cloud-provider=external`. This is to ensure that the kubelet is aware that it must be initialized by the cloud controller manager before it is scheduled any work.

### Install via kubectl

Create a secret:
```
apiVersion: v1
kind: Secret
metadata:
  name: advancedhosting
  namespace: kube-system
stringData:
  token: "TOKEN"
```

Create a configmap:
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: advancedhosting
  namespace: kube-system
data:
  private_network.number: "PRIVATE-NETWORK-NUMBER"
  datacenter.slug: "DATACENTER-SLUG"
```
Deploy CCM:
```
kubectl apply -f https://raw.githubusercontent.com/advancedhosting/advancedhosting-cloud-controller-manager/master/deploy/advancedhosting-ccm-{VERSION}.yaml
```

### Install via Helm
```
helm repo add ah-ccm https://advancedhosting.github.io/advancedhosting-cloud-controller-manager
helm repo update

# Put your values to these variables:
NETWORK="NET14520581"
DATACENTER="ams1"

helm install ccm ah-ccm/ah-ccm --set privateNetworkNumber=$NETWORK --set datacenterSlug=$DATACENTER
```