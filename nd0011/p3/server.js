// Setup empty JS object to act as endpoint for all routes
projectData = {};

// Require Express to run server and routes
const express = require('express');

// Start up an instance of app
const app = express();



/* Middleware*/
//Here we are configuring express to use body-parser as middle-ware.
const bodyParser = require('body-parser')
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

// Cors for cross origin allowance
const cors = require('cors');
app.use(cors());


// Initialize the main project folder
app.use(express.static('website'));

const port = 3000;
// Setup Server
const server = app.listen(port, listening);

function listening() {
    console.log(`running on localhost: ${port}`);
}

// routes
app.get('/all', function (req, res) {
  res.send(projectData);
})

app.post('/add', function(request, response){
  let data = request.body;
  response.json({"receive": "ok"});
  projectData[new Date().getTime()] = data;
})

