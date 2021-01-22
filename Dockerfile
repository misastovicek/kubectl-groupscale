FROM golang:1.15.5-alpine AS build
WORKDIR /src
COPY . .
RUN go get
RUN go build -o kubectl-groupscale ./

FROM microsoft/azure-cli AS bin
COPY --from=build /src/kubectl-groupscale /kubectl-groupscale
