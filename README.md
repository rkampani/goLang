# goLang
goLang

Adding some simple Docker configuration
1. Create DockerFile on the root of the app
2. go build
3. Run below commands to build docker image
    1. export GOOS=linux
    2. go build -o XXXXX-linux-amd64 (XXXXX=app)
    3. export GOOS=darwin

4. docker build -t someprefix/XXXXX
5. Run on docker now -  However - note that this container is not running on your host OS localhost anymore. It now lives in it’s own networking context and we can’t actually call it directly from our host operating system
    1. docker run --rm someprefix/XXXXX

6. Create overlay networl
docker network create --driver overlay my_network
7. Deploy the app now
docker service create --name=XXXXX --replicas=1 --network=my_network -p=8089:8089 someprefix/XXXXX 042lqtptrkqe0emv4q95mbgf9

Here’s a quick rundown of the arguments:

–name: Assigns a logical name to our service. This is also the name other services will use when addressing our service within the cluster. So if you had a another service that would like to call the accountservice, that service would just do a GET to http://name:8089/api/bolt/getUsers
–replicas: The number of instances of our service we want. If we’d have a multi-node Docker Swarm cluster the swarm engine would automatically distribute the instances across the nodes.
–network: Here we tell our service to attach itself to the overlay network we just created.
-p: Maps [internal port]:[external port]. Here we used 6767:6767 but if we’d created it using 6767:80 then we would access the service from port 80 when calling externally. Note that this is the mechanism that makes our service reachable from outside the cluster. Normally, you shouldn’t expose your services directly to the outside world. Instead, you’d be using an EDGE-server (e.g. a reverse proxy) that would have routing rules and security setup so external consumers wouldn’t be able to reach your services except in the way you’ve intended them to.
someprefix/accountservice: This how we specify which image we want the container to run. In our case this is the tag we specified when we created the container. Note! If we’d be running a multi-node cluster we would have had to push our image to a Docker repository such as the public (and free) Docker Hub service. One can also set up private docker repositories or use a paid service if you want your images to stay private.


8. docker service ls (Check services succesfully started)
9. 
