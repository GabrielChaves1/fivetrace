package adapters

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"fivetrace.com/iam_service/internal/ports"
)

type DynamoTokenTable struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoTokenTable(client *dynamodb.Client, tableName string) ports.AuthTokenTable {
	return &DynamoTokenTable{
		client:    client,
		tableName: tableName,
	}
}

func (a *DynamoTokenTable) PutToken(ctx context.Context, sub, token string) error {
	_, err := a.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &a.tableName,
		Item: map[string]types.AttributeValue{
			"sub":   &types.AttributeValueMemberS{Value: sub},
			"token": &types.AttributeValueMemberS{Value: token},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (a *DynamoTokenTable) GetToken(ctx context.Context, sub string) (string, error) {
	input := &dynamodb.GetItemInput{
		TableName: &a.tableName,
		Key: map[string]types.AttributeValue{
			"sub": &types.AttributeValueMemberS{Value: sub},
		},
	}

	res, err := a.client.GetItem(ctx, input)

	if err != nil {
		return "", err
	}

	if res.Item == nil {
		return "", nil
	}

	val := res.Item["token"].(*types.AttributeValueMemberS).Value

	return val, nil
}

func (a *DynamoTokenTable) DeleteToken(ctx context.Context, sub string) error {
	_, err := a.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &a.tableName,
		Key: map[string]types.AttributeValue{
			"sub": &types.AttributeValueMemberS{Value: sub},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
