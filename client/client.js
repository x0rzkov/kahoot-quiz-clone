//Import dependencies
const path = require('path');
const http = require('http');
const express = require('express');

const publicPath = path.join(__dirname, './public');
const app = express();

const server = http.createServer(app);

app.use(express.static(publicPath));

// Starting server on port 4242
server.listen(4242, () => {
    console.log("Server started on port 4242");
});

