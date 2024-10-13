// Uses Declarative syntax to run commands inside a container.
pipeline {
    agent {
        kubernetes {
            // Rather than inline YAML, in a multibranch Pipeline you could use: yamlFile 'jenkins-pod.yaml'
            // Or, to avoid YAML:
            // containerTemplate {
            //     name 'shell'
            //     image 'ubuntu'
            //     command 'sleep'
            //     args 'infinity'
            // }
            yaml '''
apiVersion: v1
kind: Pod
metadata:
  name: kaniko
spec:
  containers:
    - name: kaniko
      image: gcr.io/kaniko-project/executor:debug
      command:
        - /busybox/cat
      tty: true
      volumeMounts:
        - name: kaniko-secret
          mountPath: /kaniko/.docker/
  restartPolicy: Never
  volumes:
    - name: kaniko-secret
      secret:
        secretName: kaniko-secret
        items:
        - key: config.json
          path: config.json
'''
            // Can also wrap individual steps:
            // container('shell') {
            //     sh 'hostname'
            // }
            defaultContainer 'kaniko'
        }
    }

    environment {
        GIT_COMMIT = "${env.GIT_COMMIT}"
        // dockerfile = 'Dockerfile'
        destination = "coreharbor.azurewaf.top/devops/golang-exporter-demo"
    }

    parameters {
        string(name: 'IMAGE_TAG', description: 'The tag used in image pushing') 
        string(name: 'dockerfile', defaultValue: 'Dockerfile', description: 'Dockerfile Path in the project') 
    }

    stages {
        stage('Build Image') {
            steps {
                sh """
                /kaniko/executor -c `pwd`/ -f `pwd`/${dockerfile} -d ${destination}:${IMAGE_TAG}  --build-arg "BUILDDIR=`pwd`"
                """
            }
        }
    }
}
