FROM golang:1.20

WORKDIR /usr/src/app/SafePasswordApi

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./        
RUN go build -o safePassword 

# EXPOSE ${API_PORT}

CMD ["./safePassword"]
