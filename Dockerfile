FROM golang:1.20

WORKDIR /usr/src/app/SafePasswordApi

# ## install tools
RUN go install github.com/swaggo/swag/cmd/swag@latest


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./      

# Generate the api documentation
RUN swag init

# project build
RUN go build -o safePassword 

EXPOSE ${API_PORT}

# ## Run the application
CMD ["./safePassword"]
