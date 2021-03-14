sudo apt-get -y update 
sudo apt-get -y upgrade 
sudo apt-get -y update 

sudo apt-get -y install docker.io

USER=`whoami`
sudo groupadd docker
sudo usermod -aG docker ${USER}
newgrp docker





