package forge

import (
	clv1alpha2 "github.com/netgroup-polito/CrownLabs/operators/api/v1alpha2"
)

const (
	labelManagedByKey = "crownlabs.polito.it/managed-by"
	labelInstanceKey  = "crownlabs.polito.it/instance"
	labelWorkspaceKey = "crownlabs.polito.it/workspace"
	labelTemplateKey  = "crownlabs.polito.it/template"
	labelTenantKey    = "crownlabs.polito.it/tenant"

	labelManagedByValue = "instance"
)

// InstanceLabels receives in input a set of labels and returns the updated set depending on the specified template,
// along with a boolean value indicating whether an update should be performed.
func InstanceLabels(labels map[string]string, template *clv1alpha2.Template) (map[string]string, bool) {
	update := false
	if labels == nil {
		labels = map[string]string{}
		update = true
	}

	update = updateLabel(labels, labelManagedByKey, labelManagedByValue) || update
	update = updateLabel(labels, labelWorkspaceKey, template.Spec.WorkspaceRef.Name) || update
	update = updateLabel(labels, labelTemplateKey, template.Name) || update

	return labels, update
}

// InstanceObjectLabels receives in input a set of labels and returns the updated set depending on the specified instance.
func InstanceObjectLabels(labels map[string]string, instance *clv1alpha2.Instance) map[string]string {
	if labels == nil {
		labels = map[string]string{}
	}
	labels[labelManagedByKey] = labelManagedByValue
	labels[labelInstanceKey] = instance.Name
	labels[labelTemplateKey] = instance.Spec.Template.Name
	labels[labelTenantKey] = instance.Spec.Tenant.Name

	return labels
}

// InstanceSelectorLabels returns a set of selector labels depending on the specified instance.
func InstanceSelectorLabels(instance *clv1alpha2.Instance) map[string]string {
	return map[string]string{
		labelInstanceKey: instance.Name,
		labelTemplateKey: instance.Spec.Template.Name,
		labelTenantKey:   instance.Spec.Tenant.Name,
	}
}

// updateLabel configures a map entry to a given value, and returns whether a change was performed.
func updateLabel(labels map[string]string, key, value string) bool {
	if labels[key] != value {
		labels[key] = value
		return true
	}
	return false
}