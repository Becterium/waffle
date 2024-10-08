pipeline {
    agent {
        kubernetes {
            cloud 'k8s-dev' // 与Jenkins配置中的Kubernetes Cloud名称一致
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
    resources:                   # 在这里定义资源限制
      limits:
        memory: "512Mi"          # 内存限制
        cpu: "500m"              # CPU限制
      requests:
        memory: "256Mi"          # 内存请求
        cpu: "250m"              # CPU请求
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
    }
    post {
        always {
            // 清理工作，例如删除临时文件或发送通知
            echo 'Cleaning up...'
        }
    }
}
