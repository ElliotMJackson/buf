// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-api. DO NOT EDIT.

package registryv1alpha1api

import (
	context "context"
	v1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/image/v1"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
)

// PluginCurationService manages curated plugins.
type PluginCurationService interface {
	// ListCuratedPlugins returns all the curated plugins available.
	ListCuratedPlugins(
		ctx context.Context,
		pageSize uint32,
		pageToken string,
		reverse bool,
	) (plugins []*v1alpha1.CuratedPlugin, nextPageToken string, err error)
	// CreateCuratedPlugin creates a new curated plugin.
	CreateCuratedPlugin(
		ctx context.Context,
		owner string,
		name string,
		registryType v1alpha1.PluginRegistryType,
		version string,
		containerImageDigest string,
		defaultOptions []string,
		dependencies []*v1alpha1.CuratedPluginReference,
		sourceUrl string,
		description string,
		registryConfig *v1alpha1.RegistryConfig,
		revision uint32,
	) (configuration *v1alpha1.CuratedPlugin, err error)
	// GetLatestCuratedPlugin returns the latest version of a plugin matching given parameters.
	GetLatestCuratedPlugin(
		ctx context.Context,
		owner string,
		name string,
		version string,
		revision uint32,
	) (plugin *v1alpha1.CuratedPlugin, versions []*v1alpha1.CuratedPluginVersionRevisions, err error)
}

// CodeGenerationService generates code using remote plugins.
type CodeGenerationService interface {
	// GenerateCode generates code using the specified remote plugins.
	GenerateCode(
		ctx context.Context,
		image *v1.Image,
		requests []*v1alpha1.PluginGenerationRequest,
		includeImports bool,
		includeWellKnownTypes bool,
	) (responses []*v1alpha1.PluginGenerationResponse, err error)
}
