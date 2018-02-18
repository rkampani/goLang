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
    1.docker run --rm someprefix/XXXXX
 
