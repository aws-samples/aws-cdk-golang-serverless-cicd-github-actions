# Streamlining Multi-Account Serverless Infrastructure with AWS CDK Golang and GitHub Reusable Workflows
This repository contains a solution for deploying serverless infrastructure across multiple AWS accounts and environments using AWS CDK with Golang and GitHub Actions. It's designed to help organizations streamline their deployment processes while adhering to best practices in cloud infrastructure management.

![Alt text](Diagram/Architecture-Diagram.png?raw=true)

## Key features:

- **Cross-account deployment**: Easily manage deployments across development, staging, and production environments in separate AWS accounts.
- **Infrastructure as Code (IaC)**: Utilize AWS CDK with Golang to define and version control your infrastructure and consistent deployments.
- **Automated CI/CD**: Implement GitHub Actions workflows to automate the entire CI/CD pipeline: Run unit tests, build docker images, push to ECR repository and deploy cdk infrastructure.
- **Best practices alignment**: Implement separation of concerns, promote collaboration, and ensure governance in multi-account setups.
- **Consistency and reliability**: Achieve repeatable deployments across various environments, reducing errors and improving overall system reliability.

This solution is ideal for teams looking to scale their serverless applications across multiple environments while maintaining a high standard of infrastructure management and deployment practices.
## Architecture Explained
#### GitHub Actions Workflow
- When a user initiates a GitHub Actions workflow, it triggers a Continuous Integration (CI) process. This process assumes an OpenID Connect (OIDC) role in the AWS Central Account.
#### AWS Central Account:
- The GitHub Actions workflow builds a Docker image
- The image is pushed to an Amazon Elastic Container Registry (ECR) repository.
- The image tag is stored in AWS Systems Manager Parameter Store.
- The GitHub Actions OIDC role has the necessary permissions for these operations.
- The ECR repository is configured with cross-account permissions.
- AWS Target Account CDK bootstrap roles are included in the permissions policy.
#### AWS Target Accounts:
- An AWS CDK Golang stack is utilized to create AWS Lambda functions as container images.
- The Lambda functions are configured with appropriate IAM permissions.
- The container images are fetched from the AWS Central Account's ECR repository.
- The Continuous Deployment (CD) process is automated using GitHub Actions.
- When running the CDK bootstrap command in the target accounts, trusts the GitHub Actions OIDC role.

## Prerequisites
- An Active AWS Account
 - AWS Command Line Interface (AWS CLI) version 2.9.11 or later, [installed](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html) and    [configured](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)
 - AWS CDK version 2.114.1 or later, [installed](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html#getting_started_install) and [bootstrapped](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html#getting_started_install)
 - Install Go 1.22 or later, [installed](https://go.dev/doc/install)
 - Install Docker 24.0.6 or later, [installed](https://docs.docker.com/engine/install/)

## 1. SetUp
 **1.1 Create ECR Repository in Central Account**
```bash
aws ecr create-repository --repository-name sample-repo
```
**1.2 Create GitHub OIDC Role in Central Account**
- Follow this [link](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services)
- Once Role is created, add necessary permissions to the role. For example, ECR permissions, SSM Permissions etc.

**1.3 Bootstrap CDK in Target Accounts**
```bash
cdk bootstrap aws://<Target_Account_ID>/<Target_Region> --trust <Central_Account_ID> --cloudformation-execution-policies arn:aws:iam::aws:policy/<Least_Privilege_Policy>
```
**1.4 Provide assume role permissions to assume Target Account CDK Bootstrap Roles in Central Account OIDC Role**
```bash
Target Account CDK Bootstrap Roles
arn:aws:iam::<Target_Account_ID>:role/cdk-deploy-role-<Target_Account_ID>-<Target_Region>
arn:aws:iam::<Target_Account_ID>:role/cdk-file-publishing-role-<Target_Account_ID>-<Target_Region>
arn:aws:iam::<Target_Account_ID>:role/cdk-image-publishing-role-<Target_Account_ID>-<Target_Region>
arn:aws:iam::<Target_Account_ID>:role/cdk-lookup-role-<Target_Account_ID>-<Target_Region>
```

# 2. Build the Docker Image
**2.1 Clone the project repository**
```bash
git clone https://github.com/aws-samples/aws-cdk-golang-serverless-cicd-github-actions.git
```
**2.2 Login into Amazon ECR Repository**
```bash
aws ecr get-login-password --region <AWS_REGION> | docker login --username AWS --password-stdin <AWS_ACCOUNT_ID>.dkr.ecr.<AWS_REGION>.amazonaws.com
```
**2.3 Build the Docker Image**
```bash
docker build --platform linux/arm64 -t sample-app .
```
**2.4 Tag the Docker Image and Push to ECR Repository**
```bash
# tag
docker tag sample-app:latest <AWS_ACCOUNT_ID>.dkr.ecr.<AWS_REGION>.amazonaws.com/<ECR_REPOSITORY>:<DOCKER_TAG>
# push
docker push <AWS_ACCOUNT_ID>.dkr.ecr.<AWS_REGION>.amazonaws.com/<ECR_REPOSITORY>:<DOCKER_TAG>
```

# 3. Deploy the AWS CDK App
**3.1 Synthesize AWS CDK Golang Stack**
```bash
ENV=<environment> IMAGETAG=<image_tag> ECR_ARN=<ecr_repo_arn> cdk synth
```
**3.2 Deploy AWS CDK Golang Stack**
```bash
ENV=<environment> IMAGETAG=<image_tag> ECR_ARN=<ecr_repo_arn> cdk deploy --require-approval never
```
# 4. Sample input.yml file
```bash
LambdaVariables:
    LambdaEnvVar : 
      ENV : "<ENV like Dev, Stage and Prod>"
      LOG_LEVEL : "<LOG_LEVEL>" 
      APP_NAME : "<APP Name>"
      APP_VERSION : "<APP Version>"
    Repo: "<ECR ARN>"
    Name: "<ECR Repository Name>"
    Account_Number: "<Target AWS Accout ID>"
    Region: "<Target AWS Acount Region>"
    Timeout: <Lambda Timeout>
```

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This library is licensed under the MIT-0 License. See the LICENSE file.
