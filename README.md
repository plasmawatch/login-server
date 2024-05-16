# Plasmawatch Login Server
Plasmawatch is an Overwatch 1 Private Server recreation project. This is the login server for the project. It directly emulates the Battle.net login experience for the client.

Join the discord: https://discord.gg/T5RncE2RkN

Written in TypeScript using Node 18.18.

## Building and Running
1. Clone the repository either by using `git clone https://github.com/plasmawatch/login-server` or clicking the green `Code` button, then `Download ZIP`.
2. (Optional if cloned via the command) Unzip the ZIP file to a directory you'll remember.
3. Go into that directory and run `npm i`, this will install the necessary packages for the login server to work.
4. Finally, run `npm test` to start the server (I know it says test, it's just the easiest.)

We do not have any binaries as this is a Node applet.

## Extra Information
This is for extra information for the login server, please use our Launcher repo to learn how to setup your Overwatch 1 client for use.

### Ports Used:
- TLS3 Server: 8657 (Can be modified in .env)
- Auth Webserver: 8658 (Can be modified in .env, not recommended.)
