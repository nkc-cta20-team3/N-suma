# Student Support System
This repository was prepared for students of Nagoya Kogakuin College to develop a student support system, "NKC Smart Document" or "N-Sma" for short, as their graduation project.
(Translated by DeepL)

[日本語版はこちら](README.md)

# How to use

### Tools you need
* Git
* Docker
* docker-compose

### How do I activate N-Sma?

1. Install the above tools. 

2. Execute the commands in 2-1 to start up in the development environment, and the commands in 2-2 to start up in the production environment using various terminals (Git Bash is recommended).

2-1. 
```terminal
git clone https://github.com/nkc-cta20-team3/N-suma.git
cd N-suma
docker compose --profile dev_front --profile dev_back up -d --build
```

2-2. 
```terminal
git clone https://github.com/nkc-cta20-team3/N-suma.git
cd N-suma
docker compose -f docker-compose-preview.yaml up -d
```

3. When stopping N sma, execute command 3-1 if it was started with command 2-1, or command 2 if it was started with command 3-2.

3-1. 
``` terminal
docker compose stop   
```

3-2. 
``` terminal
docker compose -f docker-compose-preview.yaml stop
```

4. To completely destroy the environment, execute the following command. <br>
For environments built with the command 2-2, do not forget to specify the docker-compose-preview.yaml file with the -f option.
```
docker compose down --rmi all --volumes --remove-orphans
```

