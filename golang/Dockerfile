# Start from golang base image
FROM golang:1.13
ENV CGO_ENABLED 1

# Set the current working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/rakyll/gotest
RUN go get github.com/codegangsta/gin

# Copy sources to the working directory
COPY . .

# Run the Go app
#ARG project
#ENV PROJECT $project
#CMD /go/bin/gin -d ${PROJECT} run main.go
#run ls -la
#RUN cat entrypoint.sh
CMD ["/bin/sh", "entrypoint.sh"]
#RUN /bin/sh -c 'while true ;do wait ;done'
#CMD ["/bin/sh", "tail", "-f", "/dev/null"]
#CMD ["/bin/sh", "-c", "'while true ;do wait ;done'"]
