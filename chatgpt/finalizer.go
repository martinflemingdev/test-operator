package chatgpt


// You're correct that MergeFrom() may be interpreting the finalizers field replacement as "adding" foregroundDeletion, even though you’re just preserving it. The key issue is that strategic merge patches don’t work well for CRDs, and MergeFrom() internally generates a strategic merge patch.

// Alternative Fix: Use RawPatch() with JSON Merge Patch
// Instead of using MergeFrom(), explicitly patch only the finalizers field using RawPatch(), ensuring no other fields are touched.

// Fixed Code (Using RawPatch() Instead of MergeFrom())
// go
// Copy
// Edit
import (
    "context"
    "encoding/json"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    apierrors "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/apimachinery/pkg/types"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *WidgetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    widget := &v1alpha1.Widget{}
    if err := r.Get(ctx, req.NamespacedName, widget); err != nil {
        if apierrors.IsNotFound(err) {
            return ctrl.Result{}, nil
        }
        return ctrl.Result{}, err
    }

    finalizerName := "finalizer.widget.example.com"

    if controllerutil.ContainsFinalizer(widget, finalizerName) {
        // Remove only our finalizer while keeping any system finalizers like "foregroundDeletion"
        updatedFinalizers := []string{}
        for _, f := range widget.Finalizers {
            if f != finalizerName { // Keep all except ours
                updatedFinalizers = append(updatedFinalizers, f)
            }
        }

        // Explicitly update only the finalizers field via JSON Merge Patch
        patchData := map[string]interface{}{
            "metadata": map[string]interface{}{
                "finalizers": updatedFinalizers,
            },
        }
        patchBytes, err := json.Marshal(patchData)
        if err != nil {
            return ctrl.Result{}, fmt.Errorf("failed to marshal finalizers patch: %w", err)
        }

        // Apply the patch with RawPatch and MergePatchType (not StrategicMerge)
        err = r.Client.Patch(ctx, widget, client.RawPatch(types.MergePatchType, patchBytes))
        if err != nil {
            return ctrl.Result{}, fmt.Errorf("failed to patch finalizers: %w", err)
        }

        return ctrl.Result{}, nil
    }

    return ctrl.Result{}, nil
}