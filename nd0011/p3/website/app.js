/* Global Variables */
const URL = 'https://api.openweathermap.org/data/2.5/weather';
//const URL = 'https://api.openweathermap.org/geo/1.0/zip';
const API_KEY = "2ea39989ad2278692864129c61a7d64c";



const getData = async ( url = '', data = {}, apiKey = '')=>{
   const aUrl = `${url}?q=${data.zip}&appid=${apiKey}&units=imperial`;
   //console.log(aUrl);
   const response = await fetch(aUrl, {
     method: 'GET', 
    });

    try {
      const newData = await response.json();
      console.log(newData);
      return newData
    }catch(error) {
    console.log("error", error);
    // appropriately handle the error
    }
}

const postData = async ( url, data = {})=>{

    const response = await fetch(url, {
    method: 'POST',
    credentials: 'same-origin',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify(data), // body data type must match "Content-Type" header
  });

    try {
      const newData = await response.json();
      return newData
    }catch(error) {
      console.log("error", error);
    }
}

const updateUI = async () => {
  const request = await fetch('/all');
  try{
    const allData = await request.json();
    const latest = Object.keys(allData)[Object.keys(allData).length - 1];
    document.getElementById('date').innerHTML = allData[latest].date;
    document.getElementById('temp').innerHTML = allData[latest].temperature;
    document.getElementById('content').innerHTML = allData[latest].user_response;

  }catch(error){
    console.log("error", error);
  }
}




// Create a new date instance dynamically with JS
let d = new Date();
let newDate = d.getMonth()+1+'.'+ d.getDate()+'.'+ d.getFullYear();


function generateHandler(zip, userResponse) {
   const data = {zip: zip};
   getData(URL, data, API_KEY).then(function(data) {
      //console.log("add data " + JSON.stringify(data));
      const d = {temperature: data.main.temp, date: newDate, user_response: userResponse}
      postData('http://localhost:3000/add', d);
   }).then(function(data) {
      updateUI()
   });
}
