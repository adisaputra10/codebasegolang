pipeline {
    parameters {
        string(name: 'PRODUCTION_NAMESPACE',       description: 'Production Namespace',                 defaultValue: 'test')

        string(name: 'DEVELOPMENT_NAMESPACE',      description: 'Development Namespace',                defaultValue: 'namaproject-dev')

        string(name: 'DOCKER_IMAGE_NAME',          description: 'Docker Image Name',                    defaultValue: 'nginx')
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
				    projectKubernetes= "${params.PRODUCTION_NAMESPACE}"
                                envStage = "production"
                            }else if ( env.BRANCH_NAME == 'development'){
				projectKubernetes= "${params.DEVELOPMENT_NAMESPACE}"
                                envStage = "development"
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
		    script{
                	   echo "Login Docker Registry"
		         withCredentials([string(credentialsId: 'kubernetes_token', variable: 'TOKEN')]) {
        			sh "docker login docker-registry-default.apps-kubernetes.bkn.go.id -u jenkins -p ${TOKEN}"
				 
				 imagefinal = "docker-registry-default.apps-kubernetes.bkn.go.id/${projectKubernetes}/${params.DOCKER_IMAGE_NAME}"
				 sh "docker tag ${params.DOCKER_IMAGE_NAME}:${BUILD_NUMBER}-${commitId} ${imagefinal}:latest"
				 sh  "docker push  ${imagefinal}:latest"
				 if ( env.BRANCH_NAME == 'master' ){
				      sh "docker tag ${params.DOCKER_IMAGE_NAME}:${BUILD_NUMBER}-${commitId} ${imagefinal}:prod-${BUILD_NUMBER}"
				      sh  "docker push ${imagefinal}:prod-${BUILD_NUMBER}"
                                 }
    			}
		    }
		    
            }
        }


    }


post {
        always{
          sh 'ls '
        }
	
       }



}
