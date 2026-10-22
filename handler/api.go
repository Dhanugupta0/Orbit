package handler

import (
	"orbit/cache"
	"orbit/guardrails"
	"orbit/provider"
	"orbit/routing"
	"orbit/websearch"
)

// APIHandler groups all LLM API endpoint handlers with their shared dependencies.
type APIHandler struct {
	Registry   *provider.Registry
	Cache      *cache.Cache
	Router     *routing.Router
	Guardrails *guardrails.GlobalConfig
	Search     *websearch.Service
}
