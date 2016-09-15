Build the docker image and create 2 services, then loop querying each and look for not "OK" lines.

    docker build -t dangra/swarm-net-bug .
    docker service create --name s1 -p 10001:8080 dangra/swarm-net-bug app --id 1
    docker service create --name s2 -p 10002:8080 dangra/swarm-net-bug app --id 2
    while true; do curl http://localhost:10001/1; done
    while true; do curl http://localhost:10002/2; done
