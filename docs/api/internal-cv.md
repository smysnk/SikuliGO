# API: `internal/cv`

[Back to API Index](./)

## Full Package Doc

```text
package cv // import "github.com/sikulix/portgo/internal/cv"


TYPES

type NCCMatcher struct{}

func NewNCCMatcher() *NCCMatcher

func (m *NCCMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type SADMatcher struct{}

func NewSADMatcher() *SADMatcher

func (m *SADMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

```
