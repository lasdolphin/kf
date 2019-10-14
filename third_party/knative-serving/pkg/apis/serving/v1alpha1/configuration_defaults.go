/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/google/kf/third_party/knative-serving/pkg/apis/serving/v1beta1"
)

func (c *Configuration) SetDefaults(ctx context.Context) {
	ctx = apis.WithinParent(ctx, c.ObjectMeta)
	c.Spec.SetDefaults(apis.WithinSpec(ctx))
}

func (cs *ConfigurationSpec) SetDefaults(ctx context.Context) {
	if v1beta1.IsUpgradeViaDefaulting(ctx) {
		beta := v1beta1.ConfigurationSpec{}
		if cs.ConvertUp(ctx, &beta) == nil {
			alpha := ConfigurationSpec{}
			if alpha.ConvertDown(ctx, beta) == nil {
				*cs = alpha
			}
		}
	}

	cs.GetTemplate().Spec.SetDefaults(ctx)
}