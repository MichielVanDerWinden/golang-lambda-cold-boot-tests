# golang-lambda-cold-boot-tests

## Introduction
This repository is a simple test case regarding the cold boot times of Lambda using Golang.
It focuses on the time it takes to start an entirely new Lambda invocation using the Golang runtime.

Part of your cold boot times is based on the amount of initialization in your code. This is often handled by dependency injection; a method of injecting required variables into constructors of each of the required classes to start your application. Golang does this by default, using `import` and `init` to create an entire tree of dependencies that each need to be satisfied before the entire application starts.
A simple setup for this would be a `main.go` file that imports several controllers, each in turn importing several dependencies (e.g. utilities, commonly used libraries, services, models etc.)

## The possible issue with Golang on AWS Lambda
AWS Lambda can run for at most 15 minutes before an invocation or instance will be terminated. Also, an AWS Lambda can only run one invocation at a given time. As a Lambda instance has been created, it can be re-used multiple times when no invocations are used and the instance is not yet terminated. This means that your Lambda can have two states: 
1. Cold boot - No Lambda instance readily available for the given invocation, thus one has to be created "from scratch".
2. Warm boot - A previously used Lambda instance is available for the given invocation, thus it can be re-used.

During cold boot, the entire initialization of your code (and possibly container) has to be run before the invocation can be handled properly. This means that each time a cold boot occurs, your code will have to traverse the entire dependency graph and satisfy all dependencies that are necessary for proper operation.

As AWS Lambda is created specifically for short running, simple tasks, you'll want to keep your code path "as clean as possible" with as little initializations as possible to prevent latency on each invocation. By default, Golang will initialize all code in your import clauses and thus it is prone to taking a (relatively) long time before your application is ready to be invoked.
This repository will give insights in the amount of time it takes, and how we can tackle the above issue.

### My approach
As Golang is automatically running your `init()` statements on import, I've removed any direct invocations of connectivity/setup to external libraries in my optimized codebase. Instead, I've moved those to a separate function that will be called every time when the code path is initially hit by an external invocation, and will cache the resulting connection/setup response to prevent having to do that again every time the path gets hit. The changes I've made are apparent in `src/optimized/pkg/aws`; each of these utility classes now contains a `getService()` method that'll create, cache and return the object it uses for calls to (external) services.

## Test setup
This repository has a test setup that has been divided into three parts:
1. `./terraform` - contains all AWS setup files; you'll have to run this for yourself to recreate my testing methodology
2. `./src/default` - contains a "naive" method of using `import` and `init` as described above
3. `./src/optimized` - contains another approach to initialization and dependency injection "when needed"

### Test prerequisites
I've not tested this on any other machine than my own, which is running Linux (Manjaro). If you're running Windows or MacOS, I can't say for certain all paths, executables and variables will work exactly the same.

You'll at least need the following tools installed:
1. [go](https://go.dev/) >= v1.20
2. [Terraform](https://www.terraform.io/) >= v1.6

### Important notes on the test setup
During the `terraform apply` phase, a `bootstrap` file is created in both the `src/default` and `src/optimized` directories. These binaries are used by Lambda to run your code.
Building this code is handled by the module itself, and will only trigger once during `terraform apply` - afterwards you'll have to trigger a re-build as shown in the next chapter.

Do note that both code sets contain a `time.Sleep(time.Second)` in each of their code paths. This is necessary to "emulate" a heavy initialization step of your code, as importing an AWS library, starting a connection and setting up the library itself didn't give enough variance to show differences in the approach to dependency injection.
This means that as long as you're only doing some AWS interactions in your code and won't ever be using other, external libraries or heavier initialization of code paths, you'll find nothing of use in this repository ;)

### Steps to recreate
Follow these steps to recreate the test setup and start running the tests:
1. Deploy the Terraform code using `terraform init`, `terraform plan` and `terraform apply` in the `./terraform` directory
2. Both of the Lambda services should now be deployed to your AWS account; test these by using `curl` (or any other REST client for that matter) for the following API Gateway endpoints:
 - `{api-gateway-endpoint}/dev/default/s3`
 - `{api-gateway-endpoint}/dev/default/parameterstore`
 - `{api-gateway-endpoint}/dev/default/dynamodb`
 - `{api-gateway-endpoint}/dev/optimized/s3`
 - `{api-gateway-endpoint}/dev/optimized/parameterstore`
 - `{api-gateway-endpoint}/dev/optimized/dynamodb`

> **Note: If you'd like to change the code after already having ran `terraform apply` and thus creating the Lambda's, please remove the Lambda's and the `bootstrap` files in the `src/default` and `src/optimized` directories. This will trigger a re-build of the `main.go` and dependencies into the `bootstrap` binary that is used by Lambda for running the code.**

### Test results (aggregated)
#### Cold boot `src/default`:
 - Min: 4.15s
 - Avg: 4.21s
 - Max: 4.25s

#### Cold boot `src/optimized`:
 - Min: 2.08s
 - Avg: 2.12s
 - Max: 2.22s

As we can see from the aggregated data, using the optimized flow for dependency injection provides quicker cold boot results. Do note that using this way of working means that you're splitting up the load of the `default` approach over multiple code paths, meaning that invoking each code path separately will still result in nearly the same latency, but split over each invocation.

## Conclusion
The given approach can greatly improve the cold boot latency of your Golang Lambda invocations. The changes I made are minor, but they could affect your entire workflow when done properly. Still, they're prone to "it depends" - this way of working won't make a difference if you're abiding by the "one Lambda for one task" rule, or if there are no real separate code paths in your Lambda function.

All in all, I like how easy it is to start with Golang, but how we can also optimize really quickly by making simple changes to our code. The language works wonders in AWS Lambda, both when running in a docker container and when deployed as a binary. It's my new favourite programming langguage for private projects to come and will be continuously updating this repo with new insights in the future.