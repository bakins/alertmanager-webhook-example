FROM scratch
MAINTAINER Brian Akins <brian@akins.org>
COPY alertmanager-webhook-example.linux /alertmanager-webhook-example
ENTRYPOINT [ "/alertmanager-webhook-example" ]
