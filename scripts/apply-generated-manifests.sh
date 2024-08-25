kubectl apply -f /home/martinfleming/go/operators/test-operator/config/rbac/service_account.yaml
kubectl apply -f /home/martinfleming/go/operators/test-operator/config/rbac/role.yaml
kubectl apply -f /home/martinfleming/go/operators/test-operator/config/rbac/role_binding.yaml
kubectl apply -f /home/martinfleming/go/operators/test-operator/config/manager/manager.yaml
