FROM node:16.0.0-alpine

WORKDIR /app
COPY package*.json ./

RUN npm install
COPY ./src ./src

CMD ["node", "./src/index.js"]

EXPOSE 3000
