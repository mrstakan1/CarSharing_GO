pipeline {
    agent any

    environment {
        GO_VERSION = "1.22.1"  // Укажите версию Go, которая вам нужна
    }

    stages {
        stage('Checkout') {
            steps {
                dir("/e/\"3 курс 1 семестр\"/АКСП_КР/CarSharing")
            }
        }

        stage('Set Up Go') {
            steps {
                // Установка нужной версии Go
                sh 'wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz'
                sh 'sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz'
                sh 'export PATH=$PATH:/usr/local/go/bin'
            }
        }

        stage('Build') {
            steps {
                // Собираем бинарный файл проекта
                sh 'go build -o carsharing-app'
            }
        }

        stage('Test') {
            steps {
                // Запускаем тесты Go
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                // Собираем Docker-образ
                sh 'docker build -t carsharing-app .'
            }
        }

        stage('Deploy') {
            steps {
                // Разворачиваем Docker-контейнер
                sh 'docker-compose down && docker-compose up -d'
            }
        }
    }

//     post {
//         always {
//             mail to: 'mr.paxapanda@example.com',
//                  subject: "Jenkins Build: ${currentBuild.fullDisplayName}",
//                  body: "Build finished with status: ${currentBuild.currentResult}"
//         }
//     }
}
