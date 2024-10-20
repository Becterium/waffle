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
  containers:
    - name: jnlp
      image: 192.168.37.130:8009/library/jenkins/inbound-agent-kubectl:jdk17
      imagePullPolicy: IfNotPresent
      resources:
        limits:
          memory: "1Gi"
          cpu: "1000m"      
        requests:
          memory: "512Mi"
          cpu: "500m"
      volumeMounts:
        - name: kube-config
          mountPath: /root/.kube
        - name: docker-sock
          mountPath: /var/run/docker.sock
        - name: docker-bin
        mountPath: /usr/bin/docker
      env:
        - name: KUBECONFIG
          value: /root/.kube/config
  volumes:
    - name: kube-config
      hostPath:
        path: /etc/kubernetes/admin.conf
        type: File
    - name: docker-sock
      hostPath:
        path: /var/run/docker.sock
    - name: docker-bin
      hostPath:
        path: /usr/bin/docker

'''
        }
    }
    stages {
        stage('Kubernetes') {
            steps{
                script{
                  sh "kubectl get node"
                }
            }
        }
    }
    stages {
        stage('Docker') {
            steps{
                script{
                  sh "docker images"
                }
            }
        }
    }
}
