![alt tag](files/gophernotes-logo.png)

[![Build Status](https://travis-ci.org/throff-jupyter/throff-jupyter.svg?branch=master)](https://travis-ci.org/throff-jupyter/throff-jupyter)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/throff-jupyter/throff-jupyter/blob/master/LICENSE)

# throff-jupyter - Use throff in Jupyter notebooks and nteract

`throff-jupyter` is a Throff kernel for [Jupyter](http://jupyter.org/) notebooks and [nteract](https://nteract.io/).  It lets you use throff interactively in a browser-based notebook or desktop app.  Use `throff-jupyter` to create and share documents that contain live Throff code, equations, visualizations and explanatory text.  These notebooks, with the live Throff code, can then be shared with others via email, Dropbox, GitHub and the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/). 

**Acknowledgements** - This project is forked from [Gophernotes](https://github.com/gopherdata/gophernotes).

- [Examples](#examples)
- Install gophernotes:
  - [Prerequisites](#prerequisites)
  - [Linux](#linux)
  - [Mac](#mac)
  - [Windows](#windows)
  - [Docker](#docker)
- [Getting Started](#getting-started)
- [Limitations](#limitations)
- [Troubleshooting](#troubleshooting)

### Example Notebook (dowload and run it locally, follow the links to view in Github, or use the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/)):
- [Introduction](introduction.ipynb)


## Installation

### Prerequisites

- [Go 1.11+](https://golang.org/doc/install) - including GOPATH/bin added to your PATH (i.e., you can run Go binaries that you `go install`).
- [Jupyter Notebook](http://jupyter.readthedocs.io/en/latest/install.html) or [nteract](https://nteract.io/desktop)
- [ZeroMQ 4.X.X](http://zeromq.org/intro:get-the-software) - for convenience, pre-built Windows binaries (v4.2.1) are included in the zmq-win directory.
- [pkg-config](https://en.wikipedia.org/wiki/Pkg-config)

### Linux

```sh
$ go get -u github.com/throff-jupyter/throff-jupyter
$ mkdir -p ~/.local/share/jupyter/kernels/throff-jupyter
$ cp $GOPATH/src/github.com/throff-jupyter/throff-jupyter/kernel/* ~/.local/share/jupyter/kernels/throff-jupyter
```

To confirm that the `throff-jupyter` binary is installed and in your PATH, you should see the following when running `throff-jupyter` directly:

```sh
$ throff-jupyter
2017/09/20 10:33:12 Need a command line argument specifying the connection file.
```

**Note** - if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:

```sh
$ jupyter --data-dir
```

### Mac

```sh
$ go get -u github.com/throff-jupyter/throff-jupyter
$ mkdir -p ~/Library/Jupyter/kernels/throff-jupyter
$ cp $GOPATH/src/github.com/throff-jupyter/throff-jupyter/kernel/* ~/Library/Jupyter/kernels/throff-jupyter
```

To confirm that the `throff-jupyter` binary is installed and in your PATH, you should see the following when running `throff-jupyter` directly:

```sh
$ throff-jupyter
2017/09/20 10:33:12 Need a command line argument specifying the connection file.
```

**Note** - if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:

```sh
$ jupyter --data-dir
```

### Windows

Make sure you have the MinGW toolchain:

- [MinGW-w64](https://sourceforge.net/projects/mingw-w64/), for 32 and 64 bit Windows
- [MinGW Distro](https://nuwen.net/mingw.html), for 64 bit Windows only

Then:

1. build and install gophernotes (using the pre-built binaries and `zmq-win\build.bat`):

    ```
    REM Download w/o building.
    go get -d -u github.com/throff-jupyter/throff-jupyter
    cd %GOPATH%\src\github.com\throff-jupyter\throff-jupyter\zmq-win

    REM Build x64 version.
    build.bat amd64
    move throff-jupyter.exe %GOPATH%\bin
    copy lib-amd64\libzmq.dll %GOPATH%\bin

    REM Build x86 version.
    build.bat 386
    move throff-jupyter.exe %GOPATH%\bin
    copy lib-386\libzmq.dll %GOPATH%\bin
    ```

3. Copy the kernel config:

    ```
    mkdir %APPDATA%\jupyter\kernels\throff-jupyter
    xcopy %GOPATH%\src\github.com\throff-jupyter\throff-jupyter\kernel %APPDATA%\jupyter\kernels\throff-jupyter /s
    ```

    Note, if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:

    ```
    jupyter --data-dir
    ```

4. Update `%APPDATA%\jupyter\kernels\gophernotes\kernel.json` with the FULL PATH to your gophernotes.exe (in %GOPATH%\bin), unless it's already on the PATH.  For example:

    ```
    {
        "argv": [
          "C:\\gopath\\bin\\throff-jupyter.exe",
          "{connection_file}"
          ],
        "display_name": "Go",
        "language": "go",
        "name": "go"
    }
    ```

### Docker

You can try out or run Jupyter + throff-jupyter without installing anything using Docker. To run a Go notebook that only needs things from the standard library, run:

```
$ docker run -it -p 8888:8888 throff-jupyter/throff-jupyter
```

Or to run a Go notebook with access to common Go data science packages (gonum, gota, golearn, etc.), run:

```
$ docker run -it -p 8888:8888 throff-jupyter/throff-jupyter:latest-ds
```

In either case, running this command should output a link that you can follow to access Jupyter in a browser. Also, to save notebooks to and/or load notebooks from a location outside of the Docker image, you should utilize a volume mount.  For example:

```
$ docker run -it -p 8888:8888 -v /path/to/local/notebooks:/path/to/notebooks/in/docker throff-jupyter/throff-jupyter
```

## Getting Started

### Jupyter

- If you completed one of the local installs above (i.e., not the Docker install), start the jupyter notebook server:

  ```
  jupyter notebook
  ```

- Select `throff` from the `New` drop down menu.

- Have fun!

### nteract

- Launch nteract.

- From the nteract menu select Language -> throff.

- Have fun!


## Troubleshooting

### throff-jupyter not found

Depending on your environment, you may need to manually change the path to the `throff-jupyter` executable in `kernel/kernel.json` before copying it to `~/.local/share/jupyter/kernels/throff-jupyter`.  You can put the **full path** to the `throff-jupyter` executable here, and you shouldn't have any further issues.

### "Kernel error" in a running notebook

```
Traceback (most recent call last):
  File "/usr/local/lib/python2.7/site-packages/notebook/base/handlers.py", line 458, in wrapper
    result = yield gen.maybe_future(method(self, *args, **kwargs))
  File "/usr/local/lib/python2.7/site-packages/tornado/gen.py", line 1008, in run
    value = future.result()
  ...
  File "/usr/local/Cellar/python/2.7.11/Frameworks/Python.framework/Versions/2.7/lib/python2.7/subprocess.py", line 1335, in _execute_child
    raise child_exception
OSError: [Errno 2] No such file or directory
```

Stop jupyter, if it's already running.

Add a symlink to `/go/bin/throff-jupyter` from your path to the gophernotes executable. If you followed the instructions above, this will be:

```
sudo ln -s $HOME/go/bin/throff-jupyter /go/bin/throff-jupyter
```

Restart jupyter, and you should now be up and running.
