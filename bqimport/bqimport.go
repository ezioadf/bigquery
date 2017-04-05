package bqimport

import (
	"fmt"
	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func BQImport() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "bigquery-163707"

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("Failed to create client: %v", err)
	}

	// Sets the name for the new dataset.
	datasetName := "test_data"

	// Creates a Dataset instance.
	dataset := client.Dataset(datasetName)

	// Creates the new BigQuery dataset.
	if err := dataset.Create(ctx); err != nil {
		fmt.Println("Failed to create dataset: %v", err)
	}

	fmt.Println("Dataset created\n")
}

func BQRead() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "bigquery-163707"

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("Failed to create client: %v", err)
		return
	}

	q := client.Query(`
		SELECT year, SUM(number) as num
		FROM [bigquery-public-data:usa_names.usa_1910_2013]
		WHERE name = "William"
		GROUP BY year
		ORDER BY year
	`)

	it, err := q.Read(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(values)
	}
}