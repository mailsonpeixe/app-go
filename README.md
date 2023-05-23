# Project to learn kubernetes objects.

Based in Go language show an index page in endpoint http

### Prerequisite for Linux users

Docker, Kind(or K3d), GO


### Versions that you pay attention:

GO Version: **1.20**

Kubernetes Version: > **1.22**



#### For localdev run:

`./localbuild.sh`

#### To build and push:

if necessary, adjust url repository and run `./buildpush.sh`

_Note: you have need permissions to upload the image Docker to repository._


### Deploy kubernetes local

Run script `./start.sh`

NOTE: in this case, we create a namespace called `lcd-25023` to app, check the script start.sh

Check service NodePort `kubectl get services`

### Stop app in kubernetes local

Run script `./stop.sh`

### Deploy Cloud Kubernetes
NOTE: Check the context that you are connected, `kubectx` and `kubens` are useful commands
NOTE 2: in this case, we create a namespace called `lcd-25023`  and create a Ingress to app, check the script start.sh

Run script `./startcloud.sh`

and to stop and clear objects created, run `./stop.sh`

