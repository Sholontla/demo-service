package models

type ProducerChannles struct {
	ChanServiceConsumer1 chan MainProducer
	ChanServiceConsumer2 chan MainProducer
	ChanServiceConsumer3 chan MainProducer
	ChanServiceConsumer4 chan MainProducer
	ChanServiceConsumer5 chan MainProducer
}

type ProducerStringChannles struct {
	ChanStringServiceConsumer1 chan string
	ChanStringServiceConsumer2 chan string
	ChanStringServiceConsumer3 chan string
	ChanStringServiceConsumer4 chan string
	ChanStringServiceConsumer5 chan string
}

type LoggerProducerChannles struct {
	ChanLoggerServiceConsumer1 chan string
	ChanLoggerServiceConsumer2 chan string
	ChanLoggerServiceConsumer3 chan string
	ChanLoggerServiceConsumer4 chan string
	ChanLoggerServiceConsumer5 chan string
}
