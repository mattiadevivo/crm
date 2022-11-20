# build stage
FROM golang:1.19 as build-img

ARG GIT_USER
ARG GIT_TOKEN

WORKDIR /go/src/app

COPY ./ .

RUN echo "machine github.com" > /root/.netrc
RUN echo "  login $GIT_USER" >> /root/.netrc
RUN echo "  password $GIT_TOKEN" >> /root/.netrc
ENV GOPRIVATE="github.com/mattiadevivo/"
RUN go mod tidy
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/app/cmd/crm
RUN go build -a -ldflags="-w -s" -o crm-backend

# final stage
FROM alpine:3.17
WORKDIR /app
COPY --from=build-img /go/src/app/cmd/crm/crm-backend .

CMD ["/app/crm-backend"]