pipeline {
    agent {
        kubernetes {
            cloud 'kubernetes' // 与 Jenkins 配置中的 Kubernetes Cloud 名称一致
            label 'k8s-agent'
            defaultContainer 'jnlp'
            yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: jenkins-agent
spec:
  containers:
  - name: jnlp
    image: 192.168.37.130:8009/library/jenkins/agent:jdk17
    args: ['\$(JENKINS_SECRET)', '\$(JENKINS_NAME)']
    imagePullPolicy: IfNotPresent
  imagePullSecrets:
  - name: harbor-credentials
"""
        }
    }
    stages {
        stage('Checkout') {
            steps {
                // 检出代码
                checkout scm
            }
        }
        stage('Build') {
            steps {
                // 在 golang 容器中执行构建命令
                container('golang') {
                    sh 'make build'
                }
            }
        }
        stage('Test') {
            steps {
                // 在 golang 容器中执行测试命令
                container('golang') {
                    sh 'make test'
                }
            }
        }
    }
    post {
        always {
            // 清理工作，例如删除临时文件或发送通知
            echo 'Cleaning up...'
        }
    }
}
