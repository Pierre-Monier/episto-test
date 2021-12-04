# episto-test



This project is based on the [gorilla/websocket chat example](https://github.com/gorilla/websocket/tree/master/examples/chat). 

It add the folling feature : 

* Create room on deman
* Add a username to the user
* Add automatic message when user enter/leaver the room

First download the project with `git clone https://github.com/Pierre-Monier/episto-test.git` then go to the repository with `cd episto-test`

You can run the server with `go run *.go` or you can run it with docker (if you don't have Go set up on your computer) 

To build the project with docker, at the root of the project do `docker build . -t episto-test:latest` then you can run it with `docker run -p 8080:8080 episto-test:latest`

To test the server, go to `http://localhost:8080/?room=roomname&username=toto`, you can now send message to the room, to create/join a new room, just change the `room` parameter in the URL. To change your user name change the `username` parameter 
