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
      imagePullPolicy: Always
      resources:
        limits:
          memory: "1Gi"
          cpu: "1000m"      
        requests:
          memory: "512Mi"
          cpu: "500m"
      volumeMounts:
        - name: kube-config
          mountPath: /home/jenkins/.kube
        - name: docker-sock
          mountPath: /var/run/docker.sock
        - name: docker-bin
          mountPath: /usr/bin/docker
      env:
        - name: KUBECONFIG
          value: /home/jenkins/.kube/config
  volumes:
    - name: kube-config
#k8s集群的kubeconfig,需要你手动复制到各个节点,如有更好的解决方法,希望指出
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
        stage('Docker') {
            steps{
                script{
                  sh "docker images"
                }
            }
        }
    }
}
