package query

//JoinRelation join relation query with modifiable struct
type JoinRelation struct {
	//Join LEFT JOIN, INNER JOIN, RIGHT JOIN, JOIN
	Join string
	//TableName example: users User
	TableName string
	//ON example: User.company_id=Company.id
	ON string
}
