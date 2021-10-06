package v1alpha1

import (
	"fmt"

	pipeline1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var depot map[schema.GroupVersionKind]client.Object

func init() {
	depot = make(map[schema.GroupVersionKind]client.Object)
	depot[corev1.SchemeGroupVersion.WithKind("PersistentVolume")] = &corev1.PersistentVolume{}
	depot[pipeline1beta1.SchemeGroupVersion.WithKind("TaskRun")] = &pipeline1beta1.TaskRun{}
}

func GetObj(gkv schema.GroupVersionKind) (client.Object, bool) {
	obg, ok := depot[gkv]
	return obg, ok
}

type Options func(client.Object)

func WithNamespace(namespace string) Options {
	return func(obj client.Object) {
		obj.SetNamespace(namespace)
	}
}

func WithAttachedGenerateName(name string) Options {
	return func(obj client.Object) {
		obj.SetGenerateName(fmt.Sprintf("%s-%s-", name, obj.GetName()))
		obj.SetName("")
	}
}
