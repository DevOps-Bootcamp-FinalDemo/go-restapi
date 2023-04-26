pipeline {
    agent { label 'worker1' }
    stages {
        stage('Build') {
            steps {
                git branch: 'main', url: 'https://github.com/kriz23/go-restapi'
                withCredentials([string(credentialsId: 'GO_RESTAPI_DB_URL', variable: 'GO_RESTAPI_DB_URL')]) {
                    sh 'echo "DB_URL=\"${GO_RESTAPI_DB_URL}\"" > .env'
                    sh 'sleep 2'
                }
                script {
                    def envFile = readFile('.env')
                    envFile = envFile.replaceAll(/(DB_URL=)(.*)/, '$1"$2"\n')
                    writeFile(file: '.env', text: envFile)
                }
                sh 'sleep 2'
                sh 'docker build -t krizz23/go-restapi:jenkins .'
            }
        }
        stage('Push image to DockerHub') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'dockerHubUser', variable: 'dockerHubUser'), string(credentialsId: 'dockerHubPass', variable: 'dockerHubPass')]) {
                        sh 'docker login -u ${dockerHubUser} -p ${dockerHubPass}'
                    }
                    sh 'docker push krizz23/go-restapi:jenkins'
                }
            }
        }
        stage('Cleaning') {
            steps {
                sh 'docker rmi -f $(docker images -q krizz23/go-restapi)'
            }
        }
    }
}
