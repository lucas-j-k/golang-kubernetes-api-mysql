## Setting up cluster and applying resources
- task start-cluster (creates Kind Cluster)
- task apply-config (creates namespaces)
- task deploy-mysql-storage
- task deploy-mysql
- task deploy-redis
- task deploy-api (deploys the actual app container)
- task deploy-contour
- task deploy-routes
- task deploy-migrator   (mysql needs to be running)

# Contour Ingress
- simple router, alternative to the nginx ingress.
- creates HttpProxy resources to route traffic to services
- accepts external traffic, routes to internal services.
- setup with kind:
setting up a valid kind cluster with config:
https://projectcontour.io/docs/1.25/guides/kind/
- the kind config adds extra port mappings to allow us to route traffic into the cluster without the load balancer ingress type - kind doesn't expose a load balancer. Kind config with extra port mappings is in ```./bash/cluster-setup.sh```

## MySQL
- MySQL deployed with a Persistent Volume Claim and Persistent Volume for storage

## SQL migrations
- SQL migrations are deployed as a simple K8s job which deploys an ephemeral container to run Goose migrations. There is a TTL defined on the job so it deletes itself after running
