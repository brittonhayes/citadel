FROM golang

RUN mkdir -p /go/src//home/peshka/GolandProjects/citadel

ADD . /go/src//home/peshka/GolandProjects/citadel

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /home/peshka/GolandProjects/citadel/threatintel/cmd  -watch /home/peshka/GolandProjects/citadel/threatintel
