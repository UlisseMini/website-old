#!/bin/bash
cd /home/pi/website
echo "Updating..."
git pull
echo "Restarting webserver..."
sudo pkill webserver
sudo /home/pi/startup.sh
echo "Done! may take a minute to update."
