FROM golang:alpine

ENV TC_RUNTIME_ENV=dev
RUN apk update && apk add --no-cache bash zsh curl wget git stow task ncurses make cmake openssh openrc neovim
RUN ssh-keygen -A && echo -e "PasswordAuthentication no" >> /etc/ssh/sshd_config
RUN mkdir -p /run/openrc && touch /run/openrc/softlevel
