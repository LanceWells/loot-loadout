# Kubeapps Setup

## About

These steps may be performed to get kubeapps setup on a given env.

## Steps:

These steps are transcribed from [this URL](https://kubeapps.dev/).

### Step 1: Install Kubeapps (for demo purposes)

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install -n kubeapps --create-namespace kubeapps bitnami/kubeapps
```

Step 2: Create a demo credential with which to access Kubeapps and Kubernetes
```bash
kubectl create --namespace default serviceaccount kubeapps-operator
kubectl create clusterrolebinding kubeapps-operator --clusterrole=cluster-admin --serviceaccount=default:kubeapps-operator
kubectl apply -f ./services/kubeapps/secret.yaml
```

Step 3: Get the API token
```bash
make get-kubeapps-secret
```
