FROM mhart/alpine-node:latest

WORKDIR /app
COPY package.json /app/package.json

# COPY . .
RUN yarn install

EXPOSE 3000

# VOLUME ["/app"]

CMD ["node", "server.js"]
