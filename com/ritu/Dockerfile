FROM iron/base
EXPOSE 8089
EXPOSE 8089


ADD ritu-linux-amd64 /
ENTRYPOINT ["./ritu-linux-amd64", "-configServerUrl=http://configserver:8888", "-profile=test", "-configBranch=P12"]

#token=$(docker -H tcp://$92.168.99.121:2375 swarm join-token worker -q)
#192.168.99.121

#192.168.99.100
#docker swarm join --token $(docker -H tcp://192.168.99.100:2375 swarm join-token worker -q) --listen-addr "192.168.99.100:2377" --advertise-addr "192.168.99.100:2377" "$192.168.99.100:2377"