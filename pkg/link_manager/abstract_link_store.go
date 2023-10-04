package link_manager

import (
	om "github.com/TaskMasterErnest/microservice-with-kubernetes/pkg/object_model"
)

type LinkStore interface {
	GetLinks(request om.GetLinksRequest) (om.GetLinksResult, error)
	AddLink(request om.AddLinkRequest) (*om.Link, error)
	UpdateLink(request om.UpdateLinkRequest) (*om.Link, error)
	DeleteLink(username string, url string) error
}
