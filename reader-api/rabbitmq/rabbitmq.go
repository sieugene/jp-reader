package rabbitmq

import "log"

type RabbitMQConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetRabbitURL(config RabbitMQConfig) string {
	return "amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/"
}
