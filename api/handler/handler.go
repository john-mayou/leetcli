package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/john-mayou/leetcli/config"
	"github.com/john-mayou/leetcli/db"
	"github.com/john-mayou/leetcli/internal/metric"
	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/model"
)

type Store struct {
	Problems     map[string]*model.Problem
	ProblemsMeta map[string]*sandbox.ProblemMeta
}

type Handler struct {
	Config     *config.Config
	Now        func() time.Time
	DBClient   db.DBClient
	HTTPClient *http.Client
	Metrics    *metric.MetricsHandler
	Logger     *log.Logger
	Store      *Store
}

type HandlerOpts struct {
	Config     *config.Config
	Now        func() time.Time
	DBClient   db.DBClient
	HTTPClient *http.Client
	Metrics    *metric.MetricsHandler
	Logger     *log.Logger
	Store      *Store
}

func NewHandler(opts *HandlerOpts) (*Handler, error) {
	if opts.Config == nil {
		return nil, errors.New("handler: config cannot be nil")
	}
	if opts.Config == nil {
		return nil, errors.New("handler: now func cannot be nil")
	}
	if opts.DBClient == nil {
		return nil, errors.New("handler: database connection cannot be nil")
	}
	if opts.HTTPClient == nil {
		return nil, errors.New("handler: http client cannot be nil")
	}
	if opts.Metrics == nil {
		return nil, errors.New("handler: metrics handler cannot be nil")
	}
	if opts.Logger == nil {
		return nil, errors.New("handler: logger cannot be nil")
	}
	if opts.Store == nil {
		return nil, errors.New("handler: store cannot be nil")
	}

	return &Handler{
		Config:     opts.Config,
		DBClient:   opts.DBClient,
		HTTPClient: opts.HTTPClient,
		Metrics:    opts.Metrics,
		Logger:     opts.Logger,
		Store:      opts.Store,
	}, nil
}

func NewTestHandler(opts *HandlerOpts) *Handler {
	if opts == nil {
		opts = &HandlerOpts{}
	}
	if opts.Now == nil {
		opts.Now = time.Now
	}
	if opts.Metrics == nil {
		opts.Metrics = metric.NewTestMetricsHandler()
	}
	if opts.Logger == nil {
		opts.Logger = log.Default()
	}

	return &Handler{
		Config:     opts.Config,
		Now:        opts.Now,
		DBClient:   opts.DBClient,
		HTTPClient: opts.HTTPClient,
		Metrics:    opts.Metrics,
		Logger:     opts.Logger,
		Store:      opts.Store,
	}
}
