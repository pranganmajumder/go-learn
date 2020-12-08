#                                       GO

### Install Go
To install Go, run the following command in terminal
```
$ go_version=1.15.5
$ cd ~/Downloads
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz
$ sudo chown -R $(id -u):$(id -g) /usr/local/go
$ rm go${go_version}.linux-amd64.tar.gz
```

Then add go to the $PATH variable
```
$ mkdir $HOME/go
$ nano ~/.bashrc
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
$ source ~/.bashrc
$ go version
```

You can use GoLand IDE or Visual Studio code to run go executable code.
Install anyone you prefer most.

### Start to Learn Go
* [An Introduction to Programming in GO](https://www.golang-book.com/books/intro)

### Other Helpful Site
* [golangbot](https://golangbot.com/learn-golang-series/)
* [gobyexample](https://gobyexample.com/)