pipeline {
    agent {
        docker {
            image 'golang:1.23'
            args '-v /go/pkg/mod:/go/pkg/mod'
        }
    }

    stages {
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
