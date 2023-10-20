FROM node:16-alpine

WORKDIR /app

RUN npm install -g pnpm

COPY package.json .

RUN pnpm install

COPY . .

CMD [ "pnpm", "run", "dev", "--host", "--port", "3000" ]
