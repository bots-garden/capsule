var http = require('http');
var fs = require('fs');

var server = http.createServer(function (req, res) {
    //console.log(req.url);
    if (req.method === "GET") {

        if(req.url === "/html") {
            res.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
            res.end("<h1>Hello World</h1>");
        }
        if(req.url === "/json") {
            res.writeHead(200, { "Content-Type": "application/json; charset=utf-8" });
            res.end(JSON.stringify({
                "name": "Bob Morane",
                "age": 42
            }));
        }                
    } else if (req.method === "POST") {
        //console.log(JSON.stringify(req.))

        let chunks = [];
        req.on("data", (chunk) => {
          chunks.push(chunk);
        });

        req.on("end", () => {
          const data = Buffer.concat(chunks);
          console.log(data.toString());
          res.writeHead(200, { "Content-Type": "application/json; charset=utf-8" });

          res.end(JSON.stringify({
            "name": "Jane Doe",
            "age": 24,
            "parameters": JSON.parse(data.toString())
          }));
        });      

    }

}).listen(process.env.HTTP_PORT || 3000);