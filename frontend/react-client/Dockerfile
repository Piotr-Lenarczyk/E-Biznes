# Dockerfile for React frontend
FROM node:20-alpine

WORKDIR /app

RUN apk update && apk add curl && apk cache clean

COPY package.json package-lock.json ./
RUN npm install --ignore-scripts

COPY . .

EXPOSE 3000

RUN adduser -D -u 1001 appuser && chown -R appuser:appuser /app

USER appuser

CMD ["npm", "start"]
