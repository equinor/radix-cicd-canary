# Updating helm chart

1. Push updated `helm` chart to ACR. **PS: See note below if you have not used private ACR Helm Repos before.**

```
cd charts/radix-cicd-canary
helm package .
az acr helm push --name radixdev <tgz file>
az acr helm push --name radixprod <tgz file>
```