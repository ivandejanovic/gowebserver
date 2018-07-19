# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

#Get gorilla mux dependency
RUN go get github.com/gorilla/mux

#Get gowebserver
RUN go get github.com/ivandejanovic/gowebserver

#Set working directory
WORKDIR /go

# Run the goWebServer command by default when the container starts.
ENTRYPOINT ["/go/bin/gowebserver"]

# Document that the service listens on port 8000.
EXPOSE 8080