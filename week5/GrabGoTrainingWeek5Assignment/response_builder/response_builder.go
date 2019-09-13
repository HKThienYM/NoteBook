package responsebuilder

import (
	"io"

	enres "thien.com/encode-response"
	getapi "thien.com/get-api"
)

type postWithCommentsResponse struct {
	Posts []postWithComments `json:"posts" xml:"Posts>Post"`
}

type postWithComments struct {
	ID       int64            `json:"id"`
	Title    string           `json:"title"`
	Comments []getapi.Comment `json:"comments,omitempty" xml:"Comments>Comment"`
}

type ResponseBuilder interface {
	Build([]getapi.Post, []getapi.Comment) ([]byte, error)
}

type responseBuilderImpl struct {
	getEncoder func(w io.Writer) *enres.EncodeResponse
}

type resultWriter struct {
	result []byte
}

func (resWriter *resultWriter) Write(p []byte) (n int, err error) {
	resWriter.result = p
	return len(p), nil
}

//Build response the result to the writer inside render
func (builder *responseBuilderImpl) Build(posts []getapi.Post, comments []getapi.Comment) ([]byte, error) {
	postWithComments := combinePostWithComments(posts, comments)
	resp := postWithCommentsResponse{Posts: postWithComments}
	writer := &resultWriter{}
	err := (*builder.getEncoder(writer)).Encode(resp)
	return writer.result, err
}

func New(function func(w io.Writer) *enres.EncodeResponse) ResponseBuilder {
	return &responseBuilderImpl{getEncoder: function}
}

func combinePostWithComments(posts []getapi.Post, comments []getapi.Comment) []postWithComments {
	commentsByPostID := map[int64][]getapi.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]postWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, postWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}
