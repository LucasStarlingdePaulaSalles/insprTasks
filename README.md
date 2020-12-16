# InsprTasks

This project was a technical challange as well as a proof of concept for Golang as a CLI and API developing tool.

The result is an application for mannaging tasks that can be used from any directory in the user's computer. This was acheived by the combination of a dockerized REST CRUD-like api, that manages a mySQL database, and a simple CLI frontend that once compiled and installed can allways be quickly by the use of single command line action.

------
## Dependencies && Installation

### Backend
The api uses a couple of custom go packages, a mysql instance and phpmyadmin intance, however these dependencies are handled by the Docker engine on building it's containers.

Installing the backend is as simple as cloning the repository and running `docker-compose up --build` on the `/backend` directory.

### Frontend 
The CLI application is arguably even simpler. It doesn't require any custom packages only a default `1.15` go installation.

All it really takes is compiling the `main.go` file located on the `myTasks` directory and executing it. However using go's `install` built in functionality is highly reccomended. By doing this a `myTasks` binary file will be generate on your $GOPATH/bin directory, and so if you add this directory to your bash paths you can use the CLI application from anywhere on computer.

Adding your $GOPATH/bin to system's paths is also very simple. Head to /home/your-user/ and open the `.bashrc` file on your editor of choice. Then simply add `export PATH=$PATH:$GOPATH/bin` to the end of it. If you don't know your GOPATH just run `go env GOPATH` and use it on the previous command.

