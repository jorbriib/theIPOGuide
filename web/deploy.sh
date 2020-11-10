cd web

echo "Installing AWS-CLI"
sudo apt-get update
sudo apt-get install -y -qq python-pip libpython-dev
pip install --user awscli
curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip"
unzip awscli-bundle.zip
sudo ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws

echo "Installing dependencies"
npm i

echo "Building web"
NODE_ENV=${ENV} TARGET_ENV=${ENV} APP_URL=${APP_URL} API_URL=${API_URL} RECAPTCHA_SITE_KEY=${RECAPTCHA_SITE_KEY} npm run build 

echo "Pushing public to S3"
aws s3 sync public s3://${AWS_BUCKET_NAME} --delete --exclude "*.js.map"