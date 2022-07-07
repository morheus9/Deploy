#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/www/devops
git pull git@gitlab.com:gEdisonLeun/devops.git
./webserver &