FROM ubuntu:16.04

LABEL maintainer="Hama <wowlsh93@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app/go-skeleton

#ADD ./SomansaRootCA.crt /usr/local/share/ca-certificates/SomansaRootCA.crt
#RUN  chmod 644 /usr/local/share/ca-certificates/SomansaRootCA.crt && update-ca-certificates


# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o goo ./cmd/goo

CMD ["./goo aapp start"]
