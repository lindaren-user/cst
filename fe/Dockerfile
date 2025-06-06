FROM node:20 AS builder
WORKDIR /app
COPY . .
RUN npm i --registry=https://registry.npmmirror.com
RUN npm run build && ls -l build

FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80