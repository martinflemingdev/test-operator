#!/usr/bin/bash
kubectl delete -f /home/martinfleming/go/operators/test-operator/config/rbac/role_binding.yaml
kubectl delete -f /home/martinfleming/go/operators/test-operator/config/rbac/service_account.yaml
kubectl delete -f /home/martinfleming/go/operators/test-operator/config/rbac/role.yaml
kubectl delete -f /home/martinfleming/go/operators/test-operator/config/manager/manager.yaml
kubectl delete -f /home/martinfleming/go/operators/test-operator/config/manager/ns.yaml
kubectl delete -f /home/martinfleming/go/operators/test-operator/config/crd/bases/enterprise.mftest.com_widgets.yaml