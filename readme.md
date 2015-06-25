

# Install Ubuntu 14.04:
sudo apt-get update
sudo ufw allow ssh/tcp
sudo ufw enable

# Install docker: 
https://docs.docker.com/installation/ubuntulinux/

wget -qO- https://get.docker.com/ | sh
mkdir apps

# copy app to docker contianer
rm -r sparkapi/sparkapi 
rm -r sparkapi/tmp/
scp -r sparkapi/ fredrik@172.20.10.7:/home/fredrik/apps

# Build docker image
https://registry.hub.docker.com/_/golang/

docker build -t spark-api .

# Run docker image
docker run -d -p 80:5050 --name api spark-api 

# Postman
GET: 172.20.10.7:80/api/graph/1
