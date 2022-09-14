FROM golang:1.19

WORKDIR $GOPATH/src/github.com/Planducteur/project-go/tree/go-projet/blockchain_service

EXPOSE 8091

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/blockchain

CMD ["blockchain", "--host", "0.0.0.0"]
