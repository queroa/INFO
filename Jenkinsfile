pipeline {
    agent any
    tools {
        go 'golang'
        dockerTool 'docker'
        
    }
    stages {
        stage ('Install dependencies'){
            steps {
                echo 'Instalando dependencias'
                sh 'go version'
                sh 'go get -u  github.com/go-sql-driver/mysql'
                sh 'go get -u github.com/gorilla/mux'
                
            }
            
        }
        
            
        
        stage ('Build'){
            steps {
                echo 'Compilando y Build'
                sh 'go build -o getter get.go'
                
            }
            
        }
        stage('Dockerize') {
            steps{
                sh 'docker build . -t info'
                
            }
            
        }
        
    }
    
}
