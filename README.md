# ![logo](https://github.com/open-policy-agent/opa/blob/master/logo/logo-144x144.png)  O.P.A Examples

A collection of examples on how to use Open Policy Agent.

Included

- Multiple policies written in the rego language
- Using shared rego libraries
- JWT Validation and support
- Running tests in O.P.A
- Mocking data for tests
- Live docker-compose example using the bundle API
- Live reloading of policies and data
- GCS support for bundles

## Policies and Data

## Testing Policies

## Bundle API

[O.P.A Bundle API](https://www.openpolicyagent.org/docs/latest/management/#bundles)

The bundle API is a great way to transport policies and data to your O.P.A server without having to use the HTTP API. The data is loaded into O.P.A in real time without requiring a restart. Caching can be implemented using the `Etag` header to prevent the same bundle from being downloaded. Having multiple bundles allows the user to separate data from the policies which might have different development flows. The bundle API can be configured in O.P.A as shown below.

O.P.A Config ([opa.yaml](./config/opa.yaml))
```yaml
services:
  - name: bundle-api
    url: http://bundle-api:8080/

bundles:
  access:
    service: bundle-api
    resource: data/bundle.tar.gz
    polling:
      min_delay_seconds: 10
      max_delay_seconds: 20
  endpoints:
    service: bundle-api
    resource: rego/bundle.tar.gz
    polling:
      min_delay_seconds: 10
      max_delay_seconds: 20
```

### Architecture

### Bundle Layout

### GCS

## Docker Compose
Included in the root directory is a fully functional docker-compose example to test out and try! simply build the bundle-api (`docker-compose build`) and run (`docker-compose up`). You can edit the rego or data in real time and when you are ready to upload it to the bundle-api run `./bundle.sh`. No data will be uploaded if the tests are failing.

* Build the bundle-api

`docker-compose build`

* Run

`docker-compose up`

* Add / Edit Data 

`./bundle.sh`

* Useful Endpoints

- `localhost:8181/v1/data/users` (retrieve all users)
- `localhost:8181/v1/data/users/{user_name}` (retrieve a single users)

- `localhost:8181/v1/data/endpoints` (endpoints policy)
- `localhost:8181/v1/data/tables` (tables policy)


## TLDR Example

