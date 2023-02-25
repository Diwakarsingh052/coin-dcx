package main

import (
	"fmt"
	"log"
	"os"
	pb "proto-basics/proto"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	//simpleMessage()
	//nestedMessage()
	//enum()
	//OneOf()
	Maps()
}
func simpleMessage() {
	r := pb.BlogRequest{
		BlogId:  101,
		Title:   "Introduction to Protocol Buffers",
		Content: "Test",
	}
	fmt.Println(r.GetBlogId(), r.GetContent())
	//return a string representation of a Protobuf message.
	//This is useful for debugging and logging purposes
	fmt.Println(r.String())

	//reset the values of a Protobuf message to their default values.
	//This is useful when you want to reuse a message without having to create a new one from scratch
	r.Reset()
	fmt.Println(r.String())

}

func enum() {
	p := pb.Product{
		ProductId: "101",
		Category:  pb.Category_CATEGORY_CLOTHING, // only one category could be set
	}

	//r:= proto.Restaurant{Order: proto.}
	fmt.Println(p.GetCategory())
}

func OneOf() {
	// Create a new Restaurant message with a breakfast order
	restaurant := &pb.Restaurant{
		// order field is an interface // type that impl that interface could be assigned here
		Order: &pb.Restaurant_Breakfast{
			Breakfast: "pancakes",
		},
	}

	res := restaurant.GetOrder()

	// Check the type of the order
	switch order := res.(type) {
	case *pb.Restaurant_Breakfast:
		fmt.Printf("Order is breakfast: %v\n", order.Breakfast)
	case *pb.Restaurant_Lunch:
		fmt.Printf("Order is lunch: %v\n", order.Lunch)
	case *pb.Restaurant_Dinner:
		fmt.Printf("Order is dinner: %v\n", order.Dinner)
	default:
		fmt.Println("Unknown order type")
	}
}

func nestedMessage() {
	// Create a new SearchResponse message with some results
	searchResponse := &pb.SearchResponse{
		Results: []*pb.SearchResponse_Result{
			{
				Url:   "https://grpc.io/",
				Title: "gRPC",
			},
			{
				Url:   "https://pkg.go.dev/",
				Title: "go packages",
			},
		},
	}

	log.Println("using nested types")
	// Print the URLs and titles of the results
	for _, result := range searchResponse.GetResults() {
		// adding result value in search result struct which was using
		fmt.Printf("Result URL: %s\n", result.GetUrl())
		fmt.Printf("Result Title: %s\n", result.GetTitle())
	}

	fmt.Println()
	log.Println("reusing Result message type outside SearchResponse message type")
	rs := searchResponse.GetResults()
	sResult := pb.SearchResult{Result: rs[1]}
	fmt.Println(sResult.GetResult())

}

func Maps() {
	// Create a new Product message
	product := &pb.ProductDetails{
		Id:   199,
		Name: "Mac Pro",
		PricesByCurrency: map[string]float64{
			"USD": 3000.99,
			"EUR": 2800.99,
			"JPY": 400000.0,
		},
	}

	// Print the deserialized message
	fmt.Println("ID:", product.Id)
	fmt.Println("Name:", product.Name)
	fmt.Println("Prices by currency:")
	for currency, price := range product.PricesByCurrency {
		fmt.Printf("  %s: %.2f\n", currency, price)
	}

	//protojson marshal and unmarshal protocol buffer messages as JSON format.
	data, err := protojson.Marshal(product)
	if err != nil {
		log.Fatal("Error marshaling product:", err)
	}

	err = os.WriteFile("test.json", data, 0666)
	if err != nil {
		log.Fatalln(err)
	}

}
