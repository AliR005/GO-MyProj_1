# TELEGRAMM BOT

## Инструкция
        1) Создайте в папке 'cmd' файл '.env'
        2) Поместите в него подомным образом след. переменные:
                BOT_TOKEN=your_token
                BOT_DB_DRIVER=postgres

                BOT_POSTGRES_HOST=localhost
                BOT_POSTGRES_PORT=5432
                BOT_POSTGRES_USERNAME=your_username
                BOT_POSTGRES_PASSWORD=your_password
                BOT_POSTGRES_DBNAME=tgbotmessages
                BOT_POSTGRES_SSL_MODE=disable

                RABBITMQ_URI=amqp://guest:guest@localhost:5672/
                RABBITMQ_QUEUE=telegram_queue
                RABBITMQ_DEFAULT_USER=user
                RABBITMQ_DEFAULT_PASS=password
                
        3) Скачайьте след. верисию образа RabbitMQ:
            docker pull rabbitmq:3-management

        4) Запустите контейнер с RabbitMQ:
            docker run -d --name cont_rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

        5) Находясь в папке 'cmd', откройте терминал и пропишите в нём 'go run .'

        6) Наслаждайтесь ботом)