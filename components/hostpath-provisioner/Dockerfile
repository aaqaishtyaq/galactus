FROM golang:1.16 as builder
RUN : \
    && go mod tidy \
    && CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o hostpath-provisioner . \
    && :

FROM scratch
COPY --from=builder /hostpath-provisioner /
CMD ["/hostpath-provisioner"]
