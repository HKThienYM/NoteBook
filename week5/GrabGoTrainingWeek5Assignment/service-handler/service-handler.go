package servicehandler

import (
	json "encoding/json"
	xml "encoding/xml"
	"log"
	"net/http"

	getapi "thien.com/get-api"
	responsebuilder "thien.com/response_builder"
)

const (
	contentJson = "application/json"
	contentXml  = "application/xml"
)

type serviceHandlerImpl struct {
	builder     responsebuilder.ResponseBuilder
	contentType string
}

func Get(isJson bool) func(http.ResponseWriter, *http.Request) {
	var service serviceHandlerImpl
	if isJson {
		service = serviceHandlerImpl{
			builder:     responsebuilder.New(json.NewEncoder),
			contentType: contentJson,
		}
	} else {
		service = serviceHandlerImpl{
			builder:     responsebuilder.New(xml.NewEncoder),
			contentType: contentXml,
		}
	}
	return service.handler
}

func (service *serviceHandlerImpl) handler(writer http.ResponseWriter, request *http.Request) {
	getApi, _ := getapi.New(http.DefaultClient)

	// Get posts from api
	posts, err := getApi.GetPosts()
	if err != nil {
		log.Println("get posts failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	// Get comments from api
	comments, err := getApi.GetComments()
	if err != nil {
		log.Println("get comments failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	// Combine and return response
	res, err := service.builder.Build(posts, comments)
	if err != nil {
		log.Println("unable to parse response: ", err)
		writer.WriteHeader(500)
	}
	writer.Header().Set("Content-Type", service.contentType)
	_, err = writer.Write(res)

}
