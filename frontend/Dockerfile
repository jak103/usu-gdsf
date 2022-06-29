FROM node:18-slim

WORKDIR /frontend

COPY package*.json ./

RUN npm config set unsafe-perm true

RUN npm install

RUN chown -R node:node /frontend/node_modules
