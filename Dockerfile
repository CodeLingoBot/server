FROM google/debian:wheezy

RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl build-essential ca-certificates git mercurial bzr
RUN mkdir /goroot && curl https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

RUN mkdir /gopath

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN mkdir -p  $GOPATH/src/github.com/awethome/server
WORKDIR  $GOPATH/src/github.com/awethome/server
ADD . $GOPATH/src/github.com/awethome/server

EXPOSE 80
ENTRYPOINT ["/gopath/src/github.com/awethome/server/scripts/eb_run.sh"]