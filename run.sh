#!/usr/bin/env bash

tmux split-window -v
tmux split-window -v
tmux select-layout even-vertical

tmux send-keys -t 1 'sleep 2; clear; go run client/client.go localhost 9090 ' C-m
tmux send-keys -t 2 'sleep 2; clear; go run client/client.go localhost 9090 ' C-m
go run server/server.go 9090
