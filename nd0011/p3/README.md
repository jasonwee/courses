# Weather-Journal App Project

## Overview
This project requires you to create an asynchronous web app that uses Web API and user data to dynamically update the UI. 

## Instructions
This will require modifying the `server.js` file and the `website/app.js` file. You can see `index.html` for element references, and once you are finished with the project steps, you can use `style.css` to style your application to customized perfection.

start the server
```
$ source init-env-nodejs20.sh
$ npm install express body-parser cors 
$ node server.js
```

quick test
```
curl -XPOST 'localhost:3000/add' -d '{"temperature": 30.7, "date": "2023-05-05", "user_response": "enter 1"}'
```

surf to brower at localhost:3000 and play around

## Extras
If you are interested in testing your code as you go, you can use `tests.js` as a template for writing and running some basic tests for your code.
