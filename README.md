# Rest API Boilerplate Go [Gin, JWT, minIO]
_just my backup - simple boilerplate-rest-api go_

####  set environment variable

`export MONGO_URI="mongodb://127.0.0.1:27017"`

`export APPLICATION_NAME="My Simple Rest Api"`

`export JWT_SIGNATURE_KEY="SECRETABC"`

`export MINIO_ENDPOINT="127.0.0.1:9000"`

`export MINIO_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE"`

`export MINIO_SECRET_ACCESS_KEY_ID="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"`

or run
`source env.sh`

#### Running Project
`go run . -bind=:5000`
_default port :8080_

##### TODO:
[x] dockerize
[x] dockerize