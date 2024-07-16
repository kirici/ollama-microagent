FROM docker.io/library/node:lts-slim

RUN npm install -g @builder.io/micro-agent@v0.0.41

WORKDIR /home
