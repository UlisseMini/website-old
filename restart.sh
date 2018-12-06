#!/bin/bash
cd /home/pi/website
git pull
sudo pkill webserver
sudo /home/pi/startup.sh
