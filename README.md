Create config map containing tasks:

    kubectl create configmap locust-tasks --from-file=./locust-tasks/tasks.py

Deploy locust master & workers:

    kubectl apply -f k8s-locust.yaml
