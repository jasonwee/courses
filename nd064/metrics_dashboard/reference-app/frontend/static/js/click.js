$(document).ready(function () {

    // all custom jQuery will go here
    $("#firstbutton").click(function () {
        $.ajax({
            url: "http://localhost:8081/api", success: function (result) {
                $("#firstbutton").toggleClass("btn-primary:focus");
                }
        });
    });
    $("#secondbutton").click(function () {
        $.ajax({
            url: "http://localhost:8082", success: function (result) {
                $("#secondbutton").toggleClass("btn-primary:focus");
            }
        });
    });    
});
