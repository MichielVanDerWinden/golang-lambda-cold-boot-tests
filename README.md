# golang-lambda-cold-boot-tests

## Introduction
This repository is a simple test case regarding the cold boot times of Lambda using Golang.
It focuses on the time it takes to start an entirely new Lambda invocation using the Golang runtime in Docker.

Part of your cold boot times is based on the amount of initialization in your code. This is often handled by dependency injection; a method of injecting required variables into constructors of each of the required classes to start your application. Golang does this by default, using `import` and `init` to create an entire tree of dependencies that each need to be satisfied before the entire application starts.
A simple setup for this would be a `main.go` file that imports several controllers, each in turn importing several dependencies (e.g. utilities, commonly used libraries, services, models etc.)

## The possible issue with Golang on AWS Lambda
AWS Lambda can run for at most 15 minutes before an invocation or instance will be terminated. Also, an AWS Lambda can only run one invocation at a given time. As a Lambda instance has been created, it can be re-used multiple times when no invocations are used and the instance is not yet terminated. This means that your Lambda can have two states: 
1. Cold boot - No Lambda instance readily available for the given invocation, thus one has to be created "from scratch".
2. Warm boot - A previously used Lambda instance is available for the given invocation, thus it can be re-used.

During cold boot, the entire initialization of your code (and possibly container) has to be run before the invocation can be handled properly. This means that each time a cold boot occurs, your code will have to traverse the entire dependency graph and satisfy all dependencies that are necessary for proper operation.

As AWS Lambda is created specifically for short running, simple tasks, you'll want to keep your code path "as clean as possible" with as little initializations as possible to prevent latency on each invocation. By default, Golang will initialize all code in your import clauses and thus it is prone to taking a (relatively) long time before your application is ready to be invoked.
This repository will give insights in the amount of time it takes, and how we can tackle the above issue.

## Test setup
This repository has a test setup that has been divided into three parts:
1. `./terraform` - contains all AWS setup files; you'll have to run this for yourself to recreate my testing methodology
2. `./src/default` - contains a "naive" method of using `import` and `init` as described above
3. `./src/optimized` - contains another approach to initialization and dependency injection "when needed"

## Important notes on the test setup
To simplify deployment and maintenance of the code as well as handling dependencies, the deployment of the code to AWS Lambda is done in a Docker container image. Although the overhead of using Docker might be larger than the overhead of Golang's dependency management, it is used in both the `default` and `optimized` codebases, thus having no impact on total runtime other than marginal variance between Docker startup times within AWS Lambda.
As AWS Lambda keeps a "local" copy of the Docker image when it is invoked at least once in the last X amount of time, each first cold boot execution time result will be discarded. This prevents us from having execution time variance coming from the interaction between AWS Lambda and ECR.

## Steps to recreate
Follow these steps to recreate the test setup and start running the tests:
1. Change the S3 bucket name that is used for the Terraform state file in `./terraform/state.tf`
2. Enter the name of the above S3 bucket into `./src/default/pkg/controllers/s3/controller.go`'s `init()` function
3. Enter the name of the above S3 bucket into `./src/optimized/pkg/controllers/s3/controller.go`'s `init()` function
4. Deploy the Terraform code using `terraform init`, `terraform plan` and `terraform apply` in the `./terraform` directory
5. Both of the Lambda services should now be deployed to your AWS account; test these by using `curl` for the following API Gateway endpoints:
 - `{api-gateway-endpoint}/dev/default/s3`
 - `{api-gateway-endpoint}/dev/default/parameterstore`
 - `{api-gateway-endpoint}/dev/default/dynamodb`
 - `{api-gateway-endpoint}/dev/optimized/s3`
 - `{api-gateway-endpoint}/dev/optimized/parameterstore`
 - `{api-gateway-endpoint}/dev/optimized/dynamodb`
6. The execution time results may vary depending on cold boot, warm boot and of course the availability of the docker image for AWS Lambda itself, but when taking note of the averages and standard deviations, you'll see something comparable to the figures below.

## Test results
