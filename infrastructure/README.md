# Infrastructure

Tools and packages that are used to run a deployment of Taskcluster.

To generate Helm charts for Taskcluster services, run `k8s/bin/helmit.py`. Use your own secrets management to fill out the values listed in `k8s/sample-helm-values.yaml`. 
