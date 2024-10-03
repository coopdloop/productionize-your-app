# Setup kubernetes to deploy docker images, and monitor applications

### shortcuts

        alias k='kubectl'

### links
kind (kubernetes in docker)
https://kind.sigs.k8s.io/


## steps:

1. Create a kubernetes cluster "monitoring". This creates a cluster with kubernetes version: 1.28.13

        kind create cluster --name monitoring --image kindest/node:v1.28.13

2. Validate running cluster "monitoring"

        k get nodes

3. Deploy kube-prometheus to our cluster. Deploy CRD (custom resource definitions)

        k apply --server-side -f ./monitoring/prometheus/kubernetes/main/manifests/setup

- Run this command to wait for kube-prometheus to be established

        k wait \
            --for condition=Established \
            --all CustomResourceDefinition \
            --namespace=monitoring

- Once established, apply the whole remaining manifests

        k apply -f ./monitoring/prometheus/kubernetes/main/manifests/

4. Let's check the install

        k -n monitoring get pods

5. Let's view the Grafana dashboard

        k -n monitoring port-forward svc/grafana 3000

6. Grafana Datasource fix

Now for some reason, the Prometheus data source in Grafana does not work out the box. To fix it, we need to change the service endpoint of the data source.

To do this, edit manifests/grafana-dashboardDatasources.yaml and replace the datasource url endpoint with http://prometheus-operated.monitoring.svc:9090

7. We'll need to patch that and restart Grafana

        k apply -f ./monitoring/prometheus/kubernetes/main/manifests/grafana-dashboardDatasources.yaml
        k -n monitoring delete po <grafana-pod>
        k -n monitoring port-forward svc/grafana 3000

8. Let's view the Prometheus UI

        k -n monitoring port-forward svc/prometheus-operated 9090

9. Create our own Prometheus

        k apply -n monitoring -f ./monitoring/prometheus/prometheus.yaml

10. View our prometheus `prometheus-applications-0` instance

        k -n monitoring get pods

11. Checkout our own prometheus UI

        k -n monitoring port-forward prometheus-applications-0 9090

    We will notice that service-discovery and targets are empty. This is because our application does not have a configured servicemonitor for it's kube deployment. We need to deploy a servicemonitor. The servicemonitor is used to tell the prometheus instance where our application monitoring endpoint is.

    A general rule of thumb is: One servicemonitor per microservice/application.

12. Deploy a servicemonitor to default namespace (where all apps are running)

        k -n default apply -f ./monitoring/prometheus/servicemonitor.yaml

13. Validate servicemonitor exists

        k -n default get servicemonitors


14. Verify our own prometheus is aware of the new servicemonitors

        k -n monitoring port-forward prometheus-applications-0 9090

    Right now our servicemonitor is not able to find a service to scrape, however the servicemonitor apply did in fact work and prometheus can see discover it.

15. Deploy the example app via deployment.yaml and service.yaml from the application folder

        k -n default apply -f ./go-app/

    Now we should see a target in the Prometheus Targets page.
