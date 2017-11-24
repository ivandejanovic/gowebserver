# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY src/goWebServerBasic /go/src/goWebServerBasic/

# Copy the local templates to container
COPY tmpl/ /go/tmpl/

# Copy the static content to container
COPY static/ /go/static/

#Get gorilla mux dependency
RUN go get github.com/gorilla/mux

# Build the goWebServer command inside the container.
RUN go install goWebServerBasic

# Run the goWebServer command by default when the container starts.
ENTRYPOINT ["/go/bin/goWebServerBasic"]

# Document that the service listens on port 8000.
EXPOSE 8000