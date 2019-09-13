package responsebuilder

type testCaseBuild struct {
	desc string
	// posts                []getapi.Post
	// comments             []getapi.Comment
	// render               ResponseBuilder

	// 	expectedBuildResult  error
	// 	expectedWriterResult []byte
}

// func getScenarioBuild() []testCaseBuild {
// 	var resWriter = servicehandler.ResultWriter{}
// 	return []testCaseBuild{
// 		testCaseBuild{
// 			desc:  "success",
// 			posts: []getapi.Post{{ID: 1, Title: "Post 1"}},
// 			comments: []getapi.Comment{
// 				{
// 					ID:     1,
// 					Body:   "Comment 1",
// 					PostID: 1,
// 				},
// 			},
// 			expectedBuildResult:  nil,
// 			expectedWriterResult: []byte("asdasdasd"),
// 			writer:               resWriter,
// 			render:               New(json.NewEncoder(&resWriter)),
// 		},
// 	}
// }

// func TestBuild(t *testing.T) {
// 	testCases := getScenarioBuild()
// 	for _, tc := range testCases {
// 		tc := tc
// 		t.Run(tc.desc, func(t *testing.T) {
// 			err := tc.render.Build(tc.posts, tc.comments)
// 			assert.Equal(t, tc.writer.Result, tc.expectedWriterResult)
// 			assert.Equal(t, err, tc.expectedBuildResult)
// 		})
// 	}

// }
