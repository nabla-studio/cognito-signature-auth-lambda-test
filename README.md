# Cognito Signature Auth Lambda 🔏λ

## Repository Name
`cognito-signature-auth-lambda`

## Description
Cognito Signature Auth Lambda is a project that to test a custom signature-based authentication flow for AWS Cognito using Lambda functions and the Serverless Application Model (SAM). This service provides a sample framework for integrating public key cryptography and digital signatures into Cognito's authentication process. 🔐🚀

## Key Features 🌟
- Custom Cognito User Pool Lambda triggers for signature-based authentication
- Public key and signature verification during user registration and authentication
- Custom challenge creation and verification using cryptographic signatures
- Flexible authentication flow management centered around digital signatures
- Extensible post-authentication actions

## Project Structure 📁
```
.
├── Makefile
├── README.md
├── internal
│   └── helpers
│       └── signature.go
├── lambdas
│   ├── create-auth-challenge
│   ├── define-auth-challenge
│   ├── post-authentication
│   ├── pre-sign-up
│   └── verify-auth-challenge-response
├── samconfig.toml
└── template.yaml
```

## Main Components 🧩

### Cognito Lambda Triggers λ
1. **Pre Sign-Up** (`lambdas/pre-sign-up/`)
   - Validates user's public key and initial signature during Cognito sign-up

2. **Define Auth Challenge** (`lambdas/define-auth-challenge/`)
   - Manages the flow of signature-based authentication challenges

3. **Create Auth Challenge** (`lambdas/create-auth-challenge/`)
   - Generates cryptographic challenges for signature-based authentication

4. **Verify Auth Challenge Response** (`lambdas/verify-auth-challenge-response/`)
   - Verifies digital signatures provided in response to authentication challenges

### Helpers 🛠️
- **Signature Verification** (`internal/helpers/signature.go`)
   - Provides core functionality for cryptographic operations and signature verification

## Configuration ⚙️
AWS SAM is used for configuration and deployment:
- `template.yaml`: Defines AWS resources, including Cognito User Pool and Lambda functions
- `samconfig.toml`: Configuration for SAM CLI

## Prerequisites 📋
- Go 1.x 🐹
- AWS CLI 🧰
- AWS SAM CLI 🛠️
- AWS account with necessary permissions 🔑

## Installation and Deployment 🚀
1. Clone the repository: `git clone https://github.com/your-username/cognito-signature-auth-lambda.git`
2. Navigate to the project directory: `cd cognito-signature-auth-lambda`
3. Run `sam build` to compile the Lambda functions
4. Run `sam deploy` to deploy the stack to AWS

## Usage 🖥️
Once deployed, this service creates a Cognito User Pool with custom Lambda triggers that implement a signature-based authentication flow. Users can register and authenticate using their public keys and digital signatures.

## Development 👨‍💻👩‍💻
To modify the project:
1. Update Lambda functions in the `lambdas/` directory to adjust signature verification or challenge processes
2. Modify `template.yaml` to fine-tune Cognito User Pool or Lambda function configurations

## Notes 📝
- Ensure proper configuration of AWS environment variables and permissions
- This project provides a foundation for implementing robust, signature-based authentication in Cognito
- Thorough cryptographic review and security testing is recommended before any production use