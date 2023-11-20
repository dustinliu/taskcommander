#!/bin/zsh
cp test/taskrc /root/.taskrc
go install github.com/golang/mock/mockgen@v1.6.0
rc-status && rc-service sshd start
