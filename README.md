alertmanager-webhook-example
============================

`alertmanager-webhook-example` is a simple example of a webhook receiver for the [Prometheus
Alertmanager](https://prometheus.io/docs/alerting/alertmanager/). It expect to receive
aleter messages in POST bodies to `/alerts` in JSON in the format described in
the [receiever webhook docs](https://prometheus.io/docs/alerting/configuration/#webhook-receiver-<webhook_config>)

Usage
=====

`alertmanager-webhook-example` stores an array of received alert messages in memory.

```
$ ./alertmanager-webhook-example
Usage of alertmanager-webhook-example
  -addr string
    	address to listen for webhook (default ":8080")
  -cap int
    	capacity of the simple alerts store (default 64)
exit status 2
```

It expects alerts to be POSTed in JSON to `/alerts`. An HTTP GET to `/alerts` will return
the alerts in memory.

Building
========
`alertmanager-webhook-example` has no external dependencies.  Clone this repo into
your GOPATH and run `go build .` in the top level of this repo.

prebuilt Docker images are available at https://quay.io/repository/bakins/alertmanager-webhook-example

LICENSE
=======

See [LICENSE](./LICENSE)
