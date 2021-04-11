package main
import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	// "errors"
	// "github.com/aws/aws-sdk-go/aws"
	// session "github.com/aws/aws-sdk-go/aws/session"
	// dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
// TODO simulate database by using json file

type DbController struct {
	// conn *dynamodb.DynamoDB
	data map[string]interface{}
}

func (c *DbController) InitDbConnection(url string) error{
	data,err := c.LoadAllData()
	if err != nil { 
		fmt.Println("FILE ERROR ",err)
		return  err
	}
	c.data = data
	// fmt.Println("DATABASE URI ",url)
	// c.conn = dynamodb.New(session.New(&aws.Config{
	// 	Region:   aws.String("us-east-1"),
	// 	Endpoint: aws.String(url),
	// }))
	return nil
}

func (c *DbController) LoadAllData() (map[string]interface{},error){

	// TODO read in mock database file 
	filPath := filepath.Join("./", "data.json")
	fmt.Printf("FILEPATH %s \n",filPath)
	data, err := ioutil.ReadFile(filPath)
	if err != nil { 
		fmt.Println("FILE ERROR ",err)
		return nil, err
	}
	jsonData := make(map[string]interface{})
	err = json.Unmarshal(data, &jsonData)
	if err != nil { 
		fmt.Println("UNMARSHAL ERROR",err)
		return nil, err
	}
	fmt.Println("DATA ", jsonData)
	return jsonData, nil
}

// func (ctrl *DbController) Scan(table string) (interface{}, error) {
// 	if ctrl.conn == nil {
// 		return nil, errors.New("db connection error")
// 	}
// 	scanInput := &dynamodb.ScanInput{
// 		TableName: aws.String(table),
// 	}
// 	fmt.Print("SCAN TABLE ", scanInput)
// 	// get all items in table
// 	scanOutput, err := ctrl.conn.Scan(scanInput)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var castTo []interface{}
// 	// https://github.com/mczal/go-gellato-membership/blob/master/service/UserService.go
// 	err = dynamodbattribute.UnmarshalListOfMaps(scanOutput.Items, &castTo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return castTo, nil
// }
