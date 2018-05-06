Create config map containing tasks:

    kubectl create configmap locust-tasks --from-file=./locust-tasks/tasks.py

Deploy locust master & workers:

    kubectl apply -f k8s-locust.yaml

Deploy test app:

    kubectl apply -f locust-test-app/k8s-deployment.yaml

Navigate to Locust master web ui to run tests.
