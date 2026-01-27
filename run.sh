#!/usr/bin/env bash

tmux split-window -v
tmux split-window -v
tmux select-layout even-vertical

#will be player 2
tmux send-keys -t 1 'sleep 1.5; clear; go run client/client.go localhost 9090 ' C-m
#will be player 1
tmux send-keys -t 2 'sleep 1; clear; go run client/client.go localhost 9090 ' C-m
go run server/server.go 9090
