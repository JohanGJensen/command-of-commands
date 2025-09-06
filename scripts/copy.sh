#!/bin/bash

mkdir -p $HOME/bin
mv ../coc  $HOME/bin

export PATH="$HOME/bin:$PATH"
source ~/.zshrc # chane to approiate shell config
