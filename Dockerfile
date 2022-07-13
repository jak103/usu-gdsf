FROM node:18-slim AS client
ENV NODE_ENV production
WORKDIR /client
COPY ./client/ ./
RUN npm install
RUN npm run build

FROM golang:1.18 AS server
WORKDIR /server/
COPY ./server ./
#RUN go build -o usu-gdsf .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o usu-gdsf .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /usu-gdsf
COPY --from=server /server/usu-gdsf /usu-gdsf/usu-gdsf
COPY --from=client /client/dist /client/dist
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/usu-gdsf/usu-gdsf"]