# Plasmawatch Login Server
Plasmawatch is an Overwatch 1 Private Server recreation project. This is the login server for the project. It directly emulates the Battle.net login experience for the client.

Join the discord: https://discord.gg/T5RncE2RkN

Written in Go.

## Building and Running
1. Clone the repository either by using `git clone https://github.com/plasmawatch/login-server` or clicking the green `Code` button, then `Download ZIP`.
2. (Optional if cloned via the command) Unzip the ZIP file to a directory you'll remember.
3. Go into that directory and run `go build`, this will install the necessary packages for the login server to work and compile the executable to be ran.
4. Double click on the `bnet-emu.fish.pfx` file.
5. Set store location to `Current User`.
6. Click next after making sure the file name is correct. (It should be as it's automatically grabbed)
7. Type in `generated-pass-abcd` as the password in the password box. Click next.
8. Click `Place all certificates in the following store` and click the `Browse` button.
9. Click on `Trusted Root Certification Authorities` and click `Ok`. Click next.
10. Click finish and press Ok or Yes to the rest of the following windows.
11. Open the `bnet-mock.exe` file and the rest should work. 

## Running (Release Tab)
1. Extract all the files from the 7z file into an empty directory.
2. Double click on the `bnet-emu.fish.pfx` file.
3. Set store location to `Current User`.
4. Click next after making sure the file name is correct. (It should be as it's automatically grabbed)
5. Type in `generated-pass-abcd` as the password in the password box. Click next.
6. Click `Place all certificates in the following store` and click the `Browse` button.
7. Click on `Trusted Root Certification Authorities` and click `Ok`. Click next.
8. Click finish and press Ok or Yes to the rest of the following windows.
9. Open the `bnet-mock.exe` file and the rest should work. 

## Extra Information
This is for extra information for the login server, please use our Launcher repo to learn how to setup your Overwatch 1 client for use.

### Ports Used:
- TLS3 Server: 1119
- Auth Webserver: 6969

### Public Hosting of Login Server
While hosting the Login server is allowed. It is recommended that you use a reverse proxy using Nginix or Apache to make sure the Webserver is using HTTPS.
