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

### Debugging

At any point in time, you can create an interactive busybox Pod to help you debug networking issues:

```bash
kubectl run -i --rm --tty busybox --image=busybox --restart=Never -- sh
```
