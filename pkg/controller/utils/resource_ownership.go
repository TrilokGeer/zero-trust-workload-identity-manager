package utils

import (
	"fmt"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
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
	return ResourceConflictError(existing.GetNamespace(), existing.GetName())
}

// IsResourceConflictOnCreate checks if a Create error is an AlreadyExists error,
// which indicates a naming conflict with a pre-existing resource not visible in
// the operator's label-filtered cache.
func IsResourceConflictOnCreate(err error) bool {
	return kerrors.IsAlreadyExists(err)
}

// ResourceConflictError returns a formatted error for a resource conflict.
func ResourceConflictError(namespace, name string) error {
	if namespace != "" {
		return fmt.Errorf("resource %s/%s already exists but is not managed by the operator", namespace, name)
	}
	return fmt.Errorf("resource %s already exists but is not managed by the operator", name)
}
