const cpuUtilization = document.getElementById("progress-cpu-utilization");
const memoryUtilization = document.getElementById("progress-memory-utilization");
const cpuTemperature = document.getElementById("temperature-cpu");
const memoryTemperature = document.getElementById("temperature-memory");
const roomTemperature = document.getElementById("temperature-room");

setInterval(function () {
    var request = new XMLHttpRequest();

    request.timeout = 500
    request.open("GET", "/data.json", true);
    request.onload = function () {
        var data = JSON.parse(this.response);

        cpuUtilization.style.width = `${data.CPUUtilization}%`;
        memoryUtilization.style.width = `${data.MemoryUtilization}%`;
        cpuTemperature.innerHTML = `${data.CPUTemperature}`;
        memoryTemperature.innerHTML = `${data.MemoryTemperature}`;
        roomTemperature.innerHTML = `${data.RoomTemperature}`;
    };
    request.ontimeout = function () {
        alert("Request Timeout");
    };
    request.onerror = function () {
        alert("Request Error");
    };
    request.send();
} , 1000);
