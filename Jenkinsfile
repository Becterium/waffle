pipeline {
    agent {
        kubernetes {
            cloud 'k8s-dev' // 与Jenkins配置中的Kubernetes Cloud名称一致
            defaultContainer 'jnlp'
            yaml '''
---
kind: Pod
apiVersion: v1
metadata:
  labels:
    k8s-app: jenkins-agent
  name: jnlp
  namespace: waffle
spec:
  imagePullSecrets:
    - name: myregistrykey
  containers:
    - name: jnlp
      image: 192.168.37.130:8009/library/jenkins/agent:jdk17
      imagePullPolicy: IfNotPresent
      resources:
        limits:
          memory: "512Mi"
          cpu: "500m"      
        requests:
          memory: "256Mi"
          cpu: "250m"
'''
        }
    }
    stages {
        stage('Checkout') {
            steps{
                script{
                  sh "sleep 30"
                }
            }
        }
    }
}
