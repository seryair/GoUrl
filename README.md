# GoUrl

Go Project location manager

build instructions:

1. copy the repository to your GOPATH and enter this directory, for example: 
cd ~/go/src
2. build first urlmanager package by running: 
go build github.com/seryair/golang/internal/app/urlmanager/
3. upon success run the following: 
go install github.com/seryair/golang/cmd/urlmanager
4. make sure bin folder exists and urlmanager.exe created.
run urlmanager.exe
5. window will pop-up, indicating the server is up (the port is 8080 by default).
in case of problems check the port is not taken, or change the port in file golang/cmd/urlmanager
6. now you can test the app by invoking POST request with the following request template in the body:
{
"domain" : "ynet.co.il",
"path" : "/abc.html"
}


Note:
there is a predefined DB i initialized (implemented as Map) which contains the following domain+path:
1. domain: "ynet.co.il", path: "/home/0,7340,L-8,00.html"
2. domain: "ynet.co.il", path: "/home/abc.html"
3. domain: "sport5", path: "/home/123/ac.html"
3. domain: "sport5", path: "/fds.html"
4. domain: "walla", path:"/home/12.html"

i've created extra api so you can use to add new locations.
POST http://127.0.0.1:8080/add
and supply in the body the request, for example:
{
"domain" : "ynet.co.il",
"path" : "/homegdkjd.html"
}
