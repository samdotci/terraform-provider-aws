// Code generated by "aws/internal/generators/listpages/main.go -function=DescribeImageBuilders github.com/aws/aws-sdk-go/service/appstream"; DO NOT EDIT.

package lister

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appstream"
)

func DescribeImageBuildersPages(conn *appstream.AppStream, input *appstream.DescribeImageBuildersInput, fn func(*appstream.DescribeImageBuildersOutput, bool) bool) error {
	return DescribeImageBuildersPagesWithContext(context.Background(), conn, input, fn)
}

func DescribeImageBuildersPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.DescribeImageBuildersInput, fn func(*appstream.DescribeImageBuildersOutput, bool) bool) error {
	for {
		output, err := conn.DescribeImageBuildersWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}