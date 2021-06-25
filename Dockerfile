# build stage
FROM golang:1.16 as build

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /go/src/lister
COPY . .

RUN go get -v
RUN go vet -v
RUN make installwithvars

# run stage
FROM busybox as run

ENV TEMPLATE_PATH "./templates"

COPY --from=build /go/bin/lister /
COPY --from=build /go/src/lister/frontend /

EXPOSE 15200
EXPOSE 9392

CMD ["/lister"]
