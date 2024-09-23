pipeline {
    agent any
    stages {
        stage('Test SSH Connection') {
            steps {
                sh 'ssh -T git@github.com'
            }
        }
    }
}
