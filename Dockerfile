FROM golang:1.20 AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o event-hunters . 

# KETIKA CGO_ENABLED = 0 DIHAPUS, EXECUTEABLE TIDAK DITEMUKAN DI DIRECTORY
# CGO_ENABLED = 0 untuk membuat hasil binerny tidak akan terhubung ke library C apapun
# With CGO_ENABLED=0 you got a staticaly-linked binary 
#(see: https://en.wikipedia.org/wiki/Static_build) 
# so it will run without any external dependencies 
# (you can buld your dockers from 'scratch' image) 
# Like that: https://github.com/s0rg/microapp/blob/master/Dockerfile

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /app/event-hunters /
COPY ./public /public
COPY ./views /views  
WORKDIR /
EXPOSE 8080
CMD [ "./event-hunters" ]