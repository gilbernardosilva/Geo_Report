FROM node:14 AS build

WORKDIR /app

COPY package*.json ./

RUN npm install --silent --ignore-scripts

COPY public ./public
COPY src ./src

ARG REACT_APP_API_URL

ENV REACT_APP_API_URL=$REACT_APP_API_URL

RUN npm run build

FROM nginx:alpine

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN mkdir -p /var/cache/nginx/client_temp && \
    chown -R appuser:appgroup /var/cache/nginx /var/run /var/log/nginx

RUN chmod -R 755 /var/cache/nginx/client_temp

RUN touch /var/run/nginx.pid && \
    chown -R appuser:appgroup /var/run/nginx.pid && \
    chmod 644 /var/run/nginx.pid

COPY --from=build /app/build /usr/share/nginx/html

RUN chown -R appuser:appgroup /usr/share/nginx/html

USER appuser

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]