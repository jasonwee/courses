# Project Description
This is a web tool that allows users to run Natural Language Processing (NLP) on articles or blogs found on other websites.


# Setup (one time only)
1. install packages
```
$ source init-env-nodejs12.sh
$ node -v
$ npm install
```
2. register an account at [MeaningCloud and get the license key](https://www.meaningcloud.com/developer/account/subscriptions). 
3. input the license key by create a file in the root directory of this project.
```
$ vim .env
API_KEY=[key from from step 2]
```
4. build
```
$ npm run build-prod
$ ls dist
```


# how-to run the app
1. start it
```
$ source init-env-nodejs12.sh
$ npm run start
```
2. go to your brower and load the url localhost:8081
3. if you want to test for offline, do step 1 and step 2. then do the following
```
commented out line in src/server.js

/*
app.use(express.static("dist"));
app.get("/", function (req, res) {
   console.log("sending file");
   res.sendFile("dist/index.html");
   //res.sendFile(path.resolve('src/client/views/index.html'))
});
*/

then restart npm run start, fresh your browser, the page should be still there and wget to the api server work as expected
```

# dependencies
see [package.json](package.json)
