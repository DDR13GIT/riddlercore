FROM golang:alpine AS builder

RUN apk add --no-cache --update git gcc g++ bash openssh tzdata

# Set go env
ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPRIVATE="magic.pathao.com"
ENV CGO_ENABLED=1

# Take SSH private key from argument
# create ssh dir and known_hosts file
# copy ssh private key, add magic as known host and
# use ssh over https for cloning

ARG SSH_KEY

RUN mkdir /root/.ssh/
RUN touch /root/.ssh/known_hosts
RUN echo "$SSH_KEY" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa
RUN git config --global url."git@magic.pathao.com:".insteadOf https://magic.pathao.com/ \
    && ssh-keyscan magic.pathao.com >> ~/.ssh/known_hosts

COPY . $GOPATH/src/magic.pathao.com/platform/riddlercore
WORKDIR $GOPATH/src/magic.pathao.com/platform/riddlercore

RUN chmod +x ./build.sh
RUN ./build.sh

RUN mv ./riddlercore /go/bin/

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates openssl && apk add --no-cache tzdata
COPY --from=0 /go/bin/riddlercore /usr/local/bin/riddlercore

ENTRYPOINT ["riddlercore"]
