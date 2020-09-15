setInterval(function () {
	var request = new XMLHttpRequest();

	request.timeout = 500
	request.open("GET", "/data.json", true);
	request.onload = function () {
		var data = JSON.parse(this.response)

		document.writeln(`${data.Year}-${data.Month}-${data.Day}`)
		document.writeln(`${data.Hour}:${data.Minute}:${data.Second}`)
	};
	request.ontimeout = function () {
		document.writeln("Request Timeout");
	};
	request.onerror = function () {
		document.writeln("Request Error");
	};
	request.send();
} , 1000);
