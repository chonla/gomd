package gomd

// func TestParseMixedATXHeadingDocument(t *testing.T) {
// 	mdText := []byte(`# foo
// ## foo
// ### foo
// #### foo
// ##### foo
// ###### foo`)
// 	expectedMdDoc := &mde.Document{
// 		Children: []mde.Block{
// 			&mde.H1{
// 				Value: "foo",
// 			},
// 			&mde.H2{
// 				Value: "foo",
// 			},
// 			&mde.H3{
// 				Value: "foo",
// 			},
// 			&mde.H4{
// 				Value: "foo",
// 			},
// 			&mde.H5{
// 				Value: "foo",
// 			},
// 			&mde.H6{
// 				Value: "foo",
// 			},
// 		},
// 	}

// 	parser := NewParser()
// 	mdDoc, e := parser.Parse(mdText)

// 	assert.Nil(t, e)
// 	assert.Equal(t, expectedMdDoc, mdDoc)
// }

// func TestParseMoreThan6HashesATX(t *testing.T) {
// 	mdText := []byte(`####### foo`)
// 	expectedMdDoc := &mde.Document{
// 		Children: []mde.Block{
// 			&mde.Text{
// 				Value: "####### foo",
// 			},
// 		},
// 	}

// 	parser := NewParser()
// 	mdDoc, e := parser.Parse(mdText)

// 	assert.Nil(t, e)
// 	assert.Equal(t, expectedMdDoc, mdDoc)
// }

// func TestParseMissingSpaceOrTabATX(t *testing.T) {
// 	mdText := []byte(`#5 bolt

// #hashtag`)
// 	expectedMdDoc := &mde.Document{
// 		Children: []mde.Block{
// 			&mde.Text{
// 				Value: "#5 bolt",
// 			},
// 			&mde.Text{
// 				Value: "#hashtag",
// 			},
// 		},
// 	}

// 	parser := NewParser()
// 	mdDoc, e := parser.Parse(mdText)

// 	assert.Nil(t, e)
// 	assert.Equal(t, expectedMdDoc, mdDoc)
// }

// func TestParseEscapedATX(t *testing.T) {
// 	mdText := []byte(`\## foo`)
// 	expectedMdDoc := &mde.Document{
// 		Children: []mde.Block{
// 			&mde.Text{
// 				Value: "## foo",
// 			},
// 		},
// 	}

// 	parser := NewParser()
// 	mdDoc, e := parser.Parse(mdText)

// 	assert.Nil(t, e)
// 	assert.Equal(t, expectedMdDoc, mdDoc)
// }
