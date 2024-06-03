# Go Go Gadget
This project provides a simple tool to run arbitrary commands and log the start time, STDOUT and stop time of commands. This information is sent to a log file you specify. This can be useful in penetration testing, bug bounty, red teaming or any other situation in which it is important to provide precise execution times of commands. This can be accomplished in several other ways, but was just a fun way to do something in Go

## Install
```sh 
    git clone https://github.com/skribblez2718/go_go_gadget.git
    cd go_go_gadget
    go build
```

## Usage
```sh
    go_go_gadget "<command>" <output_file>
```
