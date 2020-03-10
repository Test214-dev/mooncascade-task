FROM golang:1.13

WORKDIR /mooncascade-task

COPY ./ /mooncascade-task

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

RUN go get github.com/Altoros/gorm-goose/cmd/gorm-goose

RUN go get -u github.com/swaggo/swag/cmd/swag

ENTRYPOINT ./start.sh