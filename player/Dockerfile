FROM golang:1.19

WORKDIR $GOPATH/src/github.com/Planducteur/project-go/tree/go-projet/player

EXPOSE 8090
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/player

CMD ["player"]
