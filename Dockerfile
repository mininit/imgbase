FROM ubuntu:latest 
COPY imgbase /usr/local/bin/imgbase 
ENTRYPOINT [ "/usr/local/bin/imgbase" ]
