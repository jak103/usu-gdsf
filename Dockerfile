FROM node:18-slim AS client
ENV NODE_ENV production
WORKDIR /frontend
COPY ./frontend/ ./
RUN npm install
RUN npm run build

FROM golang:1.18 AS backend
WORKDIR /backend/
COPY ./backend ./
#RUN go build -o usu-gdsf .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o usu-gdsf .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /usu-gdsf
COPY --from=backend /backend/usu-gdsf /usu-gdsf/usu-gdsf
COPY --from=client /frontend/dist /frontend/dist
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/usu-gdsf/usu-gdsf"]