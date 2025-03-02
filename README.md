# Setup Instruction

## Using Localstack

If you are working with Localstack, ensure you have the following up and running

1. Setup a python virtual environment to install `tflocal` and `awslocal`.
(These utilities are wrappers for `terraform` and `awscli` respectively. They enable you
connect to localstack without needing to manually configure the endpoint.)

```shell
python3 -m venv venv
source venv/bin/activate
pip install --upgrade pip
pip install awscli-local
pip install terraform-local
```

2. Start up local stack using the docker compose file at `localstack-setup/compose.yaml`

```shell
cd localstack-setup
docker compose up -d
```

## Setup Atlantis Using Helm

1. Add the helm repo for Atlantis

```shell
helm repo add runatlantis https://runatlantis.github.io/helm-charts
```

1. Get and update the values file for the Atlantis helm charts

```shell
helm inspect values runatlantis/atlantis > values.yaml
```

2. Install atlantis with the updated `values.yaml` file

```shell
helm install atlantis runatlantis/atlantis -f values.yaml
```

3. Setup a webhook on Github (or your preferred Git server)
   Ensure you set a webhook secret

4. Ensure you create a kubernetes secret to handle github token and webhook secret

```shell
echo -n "yourtoken" > token
echo -n "yoursecret" > webhook-secret
kubectl create secret generic atlantis-vcs --from-file=token --from-file=webhook-secret
```

## References

- [Atlantis Helm Install](https://www.runatlantis.io/docs/deployment.html#kubernetes-helm-chart)
