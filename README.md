## webui

Binary webui [golang]

## build

~~~bash
➜ make clean
➜ make releases
~~~

## download

[the latest Binary](https://github.com/hyxf/webui/releases/latest)

## help

~~~bash
➜ ./webui -h
NAME:
   webui - aria2 web ui

USAGE:
   webui [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value              server host (default: "127.0.0.1")
   --port value              server port (default: "8089")
   --signal value, -s value  Usage: install | remove | start | stop | status
   --help, -h                show help
   --version, -v             print the version
~~~

## run

~~~bash
➜ ./webui
2018/09/02 11:47:45 http://127.0.0.1:8089/
~~~

## daemon run

~~~bash
➜ sudo ./webui -s start
2018/09/02 11:53:47 Starting aria2 webui service:				[  OK  ]
~~~

## daemon stop

~~~bash
➜ sudo ./webui -s stop
2018/09/02 11:55:20 Service has already been stopped
~~~

---------------------------------------------------------

## start aria2 rpc

~~~bash
➜ aria2c --enable-rpc --rpc-listen-all -D
~~~

## support

[aria2](https://github.com/aria2/aria2)

[webui-aria2](https://github.com/ziahamza/webui-aria2)



