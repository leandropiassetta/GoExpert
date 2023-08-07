package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.DynamoDB
	TableName      string
	AttributesName []string
}

// // official documentation https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.CreateTable
// func (basics TableBasics) tableExists() (bool, error) {
// 	exists := true
// 	// Verifique se a tabela existe. Se não existir, crie-a.
// 	_, err := basics.DynamoDbClient.DescribeTable(
// 		&dynamodb.DescribeTableInput{TableName: aws.String(basics.TableName)},
// 	)
// 	// Se a tabela não existir, o erro será do tipo ResourceNotFoundException.
// 	if err != nil {
// 		var notFoundTable *dynamodb.ResourceNotFoundException
// 		// Se o erro for do tipo ResourceNotFoundException, a tabela não existe.

// 		// errors.As() verifica se o erro é do tipo ResourceNotFoundException.
// 		// errors.As() -> faz a mesma coisa que o errors.Is(), mas verifica se o erro é do tipo ResourceNotFoundException.
// 		if errors.As(err, &notFoundTable) {
// 			log.Printf("Table %v does not exist.\n", basics.TableName)
// 			err = nil
// 			// Se o erro não for do tipo ResourceNotFoundException, algo deu errado.
// 		} else {
// 			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", basics.TableName, err)
// 		}
// 		exists = false
// 	}
// 	return exists, err
// }

// func (basics TableBasics) waitForTableExists() error {
// 	waiter := dynamodb.NewTableExistsWaiter(basics.DynamoDbClient)

// 	// Esperar até que a tabela exista (no máximo 5 minutos)
// 	return waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
// 		TableName: aws.String(basics.TableName),
// 	}, 5*time.Minute)
// }

// func (basics TableBasics) CreateMovieTable() (*dynamodb.TableDescription, error) {
// 	var tableDesc *dynamodb.TableDescription

// 	// Definir atributos da tabela
// 	attributeDefinitions := []*dynamodb.AttributeDefinition{
// 		{
// 			AttributeName: aws.String("year"),
// 			AttributeType: aws.String(dynamodb.ScalarAttributeTypeN),
// 		},
// 		{
// 			AttributeName: aws.String("title"),
// 			AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
// 		},
// 	}

// 	// Definir esquema de chave primária
// 	keySchema := []*dynamodb.KeySchemaElement{
// 		{
// 			AttributeName: aws.String("year"),
// 			KeyType:       aws.String(dynamodb.KeyTypeHash),
// 		},
// 		{
// 			AttributeName: aws.String("title"),
// 			KeyType:       aws.String(dynamodb.KeyTypeRange),
// 		},
// 	}

// 	// Criar uma tabela usando CreateTableInput
// 	table, err := basics.DynamoDbClient.CreateTable(&dynamodb.CreateTableInput{
// 		AttributeDefinitions: attributeDefinitions,
// 		KeySchema:            keySchema,
// 		TableName:            aws.String(basics.TableName),
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(10),
// 		},
// 	})

// 	if err != nil {
// 		// Se ocorrer um erro ao criar a tabela, registre o motivo do erro
// 		log.Printf("Couldn't create table %v. Here's why: %v\n", basics.TableName, err)
// 	} else {
// 		// Criar um monitor para aguardar a existência da tabela
// 		waiter := dynamodb.NewTableExistsWaiter(basics.DynamoDbClient)

// 		// Esperar até que a tabela exista (no máximo 5 minutos)
// 		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
// 			TableName: aws.String(basics.TableName),
// 		}, 5*time.Minute)

// 		if err != nil {
// 			// Se ocorrer um erro ao esperar pela existência da tabela, registre o motivo do erro
// 			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
// 		}

// 		// Atribuir a descrição da tabela criada à variável tableDesc
// 		tableDesc = table.TableDescription
// 	}

// 	// Retornar a descrição da tabela e o possível erro
// 	return tableDesc, err
// }

// tableExists verifica se a tabela já existe no DynamoDB.
func (table *TableBasics) tableExists() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// &dynamodb.DescribeTableInput -> cria um ponteiro para um objeto do tipo DescribeTableInput.

	// DescribeTableInput -> descreve uma tabela e retorna informações sobre ela.
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(table.TableName),
	}

	_, err := table.DynamoDbClient.DescribeTableWithContext(ctx, input)
	if err != nil {
		// Se o erro for do tipo ResourceNotFoundException, a tabela não existe.
		if aerr, ok := err.(awserr.Error); ok {
			// Se a tabela não existe, retornamos falso sem erro.
			log.Printf("Table %s does not exist", table.TableName)
			if aerr.Code() == dynamodb.ErrCodeResourceNotFoundException {
				return false, nil
			}
		}

		// Se ocorrer outro erro, retornamos o erro.
		log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", table.TableName, err)

		return false, err
	}

	// Se não houver erro, a tabela existe.
	return true, nil
}

// CreateTable cria uma tabela no DynamoDB, se ela ainda não existir.
func (table *TableBasics) CreateTable() {
	exists, err := table.tableExists()
	if err != nil {
		log.Println("Error checking table existence:", err)
		return
	}

	if exists {
		log.Println("Table already exists")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Crie a tabela. Defina os atributos e as chaves primárias. Defina a capacidade de leitura/gravação.

	// Defina o nome da tabela. Defina o tipo de atributo. Defina o tipo de chave primária. Defina a capacidade de leitura/gravação.
	input := &dynamodb.CreateTableInput{
		// Defina os atributos e as chaves primárias. Defina o nome da tabela. Defina o tipo de atributo. Defina o tipo de chave primária. Defina a capacidade de leitura/gravação.
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				// Defina o nome do atributo. O nome do atributo pode ser qualquer nome. O nome do atributo é o nome da coluna.
				AttributeName: aws.String("id"),
				// Defina o tipo de atributo. O tipo de atributo pode ser S (string), N (número) ou B (binário).
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("song_title"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("artist"),
				AttributeType: aws.String("S"),
			},
		},

		// Defina as chaves primárias. Defina o nome da tabela. Defina o tipo de atributo. Defina o tipo de chave primária. Defina a capacidade de leitura/gravação.

		// key schema -> define as chaves primárias da tabela. A chave primária pode ser uma chave primária de partição ou uma chave primária de partição e classificação. A chave primária é o nome da coluna.

		// key type -> define o tipo de chave primária. O tipo de chave primária pode ser HASH (chave primária de partição) ou RANGE (chave primária de partição e classificação).

		// a diferença entre chave primária de partição e chave primária de partição e classificação é que a chave primária de partição é uma chave primária simples, enquanto a chave primária de partição e classificação é uma chave primária composta. A chave primária composta é uma combinação de chave primária de partição e chave primária de classificação. A chave primária composta é usada para classificar os itens na tabela, ex: classificar por data de criação.
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},

		TableName: aws.String(table.TableName),

		// ProvisionedThroughput -> define a capacidade de leitura/gravação da tabela. A capacidade de leitura/gravação é a quantidade de leitura/gravação que a tabela pode suportar

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			// ReadCapacityUnits -> define a capacidade de leitura da tabela. A capacidade de leitura é a quantidade de leitura que a tabela pode suportar. aws.Int64(10) -> define a capacidade de leitura como 10 unidades de leitura.
			ReadCapacityUnits: aws.Int64(10),

			// WriteCapacityUnits -> define a capacidade de gravação da tabela. A capacidade de gravação é a quantidade de gravação que a tabela pode suportar. aws.Int64(10) -> define a capacidade de gravação como 10 unidades de gravação.
			WriteCapacityUnits: aws.Int64(10),
		},
	}

	_, err = table.DynamoDbClient.CreateTableWithContext(ctx, input)
	if err != nil {
		log.Println("Error creating table:", err)
		return
	}

	log.Println("Table created")
}

// insert a record into dynamoDB

func (table *TableBasics) InsertRecord() {
	// Criar contexto com timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar item para inserir na tabela
	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String("1"),
		},
		"song_title": {
			S: aws.String("Yellow Submarine"),
		},
		"artist": {
			S: aws.String("The Beatles"),
		},
	}

	// Criar input para inserir item na tabela
	input := &dynamodb.PutItemInput{
		// Definir nome da tabela
		TableName: aws.String(table.TableName),
		// Definir item a ser inserido na tabela
		Item: item,
	}

	// Inserir item na tabela
	_, err := table.DynamoDbClient.PutItemWithContext(ctx, input)
	if err != nil {
		log.Println("Error inserting item into table:", err)
		return
	}

	log.Println("Item inserted into table")
}

// update a record in dynamoDB

// delete a record from dynamoDB

// query a record from dynamoDB

// scan a record from dynamoDB
func (table *TableBasics) ScanRecords() {
	// Criar contexto com timeout de 5 segundos
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar input para scan
	input := &dynamodb.ScanInput{
		// Definir nome da tabela
		TableName: aws.String(table.TableName),
		Limit:     aws.Int64(10000),
	}

	// Scan -> retorna todos os itens da tabela
	// result, err := table.DynamoDbClient.ScanWithContext(ctx, input)

	result, err := table.DynamoDbClient.Scan(input)
	if err != nil {
		log.Println("Error scanning table:", err)
		return
	}

	// range -> percorre todos os itens da tabela
	for _, item := range result.Items {
		// range -> percorre todos os atributos do item
		for key, value := range item {
			log.Println("key:", key, "value:", value)
		}
	}
}

// pega atributos de uma tabela
func (table *TableBasics) GetTableAttributes() {
	// Criar contexto com timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criar input para pegar atributos da tabela
	input := &dynamodb.DescribeTableInput{
		// Definir nome da tabela
		TableName: aws.String(table.TableName),
	}

	// Pegar atributos da tabela
	result, err := table.DynamoDbClient.DescribeTableWithContext(ctx, input)
	if err != nil {
		log.Println("Error getting table attributes:", err)
		return
	}

	log.Println("Table attributes:", result.Table.AttributeDefinitions)
}

func main() {
	// Configurar sessão AWS

	sess, err := session.NewSession(&aws.Config{
		// Definir região AWS
		Region: aws.String("us-east-1"),

		// Definir credenciais

		// Definir credenciais estáticas

	})
	if err != nil {
		log.Println("Error creating session:", err)
	}

	// Criar cliente DynamoDB

	// dynamodb.New -> retorna um novo cliente DynamoDB. O cliente DynamoDB é usado para interagir com o DynamoDB.

	// sess -> sessão AWS. A sessão AWS é usada para criar um cliente DynamoDB. O cliente DynamoDB é usado para interagir com o DynamoDB.
	dynamoDBClient := dynamodb.New(sess)

	// Criar instância de TableBasics
	basics := TableBasics{
		DynamoDbClient: dynamoDBClient,
		TableName:      "DiscoveryDynamoTest",
	}

	// Chamar o método para criar a tabela
	basics.CreateTable()

	// Chamar o método para ler os itens da tabela
	basics.ScanRecords()

	// Chamar o método para pegar os atributos da tabela
	basics.GetTableAttributes()
}
