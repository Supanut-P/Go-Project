pipeline {
    agent any

    stages {
        stage('Clone') {
            steps {
                echo 'Cloning repository...'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -v ./...'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
