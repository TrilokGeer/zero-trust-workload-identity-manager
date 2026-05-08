package utils

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CheckResourceConflict verifies that an existing resource is managed by the operator
// by checking the managed-by label. Returns an error if the resource exists but does not
// have the operator's managed-by label, indicating a naming conflict with a pre-existing resource.
func CheckResourceConflict(existing client.Object) error {
	labels := existing.GetLabels()
	if labels != nil && labels[AppManagedByLabelKey] == AppManagedByLabelValue {
		return nil
	}
	ns := existing.GetNamespace()
	name := existing.GetName()
	if ns != "" {
		return fmt.Errorf("resource %s/%s already exists but is not managed by the operator", ns, name)
	}
	return fmt.Errorf("resource %s already exists but is not managed by the operator", name)
}
