/*
Copyright 2023 The Kubernetes Authors.

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

package logical

// AnnotatedObject is an interface that defines a Kubernetes object that has annotations.
// It's a slimmed down version of metav1.Object and defined here to reduce the amount
// of dependencies required to use this package.
type AnnotatedObject interface {
	GetAnnotations() map[string]string
}

// Annotation defines the logical annotation that can be attached to a Kubernetes object
// to determine the logical cluster name the object belongs to.
const Annotation = "logical.sigs.x-k8s.io/cluster"

// NameFromObject returns the logical cluster name from the given Kubernetes object, if any.
func NameFromObject(obj AnnotatedObject) Name {
	return Name(obj.GetAnnotations()[Annotation])
}
