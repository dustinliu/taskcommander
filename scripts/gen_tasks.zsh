#!/bin/zsh
gentitle() {
    cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 20 | head -n 1
}

gentask() {
    task add inbox0 $(gentitle) category:Inbox project:test1
    task add inbox1 $(gentitle) category:Inbox project:test2
    task add inbox2 $(gentitle) category:Inbox project:3test

    task add next0 $(gentitle) category:Next +test1
    task add next1 $(gentitle) category:Next +test2
    task add next2 $(gentitle) category:Next +test1 +test2
    task add next3 $(gentitle) category:Next
    task add next4 $(gentitle) category:Next
}


case $1 in
clear)
    rm -rf ~/.task/*
    ;;
*)
    gentask
    ;;
esac