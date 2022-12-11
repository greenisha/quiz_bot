FROM golang:1.19-alpine  
WORKDIR /app    
COPY go.mod ./   
COPY go.sum ./   
COPY *.go ./ 
RUN go build -o /go-server 
EXPOSE 8080
CMD ["/go-server"] 