# go cluster test env

### setting up
- run ./bash/cluster-setup.sh to init a kind cluster with the registry attached
- registry runs at localhost:5001.
- Taskfile contains scripts to build and push the service images to the registry on :5001.
- go images are just alpine + binary so are small.

ok so after the ingress etc.. has been applied, and the go service_one has been deployed, we can access the service one endpoint at localhost/ping  (default port 80 used by browser)


now we need to set it up so that we can route to different containers based on different urls



Things to add
- 2 other microservices
- configmap
- loggin - prometheus
- kubernetes-dashboard
- redis?
- proper apigateway?
- custom operator??
- cronjob
- database?
- some kind of frontend?
- kubernetes security stuff? rbac?
- tilt?

