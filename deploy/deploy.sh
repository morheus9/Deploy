#! /bin/sh

# sudo chmod 666 /var/run/docker.sock  -- if don't login
docker login -u morheus -p morheus1917
docker pull morheus/go_app
docker container rm -f go_app_latest || true
docker run -d -p 8080:8080 --name go_app_latest morheus/go_app
