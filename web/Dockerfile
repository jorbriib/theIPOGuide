FROM node:13-alpine as base

WORKDIR /home/node/app

COPY package.json .
COPY package-lock.json .

RUN npm ci

COPY . .

FROM base as development

EXPOSE 9000

CMD [ "npm", "start" ]

FROM base as builder

ARG API_URL
ARG APP_URL
ARG RECAPTCHA_SITE_KEY

RUN npm run build

FROM nginx:1.19-alpine as production

WORKDIR /usr/share/nginx/html

COPY --from=builder /home/node/app/public .
COPY --from=builder /home/node/app/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]