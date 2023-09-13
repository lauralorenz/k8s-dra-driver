# Dynamic Resource Allocation (DRA) for NVIDIA GPUs in Kubernetes

This DRA resource driver is currently under active development and not yet
designed for production use.
We will continually be force pushing over `main` until we have something more stable.
Use at your own risk.

A document and demo of the DRA support for GPUs provided by this repo can be found below:
|                                                                                                                          Document                                                                                                                          |                                                                                                                                                                   Demo                                                                                                                                                                   |
|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| [<img width="300" alt="Dynamic Resource Allocation (DRA) for GPUs in Kubernetes" src="https://drive.google.com/uc?export=download&id=12EwdvHHI92FucRO2tuIqLR33OC8MwCQK">](https://docs.google.com/document/d/1BNWqgx_SmZDi-va_V31v3DnuVwYnF2EmN7D-O_fB6Oo) | [<img width="300" alt="Demo of Dynamic Resource Allocation (DRA) for GPUs in Kubernetes" src="https://drive.google.com/uc?export=download&id=1UzB-EBEVwUTRF7R0YXbGe9hvTjuKaBlm">](https://drive.google.com/file/d/1iLg2FEAEilb1dcI27TnB19VYtbcvgKhS/view?usp=sharing "Demo of Dynamic Resource Allocation (DRA) for GPUs in Kubernetes") |

## Demo

This section describes using `kind` to demo the functionality of the NVIDIA GPU DRA Driver.

First since we'll launch kind with GPU support, ensure that the following prerequisites are met:
1. `kind` is installed. See the official documentation [here](https://kind.sigs.k8s.io/docs/user/quick-start/#installation).
1. Ensure that the NVIDIA Container Toolkit is installed on your system. This
   can be done by following the instructions
   [here](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html).
1. Configure the NVIDIA Container Runtime as the **default** Docker runtime:
   ```bash
   sudo nvidia-ctk runtime configure --runtime=docker --set-as-default
   ```
1. Restart Docker to apply the changes:
   ```bash
   sudo systemctl restart docker
   ```
1. Set the `accept-nvidia-visible-devices-as-volume-mounts` option to `true` in
   the `/etc/nvidia-container-runtime/config.toml` file to configure the NVIDIA
   Container Runtime to use volume mounts to select devices to inject into a
   container.

We start by first cloning this repository and `cd`ing into its `demo`
subdirectory. All of the scripts and example Pod specs used in this demo are
contained here, so take a moment to browse through the various files and see
what's available:

```
git clone https://github.com/NVIDIA/k8s-dra-driver.git
cd k8s-dra-driver/demo
```

### Setting up the infrastructure
First, create a `kind` cluster to run the demo:
```bash
./create-cluster.sh
```


From here we will build the image for the example resource driver:
```bash
./build-dra-driver.sh
```

This also makes the built images available to the `kind` cluster.

We now install the NVIDIA GPU DRA driver:
```
./install-dra-driver.sh
```

This should show two pods running in the `nvidia-dra-driver` namespace:
```
$ kubectl get pods -n nvidia-dra-driver
NAMESPACE           NAME                                       READY   STATUS    RESTARTS   AGE
nvidia-dra-driver   nvidia-dra-controller-6bdf8f88cc-psb4r     1/1     Running   0          34s
nvidia-dra-driver   nvidia-dra-plugin-lt7qh                    1/1     Running   0          32s
```

### Run the examples by following the steps in the demo script
Finally, you can run the various examples contained in the `demo` folder.
The files `demo/DEMO.sh` shows the full script of the demo you can walk through.
```console
cat demo/DEMO.sh
...
```

Where the running the first three examples should produce output similar to the following:
```console
$ kubectl apply --filename=gpu-test{1,2,3}.yaml
...

```
```console
$ kubectl get pod -A
NAMESPACE           NAME                                       READY   STATUS    RESTARTS   AGE
gpu-test1           pod1                                       1/1     Running   0          34s
gpu-test1           pod2                                       1/1     Running   0          34s
gpu-test2           pod                                        2/2     Running   0          34s
gpu-test3           pod1                                       1/1     Running   0          34s
gpu-test3           pod2                                       1/1     Running   0          34s
...

```
```console
$ kubectl logs -n gpu-test1 -l app=pod
GPU 0: A100-SXM4-40GB (UUID: GPU-662077db-fa3f-0d8f-9502-21ab0ef058a2)
GPU 0: A100-SXM4-40GB (UUID: GPU-4cf8db2d-06c0-7d70-1a51-e59b25b2c16c)

$ kubectl logs -n gpu-test2 pod --all-containers
GPU 0: A100-SXM4-40GB (UUID: GPU-79a2ba02-a537-ccbf-2965-8e9d90c0bd54)
GPU 0: A100-SXM4-40GB (UUID: GPU-79a2ba02-a537-ccbf-2965-8e9d90c0bd54)

$ kubectl logs -n gpu-test3 -l app=pod
GPU 0: A100-SXM4-40GB (UUID: GPU-4404041a-04cf-1ccf-9e70-f139a9b1e23c)
GPU 0: A100-SXM4-40GB (UUID: GPU-4404041a-04cf-1ccf-9e70-f139a9b1e23c)
```

### Cleaning up the environment

Running
```
$ ./delete-cluster.sh
```
will remove the cluster created in the preceding steps.

<!--
TODO: This README should be extended with additional content including:

## Information for "real" deployment including prerequesites

This may include the following content from the original scripts:
```
set -e

export VERSION=v0.1.0

REGISTRY=nvcr.io/nvidia/cloud-native
IMAGE=k8s-dra-driver
PLATFORM=ubi8

sudo true
make -f deployments/container/Makefile build-${PLATFORM}
docker tag ${REGISTRY}/${IMAGE}:${VERSION}-${PLATFORM} ${REGISTRY}/${IMAGE}:${VERSION}
docker save ${REGISTRY}/${IMAGE}:${VERSION} > image.tgz
sudo ctr -n k8s.io image import image.tgz
```

## Information on advanced usage such as MIG.

This includes setting configuring MIG on the host using mig-parted. Some of the demo scripts included
in ./demo/ require this.

```
cat <<EOF | sudo -E nvidia-mig-parted apply -f -
version: v1
mig-configs:
half-half:
   - devices: [0,1,2,3]
      mig-enabled: false
   - devices: [4,5,6,7]
      mig-enabled: true
      mig-devices: {}
EOF
```
-->