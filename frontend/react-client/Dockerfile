FROM node:20-alpine

WORKDIR /app

RUN apk update && apk add curl && apk cache clean && adduser -D -u 1001 appuser

COPY --chown=appuser:appuser package*.json ./

RUN chown -R appuser:appuser /app

USER appuser

RUN npm install --ignore-scripts

COPY --chown=appuser:appuser . .

EXPOSE 3000

CMD ["npm", "start"]

