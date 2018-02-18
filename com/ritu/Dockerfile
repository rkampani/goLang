FROM iron/base
EXPOSE 8089


ADD ritu-linux-amd64 /
ENTRYPOINT ["./ritu-linux-amd64", "-configServerUrl=http://configserver:8888", "-profile=test", "-configBranch=P12"]
