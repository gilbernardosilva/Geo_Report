FROM node:14 AS build

WORKDIR /app

COPY package*.json ./

RUN npm install --silent --ignore-scripts

COPY public ./public
COPY src ./src

ARG REACT_APP_API_URL

ENV REACT_APP_API_URL=$REACT_APP_API_URL

RUN npm run build

FROM node:14-alpine

WORKDIR /app

RUN npm install -g pm2

COPY package*.json ./

RUN npm install --production

COPY --from=build /app/build ./build

COPY src/server.js .

EXPOSE 80

CMD ["pm2-runtime", "server.js"]
