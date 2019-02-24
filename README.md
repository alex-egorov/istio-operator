# Istio-operator

Istio-operator is a Kubernetes operator to deploy and manage [Istio](https://istio.io/) resources for a Kubernetes cluster.

## Overview

[Istio](https://istio.io/) is an open platform to connect, manage, and secure microservices and it is emerging as the `standard` for building service meshes on Kubernetes. It is built out on multiple components and a rather complex deployment scheme (around 14 Helm subcharts and 50+ CRDs). Installing, upgrading and operating these components requires deep understanding of Istio and Helm (the standard/supported way of deploying [Istio](https://istio.io/)).

The goal of the **Istio-operator** is to automate and simplify these and enable popular service mesh use cases (multi cluster federation, canary releases, resource reconciliation, etc) by introducing easy higher level abstractions.

### Motivation

At [Banzai Cloud](https://banzaicloud.com) we are building a Kubernetes distribution and platform - [Pipeline](https://github.com/banzaicloud/pipeline) and operate Istio clusters for our customers. While we were comfortably operating Istio using the standard Helm deployments on 6 cloud providers and on-premise with [Pipeline](https://github.com/banzaicloud/pipeline), recently our customers were asking for multi-cloud service mesh deployments. This required lots of configurations, manual interventions during scaling or removing clusters from the mesh and become an operational burden. [Pipeline](https://github.com/banzaicloud/pipeline) automates the whole Kubernetes experience (from creating clusters, centralized logging, federated monitoring, multi-dimensional autoscaling, disaster recovery, security scans, etc) and we needed a way to `automagically` operate Istio.

At the same time there is a huge interest in the Istio community for an [operator](https://github.com/istio/istio/issues/9333), but due to resource constraints and the need of supporting Helm, building one it was discarded. There were several initiatives to simplify Istio:

- [Istio Operator for Kubernetes](https://github.com/istio/istio/issues/9333)
- [Operator](https://github.com/istio/istio/pull/10015)
- [Initial implementation of Galley registers the CRDs](https://github.com/istio/istio/pull/10120)
- [Handle upgrades with an istio-init chart](https://github.com/istio/istio/pull/10562)

however, none of these gave a full solution to **automate** the Istio experience and make it consumable for the wider audience.

Our motivation is to build an open source solution and a community which drives the innovation and features of the operator.

If you are willing to kickstart your Istio experience using Pipeline, check out the free developer beta:
<p align="center">
  <a href="https://beta.banzaicloud.io">
  <img src="https://camo.githubusercontent.com/a487fb3128bcd1ef9fc1bf97ead8d6d6a442049a/68747470733a2f2f62616e7a6169636c6f75642e636f6d2f696d672f7472795f706970656c696e655f627574746f6e2e737667">
  </a>
</p>


## Installation

The operator (`master` branch) installs the 1.0.5 version of Istio, and requires kubectl 1.13.0 and can run on Minikube v0.33.1+ and Kubernetes 1.10.0+.

As a pre-requisite it needs a Kubernetes cluster (you can create one using [Pipeline](https://github.com/banzaicloud/pipeline).

1. Set `KUBECONFIG` pointing towards your cluster
2. Run `make vendor`
3. Run `make deploy` (deploys the operator in the `istio-system` namespace to the cluster)
4. Set your Istio configurations in a  Kubernetes custom resource (sample: `config/samples/operator_v1beta1_config.yaml`) and run this command to deploy the Istio components:

```
kubectl create -n istio-system -f config/samples/operator_v1beta1_config.yaml
```

## Development

Check out the [developer docs](docs/developer.md).
Check out the [developer docs](docs/developer.md).

## Issues, feature requests and roadmap

Please note that the Istio operator is under heavy development and new releases might introduce breaking changes. We are striving to keep backward compatibility as much as possible while adding new features at a fast pace. Issues, new features or bugs are tracked on the projects [GitHub page]() - please feel free to add yours!

Some of the significant features and future items from the roadmap:

Released:

- [x] - Install Istio
- [x] - Enable global mTLS
- [x] - Federation, flat

Under development (next release)

- [x] - Federation, gateway
- [x] - Istio 1.1.0 support
- [x] - Configurable node affinity

Short term roadmap

- [ ] - Canary releases
- [ ] - Servicegraph / Kiali / Certmanager

## Contributing

If you find this project useful here's how you can help:

- Send a pull request with your new features and bug fixes
- Help new users with issues they may encounter
- Support the development of this project and star this repo!

## License

Copyright (c) 2017-2019 [Banzai Cloud, Inc.](https://banzaicloud.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.