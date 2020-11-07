#!/bin/bash
sudo yum -y update
sudo yum -y install docker

sudo service docker start

/*
sudo $(aws ecr get-login --region eu-central-1 --no-include-email)

sudo docker pull 472312519291.dkr.ecr.eu-central-1.amazonaws.com/ipo-ecr-backend-repo

sudo docker run -d --rm \
    -e MYSQL_ADDR=ipo-db.cjcdccc1jilp.eu-central-1.rds.amazonaws.com \
    -e MYSQL_DATABASE=theIPOguide \
    -e MYSQL_USER=g8FTNddx8AaESV83 \
    -e MYSQL_PASSWORD=zhBbv7SAxnPtWDqLNVVmqnJZfzvbWo \
    -e MYSQL_ROOT_PASSWORD=zhBbv7SAxnPtWDqLNVVmqnJZfzvbWo \
    -e CORS_ORIGIN=* \
    -e THROTTLE_LIMIT=40 \
    -e THROTTLE_BUCKET=4 \
    -e RECAPTCHA_SITE_URL=6LeuldkZAAAAAM7d7p74plbbGg_nT0B_4TC_CdmI \
    -e RECAPTCHA_SECRET=6LeuldkZAAAAAILfgsQi5SQy1HxU8hOEBAGs-1sq \
    -e EMAIL_HOST=smtp.gmail.com \
    -e EMAIL_FROM=jorge.br.ib@gmail.com \
    -e EMAIL_TO=jorge.br.ib@gmail.com \
    -e EMAIL_PASSWORD=Jbrisa1178R3a980 \
    -e EMAIL_PORT=587 \
    -p 80:80 \
    --name backend 472312519291.dkr.ecr.eu-central-1.amazonaws.com/ipo-ecr-backend-repo
*/