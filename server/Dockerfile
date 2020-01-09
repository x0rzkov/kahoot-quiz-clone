FROM mhart/alpine-node:latest

WORKDIR /app
COPY package.json /app/package.json
RUN npm install

COPY . .

EXPOSE 3000

VOLUME ["/app"]

CMD ["node", "server/server.js"]
