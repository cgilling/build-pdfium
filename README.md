# build-pdfium

My attempts to build pdfium and write a application that uses it

## Dockerfile

Currently the Dockerfile encodes the steps that were needed in order to build the pdfium library. Check it out for more info.

## Building the test apps

The easiest way to do this is to mount the code as a shared volume into the docker container built above.
```sh
docker build -t build-pdfium .
docker run -i -t --name build_pdfium -v $PWD/pdfium_test:/pdfium_test build-pdfium
```
You will need to install `yum install -y make` to build the `c++` application. In order to build the `go` application you will need to `yum install -y golang`, and then `export GOPATH=/pdfium_test/go`

Future version will likely encode this into the `Dockerfile` itself.
