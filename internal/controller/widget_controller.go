/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	enterprisev1alpha1 "github.com/martinflemingdev/test-operator/api/v1alpha1"
)

// WidgetReconciler reconciles a Widget object
type WidgetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=enterprise.mftest.com,resources=widgets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=enterprise.mftest.com,resources=widgets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=enterprise.mftest.com,resources=widgets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Widget object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *WidgetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	widget := &enterprisev1alpha1.Widget{}
	err := r.Get(ctx, req.NamespacedName, widget)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			logger.Info("Widget resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		logger.Error(err, "Failed to get Widget")
		return ctrl.Result{Requeue: true}, err
	}
	logger.Info("Widget acquired!")

	// extract spec data
	h := widget.Spec.Height
	w := widget.Spec.Width
	d := widget.Spec.Depth
	volume := h * w * d

	message := fmt.Sprintf("Extracted widget specs: Height=%v, Width=%v, Depth=%v", h, w, d)
	logger.Info(message)

	// check if desired state matches actual state
	if volume == widget.Status.Volume {
		logger.Info("Status matches spec, no update required")
		return ctrl.Result{}, nil
	}

	// update status data if needed
	widget.Status.Volume = volume
	_ = r.Client.Status().Update(ctx, widget)
	message = fmt.Sprintf("The volume is: %d", volume)
	logger.Info(message)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WidgetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&enterprisev1alpha1.Widget{}).
		Complete(r)
}
