# gourl - location manager

<b>Build instructions:</b>

1. Copy the repository to your GOPATH and enter this directory, for example:  <br />
cd ~/go/src
2. Build urlmanager package by running: <br />
$ go build github.com/seryair/golang/internal/app/urlmanager/
3. Run the following: <br />
$ go install github.com/seryair/golang/cmd/urlmanager
4. Make sure bin folder exists and urlmanager.exe created.<br />
run urlmanager.exe
5. Window will pop-up, indicating the server is up (the port is 8080 by default).<br />
in case of problems check the port is not taken, or change the port in file golang/cmd/urlmanager
6. now you can test the app by invoking POST request to the following url: <br />
POST http://127.0.0.1:8080/<br />
with the following request template in the body:<br />
{<br />
"domain" : "ynet.co.il",<br />
"path" : "/abc.html"<br />
}<br />


<b>Special Note: </b>
there is a predefined DB i initialized (implemented as Map) which contains the following domain+path:<br />
1. domain: "ynet.co.il", path: "/home/0,7340,L-8,00.html"
2. domain: "ynet.co.il", path: "/home/abc.html"
3. domain: "sport5", path: "/home/123/ac.html"
3. domain: "sport5", path: "/fds.html"
4. domain: "walla", path:"/home/12.html"

i've created extra api so you can use to add new locations.<br />
POST http://127.0.0.1:8080/add<br />
and supply in the body the request, for example:<br />
{<br />
"domain" : "ynet.co.il",<br />
"path" : "/homegdkjd.html"<br />
}
