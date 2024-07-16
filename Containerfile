FROM docker.io/library/node:22-alpine

RUN npm install -g @builder.io/micro-agent@v0.0.41

WORKDIR /home
