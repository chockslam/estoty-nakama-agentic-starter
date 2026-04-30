FROM heroiclabs/nakama-pluginbuilder:3.37.0 AS builder

ENV CGO_ENABLED=1 \
    GO111MODULE=on

WORKDIR /backend

COPY go.mod go.sum main.go ./

RUN go mod vendor
RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM heroiclabs/nakama:3.37.0

COPY --from=builder /backend/backend.so /nakama/data/modules/backend.so
COPY local.yml /nakama/data/local.yml
