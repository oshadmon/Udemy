sudo apt-get -y update 
sudo apt-get -y upgrade 
sudo apt-get -y update 

sudo apt-get -y install docker.io docker-compose

USER=`whoami`
sudo groupadd docker
sudo usermod -aG docker ${USER}
newgrp docker

sudo apt-get -y update
