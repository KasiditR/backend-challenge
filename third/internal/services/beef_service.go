package services

import (
	"context"
	"io"
	"net/http"
	"strings"
	"unicode"

	pb "github.com/KasiditR/backend-challenge-third/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

type BeefServiceServer struct {
	pb.UnimplementedBeefServer
}

func (s *BeefServiceServer) CountBeef(ctx context.Context, em *pb.Empty) (*pb.BeefResponse, error) {
	res, err := fetchBeefData()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	words := countBeef(res)
	beefStruct, err := convertToStruct(words)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BeefResponse{Beef: beefStruct}, nil
}

func convertToStruct(m map[string]int) (*structpb.Struct, error) {
	data := make(map[string]interface{})
	for key, value := range m {
		data[key] = value
	}
	return structpb.NewStruct(data)
}

func fetchBeefData() (string, error) {
	res, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func countBeef(text string) map[string]int {
	text = strings.ToLower(text)

	words := strings.Fields(text)

	cleanWord := func(word string) string {
		var sb strings.Builder
		for _, char := range word {
			if unicode.IsLetter(char) {
				sb.WriteRune(char)
			}
		}
		return sb.String()
	}

	wordCount := make(map[string]int)

	for _, rawWord := range words {
		word := cleanWord(rawWord)
		if word != "" {
			wordCount[word]++
		}
	}
	return wordCount
}
