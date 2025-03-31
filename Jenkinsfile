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
          mountPath: /home/jenkins/.kube/admin.conf
        - name: docker-sock
          mountPath: /var/run/docker.sock
        - name: docker-bin
          mountPath: /usr/bin/docker
        - name: docker-config
          mountPath: /root/.docker
          readOnly: true
      env:
        - name: KUBECONFIG
          value: /home/jenkins/.kube/admin.conf
      imagePullSecrets:
        - name: harbor-key
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
    - name: docker-config
      hostPath:
        path: /root/.docker
        type: Directory
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
      stage('Code') {
          steps{
              script{
                sh "cat /home/jenkins/agent/workspace/waffle/deploy/kubernetes/minio.yaml"
              }
          }
      }
      stage('BuildUser'){
        steps{
          script{
            sh "docker login 192.168.37.130:8009 -u admin -p Harbor12345"
            sh "cd /home/jenkins/agent/workspace/waffle"
            sh "docker build --build-arg APP_RELATIVE_PATH=user/service -t 192.168.37.130:8009/library/waffle/user-kratos:latest ."
            sh "docker push 192.168.37.130:8009/library/waffle/user-kratos:latest"
          }
        }
      }
      stage('BuildWaffle'){
        steps{
          script{
            sh "cd /home/jenkins/agent/workspace/waffle"
            sh "docker build --build-arg APP_RELATIVE_PATH=waffle/interface -t 192.168.37.130:8009/library/waffle/waffle-kratos:latest ."
            sh "docker push 192.168.37.130:8009/library/waffle/waffle-kratos:latest"
          }
        }
      }
      stage('BuildMedia'){
        steps{
          script{
            sh "cd /home/jenkins/agent/workspace/waffle"
            sh "docker build --build-arg APP_RELATIVE_PATH=media/service -t 192.168.37.130:8009/library/waffle/media-kratos:latest ."
            sh "docker push 192.168.37.130:8009/library/waffle/media-kratos:latest"
          }
        }
      }
      stage("run"){
        steps{
          script{
            sh "kubectl create configmap user-waffle-config --from-file=/home/jenkins/agent/workspace/waffle/app/user/service/configs --dry-run=client -o yaml | kubectl apply -f -"
            sh "kubectl create configmap gateway-waffle-config --from-file=/home/jenkins/agent/workspace/waffle/app/waffle/interface/configs --dry-run=client -o yaml | kubectl apply -f -"
            sh "kubectl create configmap media-waffle-config --from-file=/home/jenkins/agent/workspace/waffle/app/media/service/configs --dry-run=client -o yaml | kubectl apply -f -"
            sh "kubectl apply -f /home/jenkins/agent/workspace/waffle/deploy/kubernetes/user.yaml"
            sh "kubectl apply -f /home/jenkins/agent/workspace/waffle/deploy/kubernetes/waffle.yaml"
            sh "kubectl apply -f /home/jenkins/agent/workspace/waffle/deploy/kubernetes/media.yaml"
          }
        }
      }
    }
}
