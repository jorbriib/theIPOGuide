echo "--- deploying on AWS"

if [ -z $AWS_INSTANCE_NAME ]; then 
    echo "Must provide INSTANCE_NAME in environment" 1>&2
    exit 1
fi

ips=$(aws ec2 describe-instances --region ${AWS_REGION} --filters "Name=instance-state-name,Values=running" "Name=tag:Name,Values=${AWS_INSTANCE_NAME}" --query "Reservations[].Instances[0].NetworkInterfaces[].Association.PublicIp" --output text)
if [ -z $ips ]; then 
    echo "Unable to get instance IPs to deploy to" 1>&2
    exit 1
fi


for ip in $ips; do
    echo "--- deploying on ${ip} instance ---"

    REMOTE_COMMANDS=`cat <<EOF
        echo "Login";
        sudo aws ecr get-login-password --region ${AWS_REGION} | sudo docker login --username AWS --password-stdin ${AWS_ECR_URL};
            
        echo "Stoping containers";
        sudo docker stop backend || true;

        echo "Removing images";
        sudo docker rmi \\\$(sudo docker images --filter reference=${AWS_ECR_URL}/${AWS_ECR_REPO} --format "{{.ID}}")  || true;
        sudo docker rmi \\\$(sudo docker images --filter reference=${AWS_ECR_URL}/${AWS_ECR_MIGRATIONS_REPO} --format "{{.ID}}")  || true;
        
        echo "Pulling images";
        sudo docker pull ${AWS_ECR_URL}/${AWS_ECR_REPO};
        sudo docker pull ${AWS_ECR_URL}/${AWS_ECR_MIGRATIONS_REPO};

        echo "Executing migrations";
        sudo docker run --rm     
            -e ENV=${ENV}
            -e MYSQL_ADDR=${DB_HOST} \
            -e MYSQL_DATABASE=${DB_NAME} \
            -e MYSQL_USER=${DB_USER} \
            -e MYSQL_PASSWORD=${DB_PASSWORD} \
            -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} \
            --name migrations ${AWS_ECR_URL}/${AWS_ECR_MIGRATIONS_REPO};

        echo "Running backend container";
        sudo docker run -d --rm \
            -e ENV=${ENV}
            -e MYSQL_ADDR=${DB_HOST} \
            -e MYSQL_DATABASE=${DB_NAME} \
            -e MYSQL_USER=${DB_USER} \
            -e MYSQL_PASSWORD=${DB_PASSWORD} \
            -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} \
            -e CORS_ORIGIN=* \
            -e THROTTLE_LIMIT=40 \
            -e THROTTLE_BUCKET=4 \
            -e RECAPTCHA_SITE_URL=${RECAPTCHA_SITE_KEY} \
            -e RECAPTCHA_SECRET=${RECAPTCHA_SECRET} \
            -e EMAIL_HOST=${EMAIL_HOST} \
            -e EMAIL_FROM=${EMAIL_FROM} \
            -e EMAIL_TO=${EMAIL_TO} \
            -e EMAIL_PASSWORD=${EMAIL_PASSWORD} \
            -e EMAIL_PORT=${EMAIL_PORT} \
            -p 80:80 \
            --name backend ${AWS_ECR_URL}/${AWS_ECR_REPO};

        exit;
EOF    
`

    ssh -T -o "StrictHostKeyChecking=no" ec2-user@$ip ${REMOTE_COMMANDS}
done