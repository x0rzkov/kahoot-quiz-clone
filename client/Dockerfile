FROM mhart/alpine-node:latest

WORKDIR /app

COPY package.json /app/package.json
RUN npm install

COPY . .

EXPOSE 4242

VOLUME ["/app"]

CMD ["node", "client.js"]
