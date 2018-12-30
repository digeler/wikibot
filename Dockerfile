############################
# STEP 1 build executable binary
############################
#FROM dinorg/wikibotdemo
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/
# Fetch dependencies.;
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build .
EXPOSE 4445
ENTRYPOINT ["./myapp"]

############################
# STEP 2 build a small image
############################
#FROM scratch
# Copy our static executable.
#COPY --from=builder /go/bin/bot /go/bin/bot
# Run the hello binary.

#ENTRYPOINT ["/go/bin/bot"]