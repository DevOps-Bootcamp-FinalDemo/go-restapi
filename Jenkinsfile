/* groovylint-disable DuplicateStringLiteral, GStringExpressionWithinString, LineLength, NestedBlockDepth */
import groovy.json.JsonSlurperClassic

@NonCPS
def jsonParse(def json) {
    new groovy.json.JsonSlurperClassic().parseText(json)
}
pipeline {
    agent { label 'worker3' }
    environment {
        APPNAME = 'go-restapi'
        IMAGE_NAME = 'go-restapi'
        AWS_REGION = 'us-east-1'
        AWS_ACCOUNT = credentials('AWS_ACC_ID')
        CLUSTER_NAME = 'Demov2-Cluster'
        SERVICE_NAME = 'demov2'
        IMAGE_PORT = '9090'
        COMMIT = getShortCommitId()
        TASK_DEFINITION_NAME = "${APPNAME}"
        REPO_NAME = "${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com/${IMAGE_NAME}"
    }
    stages {
        stage('Build') {
            steps {
                script {
                    git branch: 'ecs-deployment', url: 'https://github.com/kriz23/go-restapi'
                    withCredentials([string(credentialsId: 'GO_RESTAPI_DB_URL', variable: 'GO_RESTAPI_DB_URL')]) {
                        sh 'echo "DB_URL=\"${GO_RESTAPI_DB_URL}\"" > .env;'
                        sh 'sleep 2;'
                    }
                    sh "docker build -t ${REPO_NAME} . --no-cache"
                    sh "docker tag ${REPO_NAME}:latest ${REPO_NAME}:${COMMIT}"
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
                    sh "docker rmi ${REPO_NAME}:${COMMIT};"
                    sh "docker rmi ${REPO_NAME}:latest;"
                }
            }
        }
        stage('Add task definition') {
            steps {
                script {
                    sh " sed -i -e 's;%APPNAME%;${APPNAME};g' -e 's;%ECRIMAGEN%;${REPO_NAME}:latest;g' deploy/ec2-task-definition.json"
                    sh " sed -i -e 's;%IMAGEPORT%;${IMAGE_PORT};g' deploy/ec2-task-definition.json"
                    TASK_DEFINITION = sh(returnStdout: true, script:"\
                    aws ecs register-task-definition --region ${AWS_REGION} --cli-input-json file://deploy/ec2-task-definition.json\
                    ")
                }
            }
        }
        stage('Create Service') {
            steps {
                script {
                    TASK_DEFINITION = sh(returnStdout: true, script: "aws ecs create-service --launch-type EC2 --cluster ${CLUSTER_NAME} --desired-count 1 --service-name ${SERVICE_NAME} --region ${AWS_REGION} --task-definition ${APPNAME} --deployment-configuration 'minimumHealthyPercent=0,maximumPercent=100' || true")
                }
            }
        }
        stage('Update Service') {
            steps {
                script {
                    TASK_DEFINITION = sh(returnStdout: true, script: "aws ecs update-service --cluster ${CLUSTER_NAME} --desired-count 1 --service ${SERVICE_NAME} --region ${AWS_REGION} --task-definition ${APPNAME} --force-new-deployment")
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
