# Dockerfile for React frontend
FROM node:16-alpine

WORKDIR /app

RUN apk update && apk add curl && apk cache clean

COPY package.json package-lock.json ./
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
