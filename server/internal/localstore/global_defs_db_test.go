// +build pgsqltest

package localstore

import (
	"math"
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"sourcegraph.com/sourcegraph/sourcegraph/go-sourcegraph/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs"
	sgtest "sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs/testing"
	"sourcegraph.com/sourcegraph/sourcegraph/store"
	"sourcegraph.com/sourcegraph/sourcegraph/store/mockstore"
	"sourcegraph.com/sourcegraph/srclib/graph"
)

func TestGlobalDefs(t *testing.T) {
	t.Parallel()

	var g globalDefs
	ctx, mocks, done := testContext()
	defer done()

	testDefs1 := []*graph.Def{
		{DefKey: graph.DefKey{Repo: "a/b", CommitID: "aaaa", Unit: "a/b/u", UnitType: "t", Path: "abc"}, Name: "ABC", Kind: "func", File: "a.go"},
		{DefKey: graph.DefKey{Repo: "a/b", CommitID: "aaaa", Unit: "a/b/u", UnitType: "t", Path: "xyz/abc"}, Name: "ABC", Kind: "field", File: "a.go"},
		{DefKey: graph.DefKey{Repo: "a/b", CommitID: "aaaa", Unit: "a/b/u", UnitType: "t", Path: "pqr"}, Name: "PQR", Kind: "field", File: "b.go"},
	}

	// if err := g.mustUpdate(ctx, t, "a/b", "a/b/u", "t", testDefs1); err != nil {
	// 	t.Fatal(err)
	// }

	// s := store.GraphFromContext(ctx).(*srcstore.MockMultiRepoStore)
	mockstore.GraphMockDefs(&mocks.Stores.Graph, testDefs1...)
	mocks.RepoVCS.Open_ = func(ctx context.Context, repo string) (vcs.Repository, error) {
		return sgtest.MockRepository{
			ResolveRevision_: func(spec string) (vcs.CommitID, error) {
				return "aaaa", nil
			},
		}, nil
	}
	// mockstore.GraphMockDefs(s, testDefs1...)
	// mockstore.Graph
	g.Update(ctx, []string{"a/b"})

	testCases := []struct {
		Query   []string
		Results []*sourcegraph.SearchResult
	}{
		{
			[]string{"abc"},
			[]*sourcegraph.SearchResult{
				{Def: sourcegraph.Def{Def: *testDefs1[0]}},
				{Def: sourcegraph.Def{Def: *testDefs1[1]}},
			},
		},
		{
			[]string{"pqr"},
			[]*sourcegraph.SearchResult{
				{Def: sourcegraph.Def{Def: *testDefs1[2]}},
			},
		},
	}
	for _, test := range testCases {
		got, err := g.Search(ctx, &store.GlobalDefSearchOp{TokQuery: test.Query})
		if err != nil {
			t.Fatal(err)
		}

		if got == nil {
			t.Errorf("got nil result from GlobalDefs.Search")
			continue
		}

		// strip score
		for _, res := range got.Results {
			res.Score = 0
		}

		if !verifyResultsMatch(got.Results, test.Results) {
			t.Errorf("got %+v, want %+v", got.Results, test.Results)
		}
	}
}

func verifyResultsMatch(got, want []*sourcegraph.SearchResult) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if !reflect.DeepEqual(got[i].Def, want[i].Def) {
			return false
		}
		if got[i].RefCount != want[i].RefCount {
			return false
		}
		if math.Abs(float64(got[i].Score-want[i].Score)) >= 0.0001 {
			return false
		}
	}
	return true
}
