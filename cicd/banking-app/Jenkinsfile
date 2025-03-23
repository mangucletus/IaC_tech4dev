pipeline {
    agent any
    
    environment {
        // DockerHub credentials and repository info
        DOCKER_HUB_CREDS = credentials('dockerhub-credentials')
        DOCKER_HUB_REPO = 'david930'  // Replace with your DockerHub username
        
        // Image names and tags
        API_IMAGE = "${DOCKER_HUB_REPO}/banking-api"
        PROCESSOR_IMAGE = "${DOCKER_HUB_REPO}/banking-processor"
        FRONTEND_IMAGE = "${DOCKER_HUB_REPO}/banking-frontend"
        
        // Version tag based on build number
        VERSION = "v1.0.${BUILD_NUMBER}"
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Lint and Format Check') {
            parallel {
                stage('Go Lint') {
                    agent {
                        docker {
                            image 'golang:1.24.1'
                        }
                    }
                    steps {
                        dir('cicd/banking-app/backend-api') {
                            sh 'go fmt ./...'
                            sh 'go vet ./...'
                        }
                    }
                }
                
                stage('Python Lint') {
                    agent {
                        docker {
                            image 'python:3.12'
                        }
                    }
                    steps {
                        dir('cicd/banking-app/transaction-service') {
                            echo 'Running flake8...'
                            sh 'python -m venv /opt/python-tools'
                            sh '/opt/python-tools/bin/pip install flake8'
                            sh '/opt/python-tools/bin/flake8 . --exit-zero --count --select=E9,F63,F7,F82 --show-source --statistics'
                        }
                    }
                }
                
                stage('JavaScript Lint') {
                    agent {
                        docker {
                            image 'node:22'
                        }
                    }
                    steps {
                        dir('cicd/banking-app/frontend') {
                            sh 'npm install eslint'
                            sh 'npx eslint . || true'  // Don't fail the build on lint errors
                        }
                    }
                }
            }
        }
        
        stage('Test') {
            parallel {
                stage('Go Tests') {
                    agent {
                        docker {
                            image 'golang:1.24.1'
                        }
                    }
                    steps {
                        dir('cicd/banking-app/backend-api') {
                            sh 'go test ./... -v'
                        }
                    }
                }
                
                stage('Python Tests') {
                    agent {
                        docker {
                            image 'python:3.12'
                        }
                    }
                    steps {
                        dir('cicd/banking-app/transaction-service') {
                            echo 'Running pytest...'
                            sh 'python -m venv /opt/python-tools'
                            sh '/opt/python-tools/bin/pip install pytest'
                            sh '/opt/python-tools/bin/pytest'
                        }
                    }
                }
            }
        }
        
        stage('Build Docker Images') {
            parallel {
                stage('Build API Image') {
                    steps {
                        dir('cicd/banking-app/backend-api') {
                            sh 'docker build -t ${API_IMAGE}:${VERSION} -t ${API_IMAGE}:latest .'
                        }
                    }
                }
                
                stage('Build Processor Image') {
                    steps {
                        dir('cicd/banking-app/transaction-service') {
                            sh 'docker build -t ${PROCESSOR_IMAGE}:${VERSION} -t ${PROCESSOR_IMAGE}:latest .'
                        }
                    }
                }
                
                stage('Build Frontend Image') {
                    steps {
                        dir('cicd/banking-app/frontend') {
                            sh 'docker build -t ${FRONTEND_IMAGE}:${VERSION} -t ${FRONTEND_IMAGE}:latest .'
                        }
                    }
                }
            }
        }
        
        stage('Push Images to DockerHub') {
            steps {
                // Login to DockerHub
                sh 'echo ${DOCKER_HUB_CREDS_PSW} | docker login -u ${DOCKER_HUB_CREDS_USR} --password-stdin'
                
                // Push all images
                sh '''
                docker push ${API_IMAGE}:${VERSION}
                docker push ${API_IMAGE}:latest
                docker push ${PROCESSOR_IMAGE}:${VERSION}
                docker push ${PROCESSOR_IMAGE}:latest
                docker push ${FRONTEND_IMAGE}:${VERSION}
                docker push ${FRONTEND_IMAGE}:latest
                '''
                echo 'Images pushed to DockerHub!'
            }
        }
    }
    
    post {
        always {
            // Clean up local Docker images
            sh '''
            docker rmi ${API_IMAGE}:${VERSION} ${API_IMAGE}:latest || true
            docker rmi ${PROCESSOR_IMAGE}:${VERSION} ${PROCESSOR_IMAGE}:latest || true
            docker rmi ${FRONTEND_IMAGE}:${VERSION} ${FRONTEND_IMAGE}:latest || true
            '''
            
            // Logout from DockerHub
            sh 'docker logout'
        }
        
        // success {
        //     echo 'Pipeline completed successfully!'
        //     // Send notifications on success
        //     emailext (
        //         subject: "Build Successful: ${currentBuild.fullDisplayName}",
        //         body: "The build was successful. Check the results at: ${env.BUILD_URL}",
        //         recipientProviders: [developers(), requestor()]
        //     )
        // }
        
        // failure {
        //     echo 'Pipeline failed!'
        //     // Send notifications on failure
        //     emailext (
        //         subject: "Build Failed: ${currentBuild.fullDisplayName}",
        //         body: "The build failed. Check the results at: ${env.BUILD_URL}",
        //         recipientProviders: [developers(), requestor()]
        //     )
        // }
    }
}
