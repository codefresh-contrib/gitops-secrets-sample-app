# Gitops with secrets

This is single application that uses [Bitnami sealed secrets](https://github.com/bitnami-labs/sealed-secrets) for
password and certificates.

## How to run locally

`go run .`

then visit http://localhost:8080 in your browser

## How to build and run container

Run

 *  `docker build . -t my-app` to create a container image 
 *  `docker run -p 8080:8080 my-app` to run it

 then visit http://localhost:8080 in your browser

You can find prebuilt images at [https://hub.docker.com/r/kostiscodefresh/gitops-secrets-sample-app/tags](https://hub.docker.com/r/kostiscodefresh/gitops-secrets-sample-app/tags)

See the [documentation page](https://codefresh.io/docs/docs/yaml-examples/examples/gitops-secrets/) for more details.

