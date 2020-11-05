# Create an instance

  - services / EC2
  - launch instance
  - choose your instance
  - add storage / 30GB free
  - add tags / webserver
  - security / ssh / http
  - launch
  - create new key pair / download

# Deploy your binary

  - mv [src] [dst] / sudo chmod 400 your.pem

1. Build hello world
  - GOOS=linux GOARCH=amd64 go build -o helloapp

2. Copy your binary to the sever
  - scp -i ~/.ssh/your-key.pem ./helloapp ec2-user@your-dns:/home/ec2-user

3. SSH into your server
  - cd ~/.ssh
  - ssh -i "your-key.pem" ec2-user@your-dns

4. Run your code
  - sudo chmod 700 helloapp
  - sudo ./helloapp


# Running app as deamon

To run our application after the terminal session has ended, we must do one of the following:

## Possible options
1. screen
2. init.d
3. upstart
4. system.d

## System.d
1. Create a configuration file
  - cd /etc/systemd/system/
  - sudo nano hello.service

```
[Unit]
Description=Go Server

[Service]
ExecStart=/home/ec2-user/helloapp
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

1. Add the service to systemd.
  - sudo systemctl enable hello.service
1. Activate the service.
  - sudo systemctl start hello.service
1. Check if systemd started it.
  - sudo systemctl status hello.service
1. Stop systemd if so desired.
  - sudo systemctl stop hello.service


# Troubleshooting

A possible issue could be that you're cross-compiling for the wrong architecture: AWS might have assigned you a different machine than the one used in this example. To solve this problem, we will install Go on the AWS machine and then run "go env" to see GOOS & GOARCH for that machine.

1. download Go
  - wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
1. unpack go
  - tar -xzf go1.7.4.linux-amd64.tar.gz
1. remove the tar file
  - rm -rf go1.7.4.linux-amd64.tar.gz
1. make your go workspace
  - mkdir goworkspace
  - cd gowoworkspace
  - mkdir bin pkg src
  - cd ../
1. add environment variables
  - nano .bashrc
```
export GOROOT=/home/ubuntu/go
export GOPATH=/home/ubuntu/goworkspace
export PATH=$PATH:/home/ubuntu/goworkspace/bin
export PATH=$PATH:/home/ubuntu/go/bin
```
1. refresh environment variables
  - source ~/.bashrc
1. confirm installation
  - go version
1. get machine GOOS & GOARCH info
  - go env

# Troubleshooting

Sometimes students miss setting port openings in security. If you are having issues, check to make sure these settings are correct - and please note, you IP address for SSH will either be 0.0.0.0/0 or something different than mine.
![](security.png)

