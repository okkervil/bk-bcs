/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"context"
	"fmt"

	k8scorev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/apis/networkextension/v1"
)

func generateRuleConflictMessage(rule *networkextensionv1.IngressRule,
	ingressName, ingressNamespace, lisIngressName, lisIngressNamespace string) string {
	return fmt.Sprintf("[conflict] rule +v of ingress %s/%s is conflict with ingress %s/%s",
		rule, ingressName, ingressNamespace, lisIngressName, lisIngressNamespace)
}

func generateMappingConflictMessage(mapping *networkextensionv1.IngressPortMapping,
	ingressName, ingressNamespace, lisIngressName, lisIngressNamespace string) string {
	return fmt.Sprintf("[conflict] mapping +v of ingress %s/%s is conflict with ingress %s/%s",
		mapping, ingressName, ingressNamespace, lisIngressName, lisIngressNamespace)
}

// func isRuleConflict(ingressName, ingressNamespace string,
// 	rule *networkextensionv1.IngressRule,
// 	existedListenerMap map[int]*networkextensionv1.Listener) (bool, string) {

// 	existedListener, ok := existedListenerMap[rule.Port]
// 	if !ok {
// 		return false, ""
// 	}

// 	// when existed listener with same port was generated by this ingress, we think it is not conflict
// 	// otherwise, return true and return conflict message
// 	lisIngressName := existedListener.Labels[networkextensionv1.LabelKeyForIngressName]
// 	lisIngressNamespace := existedListener.Labels[networkextensionv1.LabelKeyForIngressNamespace]
// 	if lisIngressName != ingressName || lisIngressNamespace != ingressNamespace {
// 		conflictMsg := generateRuleConflictMessage(rule, ingressName, ingressNamespace,
// 			lisIngressName, lisIngressNamespace)
// 		blog.Warnf(conflictMsg)
// 		return true, conflictMsg
// 	}
// 	return false, ""
// }

// func isMappingConflict(ingressName, ingressNamespace string,
// 	mapping *networkextensionv1.IngressPortMapping,
// 	existedListenerMap map[int]*networkextensionv1.Listener) (bool, string) {

// 	for i := mapping.StartIndex; i < mapping.EndIndex; i++ {
// 		existedListener, ok := existedListenerMap[mapping.StartPort+i]
// 		if !ok {
// 			continue
// 		}
// 		lisIngressName := existedListener.Labels[networkextensionv1.LabelKeyForIngressName]
// 		lisIngressNamespace := existedListener.Labels[networkextensionv1.LabelKeyForIngressNamespace]
// 		if lisIngressName != ingressName || lisIngressNamespace != ingressNamespace {
// 			conflictMsg := generateMappingConflictMessage(mapping, ingressName, ingressNamespace,
// 				lisIngressName, lisIngressNamespace)
// 			blog.Warnf(conflictMsg)
// 			return false, conflictMsg
// 		}
// 	}
// 	return false, ""
// }

// func (g *IngressConverter) checkConflicts(ingress *networkextensionv1.Ingress) (bool, error) {
// 	existedListeners := &networkextensionv1.ListenerList{}
// 	err := g.cli.List(context.TODO(), existedListeners, &client.ListOptions{})
// 	if err != nil {
// 		blog.Errorf("failed list existed Listeners err %s", err.Error())
// 		return false, fmt.Errorf("failed list existed Listeners err %s", err.Error())
// 	}

// 	// use map
// 	existedListenerMap := make(map[string]map[int]*networkextensionv1.Listener)
// 	for index, listener := range existedListeners.Items {
// 		// listener for port segment
// 		if listener.Spec.EndPort > 0 {
// 			for i := listener.Spec.Port; i <= listener.Spec.EndPort; i++ {
// 				existedListenerMap[i] = &existedListeners.Items[index]
// 			}
// 			continue
// 		}
// 		existedListenerMap[listener.Spec.Port] = &existedListeners.Items[index]
// 	}

// 	for _, rule := range ingress.Spec.Rules {
// 		isConflict, _ := isRuleConflict(ingress.GetName(), ingress.GetNamespace(), rule, existedListenerMap)
// 		if isConflict {
// 			return true, nil
// 		}
// 	}
// 	for _, mapping := range ingress.Spec.PortMappings {
// 		isConflict, _ := isMappingConflict(ingress.GetName(), ingress.GetNamespace(), mapping, existedListenerMap)
// 		if isConflict {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }

func (g *IngressConverter) convertIngress(ingress *networkextensionv1.Ingress) ([]*networkextensionv1.Listener, error) {

	var generatedListener []*networkextensionv1.Listener
	// for ingress rules
	for _, rule := range ingress.Spec.Rules {

	}

	// for ingress portmappings
	for _, mapping := range ingress.Spec.PortMappings {

	}
	return nil, nil
}

func (g *IngressConverter) convertRule(lbID string,
	rule *networkextensionv1.IngressRule) (*networkextensionv1.Listener, error) {

}

func (g *IngressConverter) generate7LayerListener() (*networkextensionv1.Listener, error) {

}

func (g *IngressConverter) convertMapping(mapping *networkextensionv1.IngressPortMapping) (
	[]*networkextensionv1.Listener, error) {

}
