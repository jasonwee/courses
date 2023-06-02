const dotenv = require("dotenv");
dotenv.config();

var path = require("path");
const express = require("express");
const mockAPIResponse = require("./mockAPI.js");
var bodyParser = require("body-parser");
var cors = require("cors");
const fetch = require("node-fetch");

const app = express();
app.use(cors());
// to use json
app.use(bodyParser.json());
// to use url encoded values
app.use(
   bodyParser.urlencoded({
      extended: true,
   })
);

//app.use(express.static('src/client'))
// NOTE: to test for offline, commented out the following
app.use(express.static("dist"));

console.log(__dirname);

console.log(`Your API key is ${process.env.API_KEY}`);
//var textapi = new aylien({
//  application_id: "your-api-id",
//  application_key: process.env.API_KEY
//});
application_key = process.env.API_KEY;

// NOTE: to test for offline, commented out this block
app.get("/", function (req, res) {
   console.log("sending file");
   res.sendFile("dist/index.html");
   //res.sendFile(path.resolve('src/client/views/index.html'))
});


// designates what port the app will listen to for incoming requests
app.listen(8081, function () {
   console.log("server listening on port 8081!");
});

app.get("/test", function (req, res) {
   res.send(mockAPIResponse);
});

async function postData(data = {}) {
   var formBody = [];
   for (var property in data) {
      var encodedKey = encodeURIComponent(property);
      var encodedValue = encodeURIComponent(data[property]);
      formBody.push(encodedKey + "=" + encodedValue);
   }
   formBody = formBody.join("&");

   const response = await fetch("https://api.meaningcloud.com/sentiment-2.1", {
      method: "POST", // *GET, POST, PUT, DELETE, etc.
      mode: "cors", // no-cors, *cors, same-origin
      cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
      credentials: "same-origin", // include, *same-origin, omit
      headers: {
         //"Content-Type": "application/json",
         //'Content-Type': 'application/x-www-form-urlencoded',
         "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
      },
      redirect: "follow", // manual, *follow, error
      referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
      //body: JSON.stringify(data), // body data type must match "Content-Type" header
      body: formBody,
   });
   return response.json(); // parses JSON response into native JavaScript objects
}

app.get("/nlp", function (req, res) {
   console.log("receving url=" + req.query.url);
   postData({ url: req.query.url, key: application_key, lang: "en" }).then(
      (data) => {
         //console.log(data); // JSON data parsed by `data.json()` call
         res.send(data);
      }
   );
   // curl 'https://api.meaningcloud.com/sentiment-2.1' -F 'key=3cf5956a7a9fb6a364fcb4a46430cb6d' -F 'url=https://jasonwee.github.io/qj-srv250/' -F 'lang=en'
});
