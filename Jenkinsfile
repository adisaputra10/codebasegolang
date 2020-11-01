@Library('shared-library')_
def deployImage = new DeployImage()
env.nodeName = ""

pipeline {
    parameters {
        string(name: 'PRODUCTION_NAMESPACE',       description: 'Production Namespace',                 defaultValue: 'mypertamina20-prod')
        string(name: 'STAGING_NAMESPACE',          description: 'Staging Namespace',                    defaultValue: 'mypertamina20-stage')
        string(name: 'DEVELOPMENT_NAMESPACE',      description: 'Development Namespace',                defaultValue: 'mypertamina20-dev')

        string(name: 'DOCKER_IMAGE_NAME',          description: 'Docker Image Name',                    defaultValue: 'myptm-external-insurance-api')

        string(name: 'CHAT_ID',                    description: 'chat id of telegram group',            defaultValue: '-1001215679728')
    }
    agent none
    options {
        skipDefaultCheckout()  // Skip default checkout behavior
    }
    stages {
        stage ( "Kill Old Build" ){
            steps{
                script{
                    KillOldBuild()
        }   }   }
        stage('Checkout SCM') {
            agent { label "nodejs" }
            steps {
                sh "rm -Rf *"
                checkout scm
                script {
                    echo "get COMMIT_ID"
                    sh 'echo -n $(git rev-parse --short HEAD) > ./commit-id'
                    commitId = readFile('./commit-id')
                }
                stash(name: 'ws', includes:'**,./commit-id') // stash this current workspace
        }   }
        stage('Initialize') {
            parallel {
                stage("Agent: nodejs") {
                    agent { label "nodejs" }
                    steps {
                        cleanWs()
                           }   }
                stage("Agent: Docker") {
                    agent { label "Docker" }
                    steps {
                        cleanWs()
                        script{
                            if ( env.BRANCH_NAME == 'master' ){
                                envStage = "Production"
                            } else if ( env.BRANCH_NAME == 'release' ){
                                envStage = "Staging"
                            } else if ( env.BRANCH_NAME == 'develop'){
                                envStage = "Development"
        }   }   }   }   }   }
        stage('Test & Build') {
            parallel {
                stage('Unit Test') {
                    agent { label "Docker" }
                    steps {
                        unstash 'ws'
                        script {
                            echo "Do Unit Test Here"
                        
                           sh 'docker run --rm -v $PWD:/build -w /build telkomindonesia/alpine:go-1.13  /bin/sh -c "make test"'
                            echo "defining sonar-scanner"
                            def scannerHome = tool 'SonarScanner' ;
                            withSonarQubeEnv('SonarQube') {
                                sh "${scannerHome}/bin/sonar-scanner"
                }   }   }   }
                stage('Build') {
                    agent { label "Docker" }
                    steps {
                        unstash 'ws'
                        script{
                            env.nodeName = "${env.NODE_NAME}"
                            sh "docker build --rm --no-cache --pull -t ${params.DOCKER_IMAGE_NAME}:${BUILD_NUMBER}-${commitId} ."
        }   }   }   }   }
        stage ('Deployment'){
            steps{
                node (nodeName as String) { 
                    echo "Running on ${nodeName}"
                    script{
                        if (env.BRANCH_NAME == 'master'){
                            echo "Deploying to ${envStage} "
                            deployImage.to_myp("${commitId}")
                        } else if (env.BRANCH_NAME == 'release'){
                            echo "Deploying to ${envStage} "
                            deployImage.to_stagemyp("${commitId}")
                        } else if (env.BRANCH_NAME == 'develop'){
                            echo "Deploying to ${envStage} "
                            deployImage.to_vsan("${commitId}")
    }   }   }   }   }   }
    post {
        always{
            node("Docker"){
                TelegramNotif(currentBuild.currentResult) 
	}   }
	failure{
            node(nodeName as String){
                script{
                    sh "docker rmi -f ${params.DOCKER_IMAGE_NAME}:${BUILD_NUMBER}-${commitId}"
}   }   }   }   }
