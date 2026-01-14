# Welcome to [Slidev](https://github.com/slidevjs/slidev)!

To start the slide show:

- `pnpm install`
- `pnpm dev`
- visit <http://localhost:3030>

Edit the [slides.md](./slides.md) to see the changes.

Learn more about Slidev at the [documentation](https://sli.dev/).

## Following the workshop

In order to follow the workshop, please make sure you have the following prerequisites installed:

- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [minikube](https://minikube.sigs.k8s.io/docs/start/)

kubectl cheat sheet: https://kubernetes.io/docs/reference/kubectl/cheatsheet/

## Debugging

At any point in time, you can create an interactive busybox Pod to help you debug networking issues:

```bash
kubectl run -i --rm --tty busybox --image=busybox --restart=Never -- sh
```

## Setting up the Ingress controller

Install the Helm chart for the Cloudflare Tunnel Ingress Controller.
First add the Helm repository and update your local Helm chart repository cache:

```shell
helm repo add strrl.dev https://helm.strrl.dev
helm repo update
```

Let's forward to dashboards minikube provides us. Enables some addons:

```shell
minikube addons enable dashboard
minikube addons enable metrics-server
```

Install the Cloudflare Tunnel Ingress Controller Helm chart.

```shell
helm upgrade --install --wait \
  -n cloudflare-tunnel-ingress-controller --create-namespace \
  cloudflare-tunnel-ingress-controller \
  strrl.dev/cloudflare-tunnel-ingress-controller \
  --set=cloudflare.apiToken="$CLOUDFLARE_API_TOKEN" \
  --set cloudflare.accountId="$CLOUDFLARE_ACCOUNT_ID" \
  --set cloudflare.tunnelName="minikube-demo" 
```

Now create an ingress:

```yaml
apiVersion: v1
items:
- apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: dashboard
    namespace: kubernetes-dashboard
  spec:
    ingressClassName: cloudflare-tunnel
    rules:
    - host: k8s-dashboard.twaslowski.com
      http:
        paths:
        - backend:
            service:
              name: kubernetes-dashboard
              port:
                number: 80
          path: /
          pathType: Prefix
```