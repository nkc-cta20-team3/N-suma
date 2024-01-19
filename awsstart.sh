#!/bin/sh
cd ~/N-suma
git pull
sudo docker compose -f docker-compose-preview.yaml up -d --build