# Swarm
Swarm is a clustering solution built inside docker 

## Starting Swarm
1. Check if swarm is active: `docker info | grep -i swarm`
2. start swarm: `docker swarm init`
    * create PKI 
    * enable swarm API 
    * Starts Raft database 
```commandline
# output
Swarm initialized: current node (wa0f5i1aay8oyxcacxg0znr11) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-0hjxqbkfbdalsfv84vs6wwa1m1u3p6qy9poeckcv1ghqhco8g2-bgt2izofzirqv6y61imvwk64b 10.0.2.15:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```
## Commands
* `docker service` replaces `docker run` when running in Swarm
