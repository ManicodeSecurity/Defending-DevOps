FROM golang:1.9 

RUN mkdir -p /go/src/unshorten-api/

WORKDIR /go/src/unshorten-api/

ADD . /go/src/unshorten-api/

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN dep ensure -vendor-only

WORKDIR /go/src/unshorten-api
RUN go install .
RUN go build

# RUN useradd -ms /bin/bash nonrootuser
# USER nonrootuser

CMD ["./unshorten-api"]




