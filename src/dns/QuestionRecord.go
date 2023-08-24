package dns

type QuestionRecord struct {
	name      string
	queryType QueryType
	class     DNSClass
}
