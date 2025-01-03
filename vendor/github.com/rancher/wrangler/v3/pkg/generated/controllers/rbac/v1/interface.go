/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	ClusterRole() ClusterRoleController
	ClusterRoleBinding() ClusterRoleBindingController
	Role() RoleController
	RoleBinding() RoleBindingController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) ClusterRole() ClusterRoleController {
	return generic.NewNonNamespacedController[*v1.ClusterRole, *v1.ClusterRoleList](schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "ClusterRole"}, "clusterroles", v.controllerFactory)
}

func (v *version) ClusterRoleBinding() ClusterRoleBindingController {
	return generic.NewNonNamespacedController[*v1.ClusterRoleBinding, *v1.ClusterRoleBindingList](schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "ClusterRoleBinding"}, "clusterrolebindings", v.controllerFactory)
}

func (v *version) Role() RoleController {
	return generic.NewController[*v1.Role, *v1.RoleList](schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "Role"}, "roles", true, v.controllerFactory)
}

func (v *version) RoleBinding() RoleBindingController {
	return generic.NewController[*v1.RoleBinding, *v1.RoleBindingList](schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "RoleBinding"}, "rolebindings", true, v.controllerFactory)
}
