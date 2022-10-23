# KnativeCon North America 2022 - Platform Building Demo

This repository contains the step-by-step instructions to run the KnativeCon NA Demo. 

The Platform that we built uses the following projects that you need to install in a Kubernetes Cluster: 
- Crossplane
- Knative Serving
- VCluster (will be automatically installed)
- ArgoCD

In this demo you will be able to request new Development Environments to the platform API.

## Prerequisites and Installation 

This tutorial creates interact with Kubernetes clusters as well as install Helm Charts hence the following tools are needed: 
- [Install `kubectl`](https://kubernetes.io/docs/tasks/tools/)
- [Install `helm`](https://helm.sh/docs/intro/install/) 
- [Install `docker`](https://docs.docker.com/engine/install/)

For this demo we will create a Kubernetes Clusters to host the Platform tools which  create Development Environments in separate namespaces using VCluster and one namespace for our Production Environment. While this tutorial uses  [KinD](https://kind.sigs.k8s.io/) for simplicity, we encourage people to try the tutorial on real Kubernetes Clusters. 


> Note: This tutorial has been tested on [GCP](https://cloud.google.com/gcp)  and using separate clusters for the platform and the production ?environment. [You can always get free credits here](https://github.com/learnk8s/free-kubernetes).


- [Installing Command-line tools](installing-clis.md)
- [Create a Platform Cluster & Install Tools](platform-cluster.md)
  

### Configuring our Platform cluster

For this demo, our platform will enable development teams to request new `Environment`s.

These `Environment`s can be configured differently depending what the teams needs to do. For this demo, we have created a [Crossplane Composition](https://crossplane.io/docs/v1.9/concepts/composition.html) that uses [VCluster](https://www.vcluster.com/) to create one virtual cluster per Development environment requested. This enable the teams with their own isolated cluster to work on features without clashing with other teams work. 

For this to work we need to create the Custom Resource Definition (CRD) that defines the APIs for creating new `Environment`s and the Crossplanae Composition that define the resources that will be created every time that a new `Environment` resource is created. 

Let's apply the Crossplane Composition and our **Environment Custom Resource Definition (CRD)** into the Platform Cluster:
```
kubectl apply -f crossplane/environment-resource-definition.yaml
kubectl apply -f crossplane/composition-devenv.yaml
```

The Crossplane Composition that we have defined and configured in our Platform Cluster uses the Crossplane Helm Provider to create a new VCluster for every `Environment` with `type: development`. The VCluster will be created inside the Platform Cluster, but it will provide its own isolated Kubernetes API Server for the team to interact with. 

The VCluster created for each development `Environment` is using the Knative Serving Plugin to enable teams to use Knative Serving inside the virtual cluster, but without having Knative Serving installed. The VCluster Knative Serving plugin shares the Knative Serving installation in the host cluster with all the virtual clusters.

Now we are ready to request environments, deploy our applications/function and promote them to production. 
