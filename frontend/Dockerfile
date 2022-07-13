FROM node:18-slim

WORKDIR /frontend

COPY package*.json ./

RUN npm install

RUN chown -R node:node /frontend/node_modules