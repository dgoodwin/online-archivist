// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/api/oauth/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OAuthAccessTokenLister helps list OAuthAccessTokens.
type OAuthAccessTokenLister interface {
	// List lists all OAuthAccessTokens in the indexer.
	List(selector labels.Selector) (ret []*v1.OAuthAccessToken, err error)
	// Get retrieves the OAuthAccessToken from the index for a given name.
	Get(name string) (*v1.OAuthAccessToken, error)
	OAuthAccessTokenListerExpansion
}

// oAuthAccessTokenLister implements the OAuthAccessTokenLister interface.
type oAuthAccessTokenLister struct {
	indexer cache.Indexer
}

// NewOAuthAccessTokenLister returns a new OAuthAccessTokenLister.
func NewOAuthAccessTokenLister(indexer cache.Indexer) OAuthAccessTokenLister {
	return &oAuthAccessTokenLister{indexer: indexer}
}

// List lists all OAuthAccessTokens in the indexer.
func (s *oAuthAccessTokenLister) List(selector labels.Selector) (ret []*v1.OAuthAccessToken, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OAuthAccessToken))
	})
	return ret, err
}

// Get retrieves the OAuthAccessToken from the index for a given name.
func (s *oAuthAccessTokenLister) Get(name string) (*v1.OAuthAccessToken, error) {
	key := &v1.OAuthAccessToken{ObjectMeta: meta_v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("oauthaccesstoken"), name)
	}
	return obj.(*v1.OAuthAccessToken), nil
}
