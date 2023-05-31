/* groovylint-disable DuplicateStringLiteral, GStringExpressionWithinString, LineLength, NestedBlockDepth */
import groovy.json.JsonSlurperClassic

@NonCPS
def jsonParse(def json) {
    new groovy.json.JsonSlurperClassic().parseText(json)
}
pipeline {
    agent { label 'ecs-agent' }
    environment {
        APPNAME = 'go-restapi'
        IMAGE_NAME = 'go-restapi'
        AWS_REGION = 'us-east-1'
        AWS_ACCOUNT = credentials('AWS_ACC_ID')
        CLUSTER_NAME = 'go-restapi-demo-cluster'
        SERVICE_NAME = 'go-restapi-demo-service'
        IMAGE_PORT = '9090'
        COMMIT = getShortCommitId()
        TASK_DEFINITION_NAME = "${APPNAME}"
        REPO_NAME = "${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com/${IMAGE_NAME}"
    }
    stages {
        stage('Build') {
            steps {
                script {
                    git branch: 'develop', url: 'https://github.com/kriz23/go-restapi'
                    withCredentials([string(credentialsId: 'GO_RESTAPI_DB_URL', variable: 'GO_RESTAPI_DB_URL')]) {
                        sh 'echo "DB_URL=\"${GO_RESTAPI_DB_URL}\"" > .env;'
                        sh 'sleep 2;'
                    }
                    sh "docker build -t ${REPO_NAME} . --no-cache;"
                    sh "docker tag ${REPO_NAME}:latest ${REPO_NAME}:${COMMIT};"
                }
            }
        }
        stage('Push image to ECR Repository') {
            steps {
                script {
                    sh "aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com;"
                    sh "docker push ${REPO_NAME}:latest;"
                    sh "docker push ${REPO_NAME}:${COMMIT};"
                }
            }
        }
        stage('Cleaning') {
            steps {
                script {
                    sh 'docker system prune -fa;'
                }
            }
        }
        stage('Update Service') {
            steps {
                script {
                    TASK_DEFINITION = sh(returnStdout: true, script: "aws ecs update-service --cluster ${CLUSTER_NAME} --service ${SERVICE_NAME} --force-new-deployment;")
                }
            }
        }
    }
}
def getShortCommitId() {
    def gitCommit = env.GIT_COMMIT
    def shortGitCommit = "${gitCommit[0..6]}"
    return shortGitCommit
}
