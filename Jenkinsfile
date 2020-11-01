pipeline {
    parameters {
        string(name: 'PRODUCTION_NAMESPACE',       description: 'Production Namespace',                 defaultValue: 'namaproject-prod')

        string(name: 'DEVELOPMENT_NAMESPACE',      description: 'Development Namespace',                defaultValue: 'namaproject-dev')

        string(name: 'DOCKER_IMAGE_NAME',          description: 'Docker Image Name',                    defaultValue: 'image')
    }
    agent any
    stages {

        stage('Checkout SCM') {
            steps {
                
                script{
                    sh 'rm -Rf *'
                }
                checkout scm
                script {
                    echo "get COMMIT_ID"
                    sh 'echo -n $(git rev-parse --short HEAD) > ./commit-id'
                    commitId = readFile('./commit-id')
                }
                stash(name: 'ws', includes:'**,./commit-id') // stash this current 
            }
        }

        stage('Initialize') {
            steps {
                script{
                            if ( env.BRANCH_NAME == 'master' ){
                                envStage = "Production"
                            }else if ( env.BRANCH_NAME == 'development'){
                                envStage = "Development"
                    }   } 
                
            }
        }


        stage('SonarQube') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    echo "Multiline shell steps works too"
                    ls -lah
                '''
                script{
                    
                    
                    def scannerHome = tool 'sonarqube4.5' ;
                            withSonarQubeEnv('SonarQube') {
                                sh "${scannerHome}/bin/sonar-scanner"
                      }
                    
                }
                
                
            }
        }

        stage('Build Docker') {
            steps {
                
                script{
                    
                    sh "docker build --rm --no-cache --pull -t ${params.DOCKER_IMAGE_NAME}:${BUILD_NUMBER}-${commitId} ."
                    
                }
            }
        }

       stage('Deploy') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    echo "Multiline shell steps works too"
                    ls -lah
                '''
            }
        }


    }


post {
        always{
          sh 'ls '
        }
	
       }



}
