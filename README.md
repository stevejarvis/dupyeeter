
## Deploy
There's a messy dependency circle to bootstrap because Lambda needs the ECR image to 
exist and the image needs the ECR to exist. Both Lambda and ECR repo defined in 
Terraform. But updating the image is an app deployment that depends on the infra from 
Terraform. Just kinda do it all twice and it'll be happy. Only a problem on initial deploy.

0. `make infra`
1. `make deploy`
