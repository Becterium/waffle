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
      image: 192.168.37.130:8009/library/jenkins/inbound-agent:jdk17
      imagePullPolicy: IfNotPresent
      resources:
        limits:
          memory: "1Gi"
          cpu: "1000m"      
        requests:
          memory: "512Mi"
          cpu: "500m"
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
