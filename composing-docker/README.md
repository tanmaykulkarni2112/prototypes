> # Docker Cleanup commands
docker container rm -f $(docker container ls -aq)

docker image rm -f $(docker image ls -q)

backend 3001

frontend 3000 communicate with backend

docker compose up to build

parsing the yml is slower

Thereby we use the yml files for config 
and the json for data exchange between systems


docker network so containers can talk to each other 