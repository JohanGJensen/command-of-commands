# Command of commands (coc)

A CLI tool written in GO that lets you list and select script commands directly from the terminal rather than inspecting the package.json file.

## Just to be clear
This functionality already largely exists in all of the major node package managers:

```sh
yarn run
npm run
pnpm run
```

Where command-of-commands differ, is that you select and run the command immediately and, eventually (WIP) list all commands from all package.json files recursively in a project.

## Build & install
--more to come--

## Getting started
--more to come--

## The concept
Have you ever had to navigate a lot of different repositories and run a lot of different package.json script commands to install and spin of different applications? Rather than concatenate and print (cat) or enter your code editor of choice, this small application lists out available options for you and lets you select a command to run.

This is a very simple cobra based CLI application written in GOLANG.
