## ShellyMicroWeb
ShellyMicroWeb is a micro webserver for testing and logging Shelly Actions. It will open a http port on port 8080. Once started on a computer point a Shelly action with a trigger to the computer.

### Example: 

webserver is running on a computer with ip 192.168.178.86

ACTION BUTTON LONG PRESSED URL is set to:

```
http://192.168.178.86:8080/button_long_pressed
```

once the action is triggered by pressing the button on the shelly it should create a webserver.log file with the following content:

```
2020/11/07 20:55:47 192.168.178.212 - button_long_pressed
```

### Build instructions

If not already done add the go gorilla mux router via console:

```
go get github.com/gorilla/mux
```

a simple `go build ShellyMicroWeb.go` will create a binary file for your platform.
