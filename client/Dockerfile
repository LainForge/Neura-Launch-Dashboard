# Build stage

FROM node:16-alpine as build

WORKDIR /app

RUN npm install yarn

COPY package.json .

RUN yarn install

COPY . .

RUN yarn build

# Serve with nginx

FROM nginx:1.23-alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=build /app/dist .

ENTRYPOINT [ "nginx", "-g", "daemon off;" ]