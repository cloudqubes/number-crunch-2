FROM golang:1.19.4-alpine3.17

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
# COPY go.sum .
# RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN go build -o /main

EXPOSE 8080

# Run
CMD [ "/main" ]