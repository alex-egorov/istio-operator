/*
Copyright 2019 Banzai Cloud.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodeagent

import (
	"github.com/banzaicloud/istio-operator/pkg/util"
	"github.com/go-logr/logr"
	"github.com/goph/emperror"
	"sigs.k8s.io/controller-runtime/pkg/client"

	istiov1beta1 "github.com/banzaicloud/istio-operator/pkg/apis/istio/v1beta1"
	"github.com/banzaicloud/istio-operator/pkg/k8sutil"
	"github.com/banzaicloud/istio-operator/pkg/resources"
)

const (
	componentName          = "nodeagent"
	serviceAccountName     = "istio-nodeagent-service-account"
	clusterRoleName        = "istio-nodeagent-cluster-role"
	clusterRoleBindingName = "istio-nodeagent-cluster-role-binding"
	daemonSetName          = "istio-nodeagent"
)

var nodeAgentLabels = map[string]string{
	"app": "istio-nodeagent",
}

var labelSelector = map[string]string{
	"istio": "nodeagent",
}

type Reconciler struct {
	resources.Reconciler
}

func New(client client.Client, config *istiov1beta1.Istio) *Reconciler {
	return &Reconciler{
		Reconciler: resources.Reconciler{
			Client: client,
			Config: config,
		},
	}
}

func (r *Reconciler) Reconcile(log logr.Logger) error {
	log = log.WithValues("component", componentName)

	log.Info("Reconciling")

	var nodeAgentDesiredState k8sutil.DesiredState
	if util.PointerToBool(r.Config.Spec.NodeAgent.Enabled) {
		nodeAgentDesiredState = k8sutil.DesiredStatePresent
	} else {
		nodeAgentDesiredState = k8sutil.DesiredStateAbsent
	}

	for _, res := range []resources.Resource{
		r.serviceAccount,
		r.clusterRole,
		r.clusterRoleBinding,
		r.daemonSet,
	} {
		o := res()
		err := k8sutil.Reconcile(log, r.Client, o, nodeAgentDesiredState)
		if err != nil {
			return emperror.WrapWith(err, "failed to reconcile resource", "resource", o.GetObjectKind().GroupVersionKind())
		}
	}

	log.Info("Reconciled")

	return nil
}
