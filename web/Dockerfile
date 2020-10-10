FROM node:13-alpine as development

RUN apk update && apk add bash

WORKDIR /home/node/app

COPY package.json .
COPY package-lock.json .

RUN npm i

COPY . .

EXPOSE 9000

CMD [ "npm", "start" ]